package telegram

import "github.com/Kwynto/read-adviser-bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client) {
}