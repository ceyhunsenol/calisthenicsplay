package config

type GeneralInfo string

const (
	AccessLevel GeneralInfo = "ACCESS_LEVEL"
)

type Audience string

const (
	Everyone     Audience = "EVERYONE"
	User         Audience = "USER"
	Subscription Audience = "SUBSCRIPTION" // ileride sub yapildigi zaman ele alinacak
)

type PlayActionType string

const (
	Watch    PlayActionType = "WATCH"
	Register PlayActionType = "REGISTER"
	Sub      PlayActionType = "SUBSCRIPTION"
)
