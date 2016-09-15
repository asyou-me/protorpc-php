package main

import (
	"C"
	"time"
)

var (
	verison = "0.1"
)

var poolMap = map[string]*Pool{}

//export Protorpc
func Protorpc(address string, max int, timeOut int) *C.char {
	var err error
	if poolMap[address] != nil {
		return C.CString("ok")
	}

	pool, err := NewPool(func() (*Client, error) {
		cli, err := DialTimeout("tcp", address, time.Duration(timeOut)*time.Millisecond)
		if err != nil {
			return nil, err
		}
		return cli, nil
	}, func(c *Client, t time.Time) error {
		return nil
	}, max)
	if err != nil {
		return C.CString("新建连接错误:" + err.Error())
	}
	poolMap[address] = pool
	return C.CString("ok")
}

//export ProtorpcCall
func ProtorpcCall(address string, serviceMethod string, data string) *C.char {
	pool, ok := poolMap[address]
	if !ok {
		return C.CString("无法找到可用连接")
	}
	replay := make([]byte, 0, 1500)
	err := pool.Call(serviceMethod, []byte(data), &replay)
	if err != nil {
		return C.CString("调用失败:" + err.Error())
	}
	toC := C.CString(string(replay))
	return toC
}

func main() {}
