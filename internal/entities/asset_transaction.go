package entities

import (
	base_entity "scaffold-game-nft-marketplace/pkg/base-entity"
	base_type "scaffold-game-nft-marketplace/pkg/base-type"
)

type TransactionType base_type.DefinedStringType

const (
	Sync     = TransactionType("sync")
	Fallback = TransactionType("fallback")
)

//AssetTransaction sync from network or insert after execute call JSON RPC
type AssetTransaction struct {
	base_entity.Base
	// TransactionHash in block
	TxHash string `json:"tx_hash" gorm:"type:varchar;uniqueIndex"`
	// (Optional) index of transaction in list block's transactions
	TxIndex int64 `json:"tx_index"`
	// Hash of block which transaction belong to
	BlockHash string `json:"block_hash" gorm:"type:varchar;index"`
	// Includes: Sync | Fallback
	Type    TransactionType `json:"type" gorm:"size:16;NOT NULL"`
	AssetID int64           `json:"asset_id"`
	Asset   Asset           `json:"asset"`
}
