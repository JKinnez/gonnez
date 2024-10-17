package dbsetup

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JKinnez/gonnez/environment"
	"github.com/samber/lo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	ErrNoDB      = "no database connection"
	ErrCanDelete = "can't delete data in production you maniac"
	emptyString  = ""
)

type DBOptions struct {
	Config *gorm.Config
	DSN    string
	Models []any
}

func NewDatabaseConnection(options DBOptions) (orm *gorm.DB, err error) {
	orm, err = options.open()
	if err != nil {
		return
	}

	err = options.schema(orm)
	return
}

func CustomLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      false,
		},
	)
}

func ClearData(orm *gorm.DB, models ...any) (err error) {
	if orm == nil {
		err = fmt.Errorf(ErrNoDB)
		return
	}

	if environment.IsProduction() {
		return fmt.Errorf(ErrCanDelete)
	}

	for _, m := range lo.Reverse(models) {
		err = orm.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&m).Error
		if err != nil {
			return
		}
	}
	return
}

func (dbo *DBOptions) open() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dbo.DSN), dbo.Config)

	return
}

func (dbo *DBOptions) schema(orm *gorm.DB) (err error) {
	for _, model := range dbo.Models {
		err = createTable(orm, model)
		if err != nil {
			return
		}

		err = orm.AutoMigrate(model)
		if err != nil {
			return
		}
	}

	return
}

func createTable(orm *gorm.DB, model any) (err error) {
	if orm.Migrator().HasTable(model) {
		return
	}

	err = orm.Migrator().CreateTable(model)
	return
}
