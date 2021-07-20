package pgx_pool_dao

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"rpc_learn/rpc_test_server/internal/models"
)

type UserDaoPgxPool struct {
	db *pgxpool.Pool
}

func NewUserDaoPgxPool(connect string) (*UserDaoPgxPool, error) {
	pool, err := pgxpool.Connect(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &UserDaoPgxPool{db: pool}, nil
}

func (dao *UserDaoPgxPool) Conn() (*pgxpool.Conn, error) {
	conn, err := dao.db.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (dao *UserDaoPgxPool) Close() error {
	return nil
}

func (dao *UserDaoPgxPool) Create(u *models.User) error {
	if u == nil {
		return errors.New("empty user pointer")
	}

	const sqlstr = `INSERT INTO public.user (` +
		`uuid, nick_name, login, password, rule, blocked, create_at, blocked_at, update_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) RETURNING id`

	conn, err := dao.Conn()
	if err != nil {
		return err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(), sqlstr, u.Uuid, u.NickName, u.Login, u.Password, u.Rule, u.Blocked, u.CreateAt, u.BlockedAt, u.UpdateAt).Scan(&u.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDaoPgxPool) Reed(u *models.User, where ...interface{}) error {
	if u == nil {
		return errors.New("empty user pointer")
	}

	sqlstr := `SELECT * FROM public.user `
	filter := u.GetFilter()
	if len(filter) == 0 {
		return errors.New("empty where parameters")
	}

	sqlstr += "WHERE " + filter

	conn, err := dao.Conn()
	if err != nil {
		return err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(), sqlstr).Scan(&u.Id, &u.Uuid, &u.NickName, &u.Login, &u.Password, &u.Rule, &u.Blocked, &u.CreateAt, &u.BlockedAt, &u.UpdateAt)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDaoPgxPool) Update(user *models.User, where ...interface{}) error {
	panic("implement me")
}

func (dao *UserDaoPgxPool) Delete(user *models.User, where ...interface{}) error {
	panic("implement me")
}
