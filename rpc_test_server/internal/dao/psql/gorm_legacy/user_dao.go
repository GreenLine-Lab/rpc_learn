package gorm_legacy_dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"rpc_learn/rpc_test_server/internal/models"
	"time"
)

type UserDaoGormLegacy struct {
	db *gorm.DB
}

func NewUserDaoGormLegacy(connect string) (*UserDaoGormLegacy, error) {
	db, err := gorm.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxIdleTime(time.Minute * 5)

	return &UserDaoGormLegacy{db: db}, nil
}

func (dao *UserDaoGormLegacy) GORM() *gorm.DB {
	return dao.db
}

func (dao *UserDaoGormLegacy) Close() error {
	return dao.db.Close()
}

func (dao *UserDaoGormLegacy) Create(user *models.User) error {
	if user == nil {
		return errors.New("empty user pointer")
	}

	return dao.GORM().Table(user.Table()).Create(&user).Error
}

func (dao *UserDaoGormLegacy) Reed(user *models.User) error {
	if user == nil {
		return errors.New("empty user pointer")
	}

	return dao.GORM().Table(user.Table()).First(&user, user.GetFilter()).Error
}

func (dao *UserDaoGormLegacy) Update(user *models.User) error {
	panic("implement me")
}

func (dao *UserDaoGormLegacy) Delete(user *models.User) error {
	panic("implement me")
}
