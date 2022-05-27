package entities

type Item struct {
	DefaultModel
	BaseGameEntity
	UserId int64 `json:"user_id"`
}
