package kafka

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	"github.com/gogo/protobuf/proto"
	"github.com/phinc275/gfas/pkg/mq"
	"github.com/phinc275/taskfi-common/go/common"
)

type Config struct {
	BootstrapServers string `mapstructure:"bootstrapServers"`
	GroupID          string `mapstructure:"groupID"`
	AutoOffsetReset  string `mapstructure:"autoOffsetReset"`

	SchemaRegistryURL string `mapstructure:"schemaRegistryURL"`
}

type MessageQueue struct {
	cfg          Config
	deserializer serde.Deserializer
}

func NewMessageQueue(cfg *Config) (*MessageQueue, error) {
	schemaRegistryConfig := schemaregistry.NewConfig(cfg.SchemaRegistryURL)
	schemaRegistryClient, err := schemaregistry.NewClient(schemaRegistryConfig)
	if err != nil {
		return nil, err
	}

	deserializer, err := protobuf.NewDeserializer(schemaRegistryClient, serde.ValueSerde, protobuf.NewDeserializerConfig())
	if err != nil {
		return nil, err
	}

	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventJobCompleted{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventJobApplied{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventJobPosted{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventUserAccessed{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventWorkspaceCompleted{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventMoneySpent{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventProfileVerified{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventProfileViewed{}).ProtoReflect().Type())

	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventSocialConnected{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventSocialRankingUpdated{}).ProtoReflect().Type())
	_ = deserializer.ProtoRegistry.RegisterMessage((&common.EventLoyaltyEarned{}).ProtoReflect().Type())

	return &MessageQueue{cfg: *cfg, deserializer: deserializer}, nil
}

func (queue *MessageQueue) SubscribeTopics(topics []string) (<-chan mq.MsgResponse, mq.Unsubscribe, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": queue.cfg.BootstrapServers,
		"group.id":          queue.cfg.GroupID,
		"auto.offset.reset": queue.cfg.AutoOffsetReset,
	})
	if err != nil {
		return nil, nil, err
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		return nil, nil, err
	}

	doneCh := make(chan struct{}, 1)

	unsubscribe := func() error {
		if c == nil {
			return nil
		}
		close(doneCh)
		return c.Close()
	}

	resultChan := make(chan mq.MsgResponse, 100) // buffer 100 is enough?
	go func() {
		for {
			select {
			case <-doneCh:
				return
			default:
				msg, err := c.ReadMessage(1 * time.Second)
				if err == nil {
					// potentially blocking, but I don't want to lose any msg
					deserializedMsg, deserErr := queue.deserializer.Deserialize(*msg.TopicPartition.Topic, msg.Value)
					if deserErr != nil {
						continue
					}

					deserializedProtoMsg, ok := deserializedMsg.(proto.Message)
					if !ok {
						continue
					}

					resultChan <- mq.MsgResponse{
						Err:     nil,
						Message: deserializedProtoMsg,
					}
					continue
				}

				if err.(kafka.Error).IsTimeout() {
					continue
				}

				resultChan <- mq.MsgResponse{Err: err}
			}
		}
	}()

	return resultChan, unsubscribe, nil
}
