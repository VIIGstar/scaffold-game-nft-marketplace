package base_entity

import "time"

type Reference struct {
	RefId        int64     `json:"ref_id" gorm:"uniqueIndex"`
	RefType      string    `json:"ref_type" gorm:"type:varchar"`
	RefCreatedAt time.Time `json:"ref_created_at"`
	LinkedAt     time.Time `json:"linked_at"`
}
