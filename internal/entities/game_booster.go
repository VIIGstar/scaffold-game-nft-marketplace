package entities

type Booster struct {
	ItemType
	Type string `json:"type"` // Exp booster | Gold booster | Luck booster...
}
