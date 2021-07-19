package test

import (
	"fmt"
	"github.com/google/uuid"
	"rpc_learn/lib/cfg"
	"rpc_learn/rpc_test_server/internal/dao"
	"rpc_learn/rpc_test_server/internal/models"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestDao(t *testing.T) {

	config := cfg.Config{
		SqlHost: "192.168.20.114",
		SqlPort: "5444",
		SqlUser: "postgres",
		SqlPass: "postgres",
		SqlDb:   "test",
	}

	t.Run("UserDaoGormPgx", func(t *testing.T) {
		config.SqlType = "gorm_pgx"

		uDao, err := dao.NewUserDao(&config)
		if err != nil {
			t.Fatalf("Unable create NewUserDao: %s", err.Error())
		}

		queryCreate(t, uDao, config.SqlType)
		queryRead(t, uDao, config.SqlType)
		uDao.Close()
	})

	time.Sleep(time.Second * 5)

	t.Run("UserDaoGormLegacy", func(t *testing.T) {
		config.SqlType = "gorm_legacy"

		uDao, err := dao.NewUserDao(&config)
		if err != nil {
			t.Fatalf("Unable create NewUserDao: %s", err.Error())
		}

		queryCreate(t, uDao, config.SqlType)
		queryRead(t, uDao, config.SqlType)
		uDao.Close()
	})

	time.Sleep(time.Second * 5)

	t.Run("UserDaoPgx", func(t *testing.T) {
		config.SqlType = "pgx"

		uDao, err := dao.NewUserDao(&config)
		if err != nil {
			t.Fatalf("Unable create NewUserDao: %s", err.Error())
		}

		queryCreate(t, uDao, config.SqlType)
		queryRead(t, uDao, config.SqlType)
		uDao.Close()
	})

	time.Sleep(time.Second * 5)

	t.Run("UserDaoPgxPool", func(t *testing.T) {
		config.SqlType = "pgx_pool"

		uDao, err := dao.NewUserDao(&config)
		if err != nil {
			t.Fatalf("Unable create NewUserDao: %s", err.Error())
		}

		queryCreate(t, uDao, config.SqlType)
		queryRead(t, uDao, config.SqlType)
		uDao.Close()
	})

}

type metric struct {
	mux              *sync.Mutex
	typeDao          string
	typeQuery        string
	totalNanoseconds int64
	countQuery       int
	min              int64
	max              int64
	avg              int64
}

func (m *metric) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("%s %s (%d query):", m.typeDao, m.typeQuery, m.countQuery))
	lines = append(lines, fmt.Sprintf("	Min: %.5f Milliseconds", float64(m.min)/float64(1000000000)))
	lines = append(lines, fmt.Sprintf("	Max: %.5f Milliseconds", float64(m.max)/float64(1000000000)))
	lines = append(lines, fmt.Sprintf("	Avg: %.5f Milliseconds", float64(m.totalNanoseconds/int64(m.countQuery))/float64(1000000000)))

	return strings.Join(lines, "\n")
}

func queryCreate(t *testing.T, uDao dao.UserDao, typeDao string) {
	const queryCount int = 1000
	wg := sync.WaitGroup{}
	m := metric{
		mux:       new(sync.Mutex),
		typeDao:   typeDao,
		typeQuery: "create",
		min:       100000000000000,
	}

	startCh := make(chan bool)
	for queryIdx := 0; queryIdx < queryCount; queryIdx++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m *metric) {
			defer wg.Done()

			user := models.User{
				Uuid:      uuid.New().String(),
				Password:  "password",
				Rule:      0,
				CreateAt:  time.Now().Unix(),
				BlockedAt: 0,
				UpdateAt:  0,
			}

			user.Login = fmt.Sprintf("%s%d", user.Uuid, time.Now().UnixNano())
			user.NickName = user.Login

			<-startCh

			now := time.Now()
			if err := uDao.Create(&user); err != nil {
				return
			}

			difference := time.Now().Sub(now).Nanoseconds()
			m.mux.Lock()
			if m.min > difference {
				m.min = difference
			}

			if m.max < difference {
				m.max = difference
			}

			m.countQuery += 1
			m.totalNanoseconds += difference
			m.mux.Unlock()
		}(&wg, &m)
	}

	close(startCh)
	wg.Wait()
	t.Log(m.String())
	return
}

func queryRead(t *testing.T, uDao dao.UserDao, typeDao string) {
	const queryCount int = 1000
	wg := sync.WaitGroup{}
	m := metric{
		mux:       new(sync.Mutex),
		typeDao:   typeDao,
		typeQuery: "read",
		min:       100000000000000,
	}

	startCh := make(chan bool)
	for queryIdx := 0; queryIdx < queryCount; queryIdx++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m *metric) {
			defer wg.Done()

			user := models.User{
				Login: "5502f4b3-ca40-4169-8a43-d840abd916e91626688379438175808",
			}

			<-startCh

			now := time.Now()
			if err := uDao.Reed(&user); err != nil {
				return
			}

			difference := time.Now().Sub(now).Nanoseconds()

			if len(user.Uuid) == 0 {
				return
			}

			m.mux.Lock()
			if m.min > difference {
				m.min = difference
			}

			if m.max < difference {
				m.max = difference
			}

			m.countQuery += 1
			m.totalNanoseconds += difference
			m.mux.Unlock()
		}(&wg, &m)
	}

	close(startCh)
	wg.Wait()
	t.Log(m.String())
	return
}
