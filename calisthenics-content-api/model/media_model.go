package model

import "calisthenics-content-api/config"

type PlayAction struct {
	ActionType config.PlayActionType
	ButtonText string
}

type PlayActionRequest struct {
	Token    string
	LangCode string
}
