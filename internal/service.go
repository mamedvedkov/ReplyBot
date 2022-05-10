package internal

import (
	"context"
	"github.com/go-logr/logr"
)

type BlacklistStore interface {
	AddBlacklist(ctx context.Context, tgId int) error
	DeleteBlacklist(ctx context.Context, tgId int) error
	IsBlacklisted(ctx context.Context, tgId int) (bool, error)
}

type AdminStore interface {
	AddAdmin(ctx context.Context, tgId int) error
	DeleteDelete(ctx context.Context, tgId int) error
	IsAdmin(ctx context.Context, tgId int) (bool, error)
}

type SessionStore interface {
	SaveSession(ctx context.Context, msgId, tgId int) error
	GetSession(ctx context.Context, msgId int) (tgId int, err error)
}

type Service struct {
	logger logr.Logger

	bl       BlacklistStore
	ad       AdminStore
	sessions SessionStore
}

func NewService(logger logr.Logger, bl BlacklistStore, ad AdminStore, sessions SessionStore) *Service {
	return &Service{logger: logger.WithName("service"), bl: bl, ad: ad, sessions: sessions}
}

func (svc *Service) SaveSession(msgId, tgId int) {
	err := svc.sessions.SaveSession(context.TODO(), msgId, tgId)
	if err != nil {
		svc.logger.Error(err, "cant save session", "msgId", msgId, "tgId", tgId)
	}
}

func (svc *Service) GetSession(msgId int) (int, bool) {
	tgId, err := svc.sessions.GetSession(context.TODO(), msgId)
	if err != nil {
		svc.logger.Error(err, "cant get session", "msgId", msgId)
		return 0, false
	}

	return tgId, true
}
