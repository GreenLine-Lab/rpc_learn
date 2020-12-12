package test

import (
	"context"
	"fmt"
	"rpc-learn/rpc-test-server/pb"
	"sync"
	"testing"
)

func TestHelloSync(t *testing.T) {

	req := pb.ReqTest{
		Name: "Vladimir",
	}

	c := Client{}
	c.Client()
	defer c.CloseConnection()
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			rpl, err := c.Client().TestSyncHello(context.Background(), &req)
			if err != nil {
				t.Fatalf("TestSyncHello return: %s", err.Error())
			}

			fmt.Printf("Rpl: %+v", rpl)
			wg.Done()
		}(&wg)
	}

	wg.Wait()
}
