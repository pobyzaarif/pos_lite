package business

import (
	"context"
	"time"
)

type ObjectMetadata struct {
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	CreatedBy  int       `json:"created_by" gorm:"created_by"`
	ModifiedAt time.Time `json:"modified_at" gorm:"modified_at"`
	ModifiedBy int       `json:"modified_by" gorm:"modified_by"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"deleted_at"`
	DeletedBy  int       `json:"deleted_by" gorm:"deleted_by"`
	Version    int       `json:"version" gorm:"version"`
}

type InternalContext struct {
	TrackerID string
}

func NewInternalContext(trackerID string) InternalContext {
	return InternalContext{
		TrackerID: trackerID,
	}
}

func (ic InternalContext) ToContext() context.Context {
	ctx := context.WithValue(context.Background(), "tracker_id", ic.TrackerID)
	return ctx
}
