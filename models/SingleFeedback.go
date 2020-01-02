package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type SingleFeedback struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Value       int       `gorm:"not null" json:"value"`
	Author      User      `gorm:"not null;foreignkey:AuthorID;association_foreignkey:ID" json:"author"`
	AuthorID    uint      `json:"-"`
	Recipient   User      `gorm:"not null;foreignkey:RecipientID;association_foreignkey:ID" json:"recipient"`
	RecipientID uint      `json:"-"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (p *SingleFeedback) Prepare() {
	p.ID = 0
	p.Author = User{}
	p.Recipient = User{}
	p.CreatedAt = time.Now()
}

func (p *SingleFeedback) Validate() error {

	if p.Value < 0 && p.Value > 10 {
		return errors.New("SingleFeedback Value is out of range!")
	}
	if p.Author.ID < 1 {
		return errors.New("Required Author")
	}
	if p.Recipient.ID < 1 {
		return errors.New("Required SingleFeedback Recipient")
	}
	return nil
}

func (f *SingleFeedback) SaveFeedback(db *gorm.DB) *SingleFeedback {
	db.Create(&f)
	return f
}
