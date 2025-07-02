package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	UserID    uint           `gorm:"index;not null;foreignKey:ID"`
	User      User           `gorm:"foreignKey:UserID"`
	ProductID uint           `gorm:"index;not null"`
	Product   Product        `gorm:"foreignKey:ProductID"`
	Quantity  uint           `gorm:"default:1"`
	CreatedAt time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Invoice struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Price     uint           `gorm:"not null"`
	Details   string         `gorm:"not null"`
	UserID    uint           `gorm:"index;not null;foreignKey:ID"`
	User      User           `gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type InvoiceOrderList struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	InvoiceID uint    `gorm:"index;not null;foreignKey:ID"`
	Invoice   Invoice `gorm:"foreignKey:InvoiceID"`
	OrderID   uint    `gorm:"index;not null;foreignKey:ID"`
	Order     Order   `gorm:"foreignKey:OrderID"`
}

type Order struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	BuyerID   uint           `gorm:"index;not null;foreignKey:ID"`
	Buyer     User           `gorm:"foreignKey:BuyerID"`
	SellerID  uint           `gorm:"index;not null;foreignKey:ID"`
	Seller    User           `gorm:"foreignKey:SellerID"`
	ProductID uint           `gorm:"index;not null;foreignKey:ID"`
	Product   Product        `gorm:"foreignKey:ProductID"`
	Price     uint           `gorm:"not null"`
	Quantity  uint           `gorm:"not null"`
	CreatedAt time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Slug        string `gorm:"uniqueIndex;not null"`
	Name        string `gorm:"not null"`
	Description string
	Photo       string         `gorm:"not null"`
	Price       uint           `gorm:"not null"`
	Stock       uint           `gorm:"not null"`
	UserID      uint           `gorm:"index;not null"`
	User        User           `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type User struct {
	ID         uint           `gorm:"primaryKey;autoIncrement"`
	Username   string         `gorm:"uniqueIndex;not null"`
	Password   string         `gorm:"not null"`
	Name       string         `gorm:"not null"`
	Email      string         `gorm:"uniqueIndex;not null"`
	UserRoleID uint           `gorm:"index;not null"`
	UserRole   UserRole       `gorm:"foreignKey:UserRoleID"`
	IsActive   uint           `gorm:"index;default:1"`
	CreatedAt  time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type UserModule struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserRole struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserRoleModuleList struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	UserRoleID   uint           `gorm:"index;not null"`
	UserRole     UserRole       `gorm:"foreignKey:UserRoleID"`
	UserModuleID uint           `gorm:"index;not null"`
	UserModule   UserModule     `gorm:"foreignKey:UserModuleID"`
	CreatedAt    time.Time      `gorm:"<-:create;not null;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"<-:update;not null;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
