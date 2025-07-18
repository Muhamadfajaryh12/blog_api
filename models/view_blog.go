package models

import "time"

type ViewBlog struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	BlogID    uint64 `gorm:"index"`
	View      int64
}