package entities

type ItemType struct {
	DefaultModel
	BaseGameEntity
	Name string `json:"name"` // Gun | Fashion | Booster
	Code string `json:"code"` // Gun | Fashion | Booster
}
