package storage

import (
	"context"
	"fmt"
)

const saveSessionQuery = `
INSERT INTO reply_bot.msg_sessions (msg_id, telegram_id, created_at)
VALUES ($1, $2, DEFAULT);
`

func (p *PG) SaveSession(ctx context.Context, msgId, tgId int) error {
	_, err := p.pool.Exec(ctx, saveSessionQuery, msgId, tgId)
	if err != nil {
		return fmt.Errorf("error from db: %w", err)
	}

	return nil
}

const getSessionQuery = `
SELECT telegram_id
FROM reply_bot.msg_sessions
WHERE msg_id = $1;
`

func (p *PG) GetSession(ctx context.Context, msgId int) (tgId int, err error) {
	err = p.pool.QueryRow(ctx, getSessionQuery, msgId).Scan(&tgId)
	if err != nil {
		return 0, fmt.Errorf("error from db: %w", err)
	}

	return tgId, nil
}
