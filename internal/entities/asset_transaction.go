package entities

import (
	base_entity "scaffold-game-nft-marketplace/pkg/base-entity"
	base_type "scaffold-game-nft-marketplace/pkg/base-type"
)

type TransactionType base_type.DefinedStringType
type TransactionStatus base_type.DefinedStringType

const (
	Sync     = TransactionType("sync")
	Fallback = TransactionType("fallback")

	Synced     = TransactionStatus("synced")
	Processing = TransactionStatus("processing")
	Verified   = TransactionStatus("verified")
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
	Type TransactionType `json:"type" gorm:"size:16;NOT NULL"`
	// Includes: Synced (right after reading from block) | Processing (if exists data to update business) | Verified (Nothing left to do)
	Status  TransactionStatus `json:"status"`
	AssetID int64             `json:"asset_id"`
	Asset   Asset             `json:"asset"`
}
