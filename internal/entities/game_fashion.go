package entities

type Fashion struct {
	ItemType
	Type string `json:"type"` // Gloves | Glass | Boot ...
}
