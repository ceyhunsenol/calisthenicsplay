package data

import (
	_ "gorm.io/gorm"
)

type User struct {
	BaseModel
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	Code       string      `json:"code"`
	Privileges []Privilege `gorm:"many2many:role_privileges"`
}

type Privilege struct {
	BaseModel
	Code          string `json:"code"`
	Description   string `json:"description"`
	EndpointsJoin string `json:"end_points_join"`
	Roles         []Role `gorm:"many2many:role_privileges"`
}
