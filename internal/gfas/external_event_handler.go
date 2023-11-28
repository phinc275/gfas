package gfas

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/phinc275/gfas/pkg/logger"
	"github.com/phinc275/gfas/pkg/mq"
	"github.com/phinc275/taskfi-common/go/core"
	"github.com/phinc275/taskfi-common/go/dlancer"
)

type ExternalEventHandler struct {
	topics  []string
	logger  logger.Logger
	mq      mq.MessageQueue
	service *UserAchievementsService

	unsubscribeFn mq.Unsubscribe
	numOfWorkers  int
	jobCh         chan job
}

type job struct {
	aggregateID string
	event       ExternalEvent
}

func NewExternalEventHandler(logger logger.Logger, topics []string, numOfWorkers int, mq mq.MessageQueue, service *UserAchievementsService) *ExternalEventHandler {
	return &ExternalEventHandler{
		logger:       logger,
		topics:       topics,
		numOfWorkers: numOfWorkers,
		mq:           mq,
		service:      service,
	}
}

func (handler *ExternalEventHandler) Start() error {
	msgCh, unsubscribeFn, err := handler.mq.SubscribeTopics(handler.topics)
	if err != nil {
		return err
	}

	numOfWorkers := handler.numOfWorkers
	if numOfWorkers == 0 {
		numOfWorkers = 1 // should raise error instead ?
	}

	handler.jobCh = make(chan job, numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {
		go handler.workerFn()
	}

	handler.unsubscribeFn = unsubscribeFn
	go func() {
		for msgResp := range msgCh {
			if msgResp.Err != nil {
				handler.logger.Errorf("consumer error: %v\n", msgResp.Err)
				continue
			}

			aggregateID, externalEvent, err := externalEventFromMessage(msgResp.Message)
			if err != nil {
				handler.logger.Errorf("parse event error: %v\n", msgResp.Err)
				continue
			}

			handler.jobCh <- job{
				aggregateID: aggregateID,
				event:       externalEvent,
			}
		}
	}()

	return nil
}

func (handler *ExternalEventHandler) Stop() error {
	if handler.unsubscribeFn != nil {
		_ = handler.unsubscribeFn()
	}

	if handler.jobCh != nil {
		close(handler.jobCh)
	}

	return nil
}

func (handler *ExternalEventHandler) workerFn() {
	for job := range handler.jobCh {
		err := handler.service.commands.applyExternalEvent.Handle(context.Background(), NewApplyExternalEventCommand(job.aggregateID, job.event))
		if err != nil {
			handler.logger.Errorf("(EventHandler).ApplyExternalEvent error: %v", err)
		}
	}
}

func externalEventFromMessage(msg proto.Message) (string, ExternalEvent, error) {
	switch v := msg.(type) {
	case *core.EventSocialConnected:
		return v.GetUserId(), ExternalEvent(&ExternalEventSocialConnected{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
			Provider:  v.GetProvider(),
		}), nil

	case *core.EventSocialRankingUpdated:
		return v.GetUserId(), ExternalEvent(&ExternalEventSocialRankingUpdated{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
			Provider:  v.GetProvider(),
			Rank:      v.GetRank(),
		}), nil

	case *core.EventLoyaltyEarned:
		return v.GetUserId(), ExternalEvent(&ExternalEventLoyaltyPointsEarned{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
			Amount:    v.GetAmount(),
		}), nil

	case *dlancer.EventJobCompleted:
		return v.GetUserId(), ExternalEvent(&ExternalEventJobCompleted{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
			JobId:     v.GetJobId(),
			JobCategory: map[string]string{
				"it-development":       string(AchievementCategoryITDevelopment),
				"smart-contract":       string(AchievementCategorySmartContract),
				"design-and-creative":  string(AchievementCategoryDesignAndCreative),
				"sales-and-marketing":  string(AchievementCategorySalesAndMarketing),
				"kol-and-web3-advisor": string(AchievementCategoryKOLAndWeb3Advisor),
			}[v.GetJobCategory()],
		}), nil

	case *dlancer.EventJobApplied:
		return v.GetUserId(), ExternalEvent(&ExternalEventJobApplied{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
			JobId:     v.GetJobId(),
			JobCategory: map[string]string{
				"it-development":       string(AchievementCategoryITDevelopment),
				"smart-contract":       string(AchievementCategorySmartContract),
				"design-and-creative":  string(AchievementCategoryDesignAndCreative),
				"sales-and-marketing":  string(AchievementCategorySalesAndMarketing),
				"kol-and-web3-advisor": string(AchievementCategoryKOLAndWeb3Advisor),
			}[v.GetJobCategory()],
		}), nil

	case *dlancer.EventJobPosted:
		return v.GetUserId(), ExternalEvent(&ExternalEventJobPosted{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
			JobID:     v.GetJobId(),
			JobCategory: map[string]string{
				"it-development":       string(AchievementCategoryITDevelopment),
				"smart-contract":       string(AchievementCategorySmartContract),
				"design-and-creative":  string(AchievementCategoryDesignAndCreative),
				"sales-and-marketing":  string(AchievementCategorySalesAndMarketing),
				"kol-and-web3-advisor": string(AchievementCategoryKOLAndWeb3Advisor),
			}[v.GetJobCategory()],
		}), nil

	case *dlancer.EventUserAccessed:
		return v.GetUserId(), ExternalEvent(&ExternalEventUserAccessed{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
		}), nil

	case *dlancer.EventWorkspaceCompleted:
		return v.GetUserId(), ExternalEvent(&ExternalEventWorkspaceCompleted{
			Timestamp:   v.GetTimestamp().AsTime(),
			UserID:      v.GetUserId(),
			WorkspaceID: v.GetWorkspaceId(),
		}), nil

	case *dlancer.EventMoneySpent:
		return v.GetUserId(), ExternalEvent(&ExternalEventMoneySpent{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
			Amount:    v.GetAmount(),
		}), nil

	case *dlancer.EventProfileVerified:
		return v.GetUserId(), ExternalEvent(&ExternalEventProfileVerified{
			Timestamp:  v.GetTimestamp().AsTime(),
			UserID:     v.GetUserId(),
			VerifyType: v.GetVerifyType(),
		}), nil

	case *dlancer.EventProfileViewed:
		return v.GetUserId(), ExternalEvent(&ExternalEventProfileViewed{
			Timestamp: v.GetTimestamp().AsTime(),
			UserID:    v.GetUserId(),
		}), nil
	default:
		return "", nil, fmt.Errorf("unrecognized event %T", v)
	}
}
