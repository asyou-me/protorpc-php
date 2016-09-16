package main

import (
	"encoding/base64"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/asyou-me/protorpc-php/rpc"
)

var (
	verison = "0.1"
)

var poolMap = map[string]*rpc.Pool{}

func Protorpc(address string, max int, timeOut int) (toC string) {
	var err error
	if poolMap[address] != nil {
		toC = "ok"
		return
	}

	pool, err := rpc.NewPool(func() (*rpc.Client, error) {
		cli, err := rpc.DialTimeout("tcp", address, time.Duration(timeOut)*time.Millisecond)
		if err != nil {
			return nil, err
		}
		return cli, nil
	}, func(c *rpc.Client, t time.Time) error {
		return nil
	}, max)
	if err != nil {
		toC = "新建连接错误:" + err.Error()
		return
	}
	poolMap[address] = pool
	toC = "ok"
	return
}

func ProtorpcCall(address string, serviceMethod string, data string) (toC string) {
	pool, ok := poolMap[address]
	if !ok {
		toC = "无法找到可用连接"
		return
	}
	replay := make([]byte, 0, 1500)
	err := pool.Call(serviceMethod, []byte(data), &replay)
	if err != nil {
		toC = "调用失败:" + err.Error()
		return
	}
	toC = string(replay)
	return
}

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	Protorpc("127.0.0.1:30015", 10, 11)
	data, _ := base64.StdEncoding.DecodeString("CAEQAg==")
	for {
		ProtorpcCall("127.0.0.1:30015", "TestHandler.Test", string(data))
	}
}
