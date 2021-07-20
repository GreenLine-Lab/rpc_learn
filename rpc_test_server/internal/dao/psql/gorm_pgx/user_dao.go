package gorm_pgx_dao

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rpc_learn/rpc_test_server/internal/models"
	"time"
)

type UserDaoGorm struct {
	db *gorm.DB
}

func NewUserDaoGorm(dsn string) (*UserDaoGorm, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxIdleTime(time.Minute * 5)

	return &UserDaoGorm{db: db}, nil
}

func (dao *UserDaoGorm) GORM() *gorm.DB {
	return dao.db
}

func (dao *UserDaoGorm) Close() error {
	db, err := dao.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (dao *UserDaoGorm) Create(user *models.User) error {
	if user == nil {
		return errors.New("empty user pointer")
	}

	return dao.GORM().Table(user.Table()).Create(&user).Error
}

func (dao *UserDaoGorm) Reed(user *models.User, where ...interface{}) error {
	if user == nil {
		return errors.New("empty user pointer")
	}

	if where == nil {
		return errors.New("empty where parameters")
	}

	return dao.GORM().Table(user.Table()).First(&user, where...).Error
}

func (dao *UserDaoGorm) Update(user *models.User, where ...interface{}) error {
	panic("implement me")
}

func (dao *UserDaoGorm) Delete(user *models.User, where ...interface{}) error {
	panic("implement me")
}
