package middlewares

import (
	"oauth2-authorization/config"
	"oauth2-authorization/models"
	"oauth2-authorization/utility"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() (err error) {
	dbinfo := config.GetDBInfo()

	db, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	if err != nil {
		utility.Log(models.LogLevelError, nil, nil, "Database Init Failed!: ", err.Error())
		return err
	}

	utility.Log(models.LogLevelDebug, nil, nil, "Database Init Complete! ", dbinfo)
	return
}
