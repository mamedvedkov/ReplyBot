package bot

import (
	"encoding/json"

	"github.com/go-logr/logr"
	tele "gopkg.in/telebot.v3"

	"github.com/mamedvedkov/ReplyBot/internal"
)

func ForwardIfNeeded(homeChat homeChat, bot *tele.Bot, svc *internal.Service) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			if ctx.Message().Chat.ID == homeChat.id {
				return next(ctx)

			}

			msg, err := bot.Forward(homeChat, ctx.Message())
			if err != nil {
				return err
			}

			svc.SaveSession(msg.ID, int(ctx.Message().Sender.ID))
			return nil
		}
	}
}

func IsSkip(me int64) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			if isSkip(ctx.Message(), me) {
				return nil
			}

			return next(ctx)
		}
	}
}

func isSkip(message *tele.Message, me int64) bool {
	return !message.IsReply() || message.ReplyTo.Sender.ID != me
}

func LogReq(logger logr.Logger) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			data, _ := json.MarshalIndent(ctx.Update(), "", "  ")
			logger.Info(string(data))
			return next(ctx)
		}
	}
}

func OnErrorLog(logger logr.Logger) func(err error, c tele.Context) {
	return func(err error, c tele.Context) {
		if err != nil {
			logger.Error(err, "error ocured")
		}
	}
}
