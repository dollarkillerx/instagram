package simple

import (
	cfg "github.com/dollarkillerx/common/pkg/config"
	"github.com/dollarkillerx/graphql_template/internal/pkg/models"
	"github.com/dollarkillerx/graphql_template/internal/utils"
	"gorm.io/gorm"

	"sync"
)

type Simple struct {
	db *gorm.DB

	inventoryMu sync.Mutex
}

func NewSimple(conf cfg.PostgresConfiguration) (*Simple, error) {
	sql, err := utils.InitPgSQL(conf)
	if err != nil {
		return nil, err
	}

	sql.AutoMigrate(
		&models.User{},
	)

	return &Simple{
		db: sql,
	}, nil
}

func (s *Simple) DB() *gorm.DB {
	return s.db
}
