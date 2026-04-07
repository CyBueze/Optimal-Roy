package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Business struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    Name      string    `gorm:"not null"`
    CreatedAt time.Time
    Branches  []Branch
}

type Branch struct {
    ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
    BusinessID uuid.UUID `gorm:"type:uuid;not null"`
    Name       string    `gorm:"not null"`
    Address    string
    CreatedAt  time.Time
    Users      []User
}

type User struct {
    ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
    BusinessID uuid.UUID `gorm:"type:uuid;not null"`
    BranchID   *uuid.UUID `gorm:"type:uuid"` // nullable — director has no branch
    Name       string    `gorm:"not null"`
    Email      string    `gorm:"uniqueIndex;not null"`
    Password   string    `gorm:"not null"` // bcrypt hash
    Role       string    `gorm:"not null"` // "director" | "manager" | "cashier"
    CreatedAt  time.Time
}

// BeforeCreate sets UUID for any model that needs it
func (b *Business) BeforeCreate(tx *gorm.DB) error {
    b.ID = uuid.New()
    return nil
}

func (br *Branch) BeforeCreate(tx *gorm.DB) error {
    br.ID = uuid.New()
    return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
    u.ID = uuid.New()
    return nil
}