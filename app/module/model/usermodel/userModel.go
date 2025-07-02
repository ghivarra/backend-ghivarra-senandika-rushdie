package userModel

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"primaryKey;autoIncrement"`
	Username   string         `gorm:"uniqueIndex;not null"`
	Password   string         `gorm:"not null"`
	Name       string         `gorm:"not null"`
	Email      string         `gorm:"uniqueIndex;not null"`
	UserRoleID uint           `gorm:"index;not null"`
	IsActive   uint           `gorm:"index;default:1"`
	CreatedAt  time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
