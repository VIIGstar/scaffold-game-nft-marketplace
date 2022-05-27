package entities

type Gun struct {
	ItemType
	Type string `json:"type"` // FPS Game: Type gun -> Assault rifle | Sniper rifles | Shotguns | Pistols | ..
}
