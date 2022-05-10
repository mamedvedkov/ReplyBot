package bot

import (
	"fmt"
)

func newRecipient(id int64) recipient {
	return recipient{id: fmt.Sprintf("%v", id)}
}

type recipient struct {
	id string
}

func (r recipient) Recipient() string {
	return r.id
}

type homeChat struct {
	id int64
	recipient
}

func newHomeChat(id int64) homeChat {
	return homeChat{
		id:        id,
		recipient: newRecipient(id),
	}
}
