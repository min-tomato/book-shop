package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; index:idx_uuid; primaryKey; autoIncrement; comment:'primary key is id'"`
	RoleName string `gorm:"column:role_name"`
	RoleNote string `gorm:"column:role_note; type:text"`
}

func (u *Role) TableName() string {
	return "db_role"
}
