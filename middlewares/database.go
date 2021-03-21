package middlewares

import (
	"database/sql"
	"oauth2-authorization/config"
	"oauth2-authorization/models"
	"oauth2-authorization/utility"

	_ "github.com/lib/pq"
)

func InitDatabase() (err error) {
	dbinfo := config.GetDBInfo()

	db, err := sql.Open(config.GetDBEngine(), dbinfo)
	if err != nil {
		utility.Log(models.LogLevelError, nil, nil, "Database Init Failed!: ", err.Error())
		return err
	}
	defer db.Close()

	utility.Log(models.LogLevelDebug, nil, nil, "Database Init Complete! ", dbinfo)
	return
}
