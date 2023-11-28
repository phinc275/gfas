package mq

import (
	"github.com/gogo/protobuf/proto"
)

type MsgResponse struct {
	Err     error
	Message proto.Message
}

type Unsubscribe func() error

type MessageQueue interface {
	SubscribeTopics([]string) (<-chan MsgResponse, Unsubscribe, error)
}
