package events

type UserRegisteredEvent struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
