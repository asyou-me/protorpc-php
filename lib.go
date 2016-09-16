package main

import (
	"C"
	"time"

	"github.com/asyou-me/protorpc-php/rpc"
)

var (
	verison = "0.1"
)

var poolMap = map[string]*rpc.Pool{}

//export Protorpc
func Protorpc(address string, max int, timeOut int) (toC *C.char) {
	var err error
	if poolMap[address] != nil {
		toC = C.CString("ok")
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
		toC = C.CString("新建连接错误:" + err.Error())
		return
	}
	poolMap[address] = pool
	toC = C.CString("ok")
	return
}

//export ProtorpcCall
func ProtorpcCall(address string, serviceMethod string, data string) (toC *C.char) {
	pool, ok := poolMap[address]
	if !ok {
		toC = C.CString("无法找到可用连接")
		return
	}
	replay := make([]byte, 0, 1500)
	err := pool.Call(serviceMethod, []byte(data), &replay)
	if err != nil {
		toC = C.CString("调用失败:" + err.Error())
		return
	}
	toC = C.CString(string(replay))
	return
}

func main() {}
