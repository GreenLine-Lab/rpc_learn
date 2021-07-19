package pgx_dao

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"rpc_learn/rpc_test_server/internal/models"
)

type UserDaoPgx struct {
	connect string
}

func NewUserDaoPgx(connect string) (*UserDaoPgx, error) {
	conn, err := pgx.Connect(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &UserDaoPgx{connect: connect}, nil
}

func (dao *UserDaoPgx) Conn() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dao.connect)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (dao *UserDaoPgx) Close() error {
	return nil
}

func (dao *UserDaoPgx) Create(u *models.User) error {
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
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), sqlstr, u.Uuid, u.NickName, u.Login, u.Password, u.Rule, u.Blocked, u.CreateAt, u.BlockedAt, u.UpdateAt).Scan(&u.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDaoPgx) Reed(u *models.User) error {
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
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), sqlstr).Scan(&u.Id, &u.Uuid, &u.NickName, &u.Login, &u.Password, &u.Rule, &u.Blocked, &u.CreateAt, &u.BlockedAt, &u.UpdateAt)
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDaoPgx) Update(user *models.User) error {
	panic("implement me")
}

func (dao *UserDaoPgx) Delete(user *models.User) error {
	panic("implement me")
}
