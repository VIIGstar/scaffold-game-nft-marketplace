package query_params

import "scaffold-game-nft-marketplace/pkg/database"

type GetUserParams struct {
	database.CommonQueryParams
	Address string `json:"address"`
}
