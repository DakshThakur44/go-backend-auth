package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = "usr_" + uuid.NewString()
	}
	return nil
}

type Session struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"not null;index" json:"user_id"`
	UserAgent string    `json:"user_agent"`
	ClientIP  string    `json:"client_ip"`
	IsRevoked bool      `gorm:"default:false" json:"is_revoked"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = "sess_" + uuid.NewString()
	}
	return nil
}

// RefreshToken entity table
type RefreshToken struct {
	Token     string    `gorm:"primaryKey" json:"token"`
	SessionID string    `gorm:"not null;index" json:"session_id"`
	UserID    string    `gorm:"not null;index" json:"user_id"`
	IsUsed    bool      `gorm:"default:false" json:"is_used"`
	IsRevoked bool      `gorm:"default:false" json:"is_revoked"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// --- Request DTOs ---

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
