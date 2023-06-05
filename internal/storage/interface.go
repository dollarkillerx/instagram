package storage

import (
	"github.com/dollarkillerx/graphql_template/internal/generated"
	"github.com/dollarkillerx/graphql_template/internal/pkg/models"
	"gorm.io/gorm"
)

type Interface interface {
	DB() *gorm.DB

	GetUserByAccount(account string) (*models.User, error)
	AccountRegistry(account string, name string, password string, role generated.Role) error
}
