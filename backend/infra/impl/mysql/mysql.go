package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/conf"
)

func New() (*gorm.DB, error) {
	dsn := conf.GetConf().MySQL.DSN
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("mysql open, dsn: %s, err: %w", dsn, err)
	}

	return db, nil
}
