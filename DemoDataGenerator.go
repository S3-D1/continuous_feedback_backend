package main

import (
	"fmt"
	"github.com/S3-D1/continuous_feedback_backend/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var demoDB *gorm.DB
var demoErr error

func createExampleData() {
	initDemoDb()
	deleteAllData()
	generateExampleData()
}

func deleteAllData() {
	demoDB.Unscoped().Delete(models.Group{})
	demoDB.Unscoped().Delete(models.SingleFeedback{})
	demoDB.Unscoped().Delete(models.User{})
}

func initDemoDb() {
	host := "0.0.0.0"
	port := "5432"
	schema := "cf"
	user := "postgres"
	pw := "postgres"
	demoDB, demoErr = gorm.Open(
		"postgres",
		"host="+host+" port="+port+" user="+user+
			" dbname="+schema+" sslmode=disable password="+pw)

	if demoErr != nil {
		fmt.Println(demoErr)
		panic("failed to connect database")
	}

	demoDB.AutoMigrate(&models.SingleFeedback{}, &models.Group{}, &models.User{})
}

func generateExampleData() {
	g1 := models.Group{
		Name: "all",
	}
	demoDB.Create(&g1)
	u1 := models.User{
		Nickname: "u1",
		Email:    "u1@e.f",
		Password: "no_pw",
	}
	u2 := models.User{
		Nickname: "u2",
		Email:    "u2@e.f",
		Password: "no_pw",
	}
	u3 := models.User{
		Nickname: "u3",
		Email:    "u3@e.f",
		Password: "no_pw",
	}
	demoDB.Create(&u1)
	demoDB.Create(&u2)
	demoDB.Create(&u3)
	f1 := models.SingleFeedback{
		Value:     7,
		Author:    u1,
		Recipient: u2,
	}
	f2 := models.SingleFeedback{
		Value:     3,
		Author:    u1,
		Recipient: u3,
	}
	f3 := models.SingleFeedback{
		Value:     5,
		Author:    u3,
		Recipient: u2,
	}
	demoDB.Create(&f1)
	demoDB.Create(&f2)
	demoDB.Create(&f3)
	g2 := models.Group{
		Name:  "g2",
		Users: []models.User{u2, u3},
	}
	demoDB.Create(&g2)
}
