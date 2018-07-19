package model

import (
	"time"
	"sync"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	//DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	Phone  string `json:"phone,omitempty"`
	Password  string `json:"-"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"-"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
