package model

import "time"

// @program:     goods_servies
// @file:        base.go.go
// @author:      bowen
// @create:      2023-01-07 17:10
// @description: 7个基础字段结构体

type BaseModel struct {
	ID       uint `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
	CreateBy string
	UpdateBy string
	Version  int16
	isDel    int8 `gorm:"index"`
}
