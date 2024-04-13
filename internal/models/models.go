// internal/models/models.go
package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique_index"`
    Password string
    Token    string `gorm:"-"`
}

type Item struct {
    gorm.Model
    Name  string
    Price float64
}

type Cart struct {
    gorm.Model
    UserID uint
    Items  []Item `gorm:"many2many:cart_items;"`
}

type Order struct {
    gorm.Model
    CartID uint
}
