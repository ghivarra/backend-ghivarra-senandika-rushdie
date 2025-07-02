package userRoleModel

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
