package model

import "time"

// UserStat model
type UserStat struct {
	UserID       uint8     `gorm:"column:user_id;type:mediumint unsigned"`
	LastWeight   float32   `gorm:"column:last_weight;type:float unsigned"`
	Height       float32   `gorm:"column:height;type:float unsigned"`
	InjuredSince time.Time `gorm:"column:injured_since;type:date"`
	InjuredUntil time.Time `gorm:"column:injured_until;type:date"`
	User         User      `gorm:"foreignKey:UserID;references:ID"`
}
