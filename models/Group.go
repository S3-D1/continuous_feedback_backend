package models

type Group struct {
	ID    uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name  string `json:"name"`
	Users []User `gorm:"many2many:members;association_foreignkey:ID;foreignkey:ID" json:"users"`
}
