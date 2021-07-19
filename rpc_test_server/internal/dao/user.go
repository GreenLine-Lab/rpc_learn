package dao

import (
	"errors"
	"fmt"
	"rpc_learn/lib/cfg"
	"rpc_learn/rpc_test_server/internal/dao/psql/gorm_legacy"
	"rpc_learn/rpc_test_server/internal/dao/psql/gorm_pgx"
	pgx_dao "rpc_learn/rpc_test_server/internal/dao/psql/pgx"
	pgx_pool_dao "rpc_learn/rpc_test_server/internal/dao/psql/pgx_pool"
	"rpc_learn/rpc_test_server/internal/models"
)

type UserDao interface {
	Create(user *models.User) error
	Reed(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error

	Close() error
}

func NewUserDao(config *cfg.Config) (UserDao, error) {
	var dao UserDao

	if len(config.SqlType) == 0 {
		return nil, errors.New("empty parameter SqlType")
	}

	switch config.SqlType {
	case "gorm_pgx":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.SqlHost,
			config.SqlUser,
			config.SqlPass,
			config.SqlDb,
			config.SqlPort)

		daoGorm, err := gorm_pgx_dao.NewUserDaoGorm(dsn)
		if err != nil {
			return nil, err
		}

		dao = daoGorm
	case "gorm_legacy":
		connect := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.SqlUser,
			config.SqlPass,
			config.SqlHost,
			config.SqlPort,
			config.SqlDb)

		daoGorm, err := gorm_legacy_dao.NewUserDaoGormLegacy(connect)
		if err != nil {
			return nil, err
		}

		dao = daoGorm
	case "pgx":
		connect := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.SqlUser,
			config.SqlPass,
			config.SqlHost,
			config.SqlPort,
			config.SqlDb)

		daoPgx, err := pgx_dao.NewUserDaoPgx(connect)
		if err != nil {
			return nil, err
		}

		dao = daoPgx
	case "pgx_pool":
		connect := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.SqlUser,
			config.SqlPass,
			config.SqlHost,
			config.SqlPort,
			config.SqlDb)

		daoPgx, err := pgx_pool_dao.NewUserDaoPgxPool(connect)
		if err != nil {
			return nil, err
		}

		dao = daoPgx
	default:
		return nil, errors.New("unknown SqlType")
	}

	return dao, nil
}
