package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/yuhando/userprofile/api/models"
)

var users = []models.User{
	models.User{
		Name:     "Yuhando",
		Email:    "yuhando74@gmail.com",
		Address:  "Jakarta",
		Password: "password",
	},
	models.User{
		Name:     "Charlilio",
		Email:    "Charlilio@gmail.com",
		Address:  "Bandung",
		Password: "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}
