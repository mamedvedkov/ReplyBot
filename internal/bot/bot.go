package bot

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	tele "gopkg.in/telebot.v3"

	"github.com/mamedvedkov/ReplyBot/internal"
)

type Config struct {
	Token      string `required:"true" desc:"Botfather token"`
	HomeChatId int64  `default:"123" split_words:"true" desc:"Chat id where to forward messages"`
}

type Bot struct {
	svc *internal.Service

	tele     *tele.Bot
	homeChat homeChat
}

func Must(cfg Config, logger logr.Logger, svc *internal.Service) *Bot {
	b, err := New(cfg, logger, svc)
	if err != nil {
		panic(err)
	}

	return b
}

func New(cfg Config, logger logr.Logger, svc *internal.Service) (*Bot, error) {
	b, err := tele.NewBot(tele.Settings{
		Token:   cfg.Token,
		OnError: OnErrorLog(logger.WithName("bot")),
	})
	if err != nil {
		return nil, err
	}

	bot := &Bot{svc: svc, tele: b, homeChat: newHomeChat(cfg.HomeChatId)}

	bot.routes()

	return bot, nil
}

func (b *Bot) routes() {
	middlewares := []tele.MiddlewareFunc{
		ForwardIfNeeded(b.homeChat, b.tele, b.svc),
		IsSkip(b.tele.Me.ID),
	}

	b.tele.Handle(tele.OnText, b.handleText, middlewares...)
	b.tele.Handle(tele.OnMedia, b.handleMedia, middlewares...)
	b.tele.Handle(tele.OnSticker, b.handleSticker, middlewares...)
	b.tele.Handle("/my_id", b.handleMyId)
	b.tele.Handle("/start", func(c tele.Context) error { return nil })
}

func (b *Bot) Start(ctx context.Context) error {
	go b.tele.Start()
	<-ctx.Done()
	b.tele.Stop()
	return ctx.Err()
}

func (b *Bot) handleMyId(ctx tele.Context) error {
	_, err := b.tele.Send(ctx.Chat(), fmt.Sprintf("%v", ctx.Message().Chat.ID))
	return err
}

func (b *Bot) handleText(ctx tele.Context) error {
	originalTgId, ok := b.svc.GetSession(ctx.Message().ReplyTo.ID)
	if !ok {
		return fmt.Errorf("cant get original tg id")
	}

	_, err := b.tele.Send(newRecipient(int64(originalTgId)), ctx.Text())
	return err

}

func (b *Bot) handleMedia(ctx tele.Context) error {
	return nil
}

func (b *Bot) handleSticker(ctx tele.Context) error {
	return nil
}
