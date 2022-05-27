package entities

import (
	base_entity "scaffold-game-nft-marketplace/pkg/base-entity"
	base_type "scaffold-game-nft-marketplace/pkg/base-type"
)

type CategoryCode base_type.DefinedStringType

const (
	Character   = CategoryCode("character")
	Weapon      = CategoryCode("weapon")
	Accessories = CategoryCode("accessories")
	Box         = CategoryCode("box")
)

// Category will insert whenever has new game-items/game-assets with un-recognized category
// its seem to never be updated
type Category struct {
	base_entity.Base
	base_entity.Reference
	Name string `json:"name" gorm:"size:256"`
	// Includes: Character | Weapon | Accessories | Box ...
	Code CategoryCode `json:"code" gorm:"size:64;unique"`
}
