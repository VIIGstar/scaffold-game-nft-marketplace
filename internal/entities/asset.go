package entities

import (
	base_entity "scaffold-game-nft-marketplace/pkg/base-entity"
	base_type "scaffold-game-nft-marketplace/pkg/base-type"
)

type AssetLevel base_type.DefinedStringType

const (
	Basic   = AssetLevel("basic")
	Rare    = AssetLevel("rare")
	Epic    = AssetLevel("epic")
	Legends = AssetLevel("legends")
)

//Asset items sync from game whenever user first list onto Marketplace through minting NFT
type Asset struct {
	base_entity.Base
	base_entity.Reference
	// Path to metadata of item
	URI string `json:"uri" gorm:"type:varchar"`
	// Attributes of item affect in game
	Attributes string `json:"attributes" gorm:"type:varchar"`
	// Image/3D model to display item
	Image string `json:"image" gorm:"type:varchar"`
	// Includes: Basic | Rare | Epic | Legends
	Level AssetLevel `json:"level" gorm:"type:varchar"`
	// References
	CategoryCode string   `json:"category_code" gorm:"size:64"`
	Category     Category `json:"category" gorm:"foreignKey:CategoryCode;references:Code"`
}
