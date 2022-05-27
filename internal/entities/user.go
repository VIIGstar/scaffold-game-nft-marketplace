package entities

import base_entity "scaffold-game-nft-marketplace/pkg/base-entity"

type User struct {
	Investor
	base_entity.Reference
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
