package business

import "time"

type ObjectMetadata struct {
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	CreatedBy  int       `json:"created_by" gorm:"created_by"`
	ModifiedAt time.Time `json:"modified_at" gorm:"modified_at"`
	ModifiedBy int       `json:"modified_by" gorm:"modified_by"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"deleted_at"`
	DeletedBy  int       `json:"deleted_by" gorm:"deleted_by"`
	Version    int       `json:"version" gorm:"version"`
}
