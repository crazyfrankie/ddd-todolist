package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/crazyfrankie/ddd-todolist/backend/conf"
)

func main() {
	db := connectDB(conf.GetConf().MySQL.DSN)

	genTask(db)
	genUser(db)
}

func genTask(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "domain/task/internal/dal/query",
		ModelPkgPath: "domain/task/internal/dal/model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateModel("task"))

	g.Execute()
}

func genUser(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "domain/user/internal/dal/query",
		ModelPkgPath: "domain/user/internal/dal/model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateModel("user"))

	g.Execute()
}

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}
