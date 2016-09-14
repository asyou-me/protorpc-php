package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/rpc"
	"time"
)

// Client 自定义 rpc 连接
type Client struct {
	*rpc.Client
}

// Dial 创建一个 rpc 连接
func Dial(network, address string, auth ...string) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn, auth...), err
}

// DialTimeout 创建一个设定超时时间的 rpc 连接
func DialTimeout(network, address string, timeout time.Duration, auth ...string) (*Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return NewClient(conn, auth...), err
}

// NewClient 创建新的客户端
func NewClient(conn io.ReadWriteCloser, auth ...string) *Client {
	client := &Client{}
	client.Client = rpc.NewClientWithCodec(NewClientCodec(conn, auth...))
	return client
}

var poolMap = map[string]*Pool{}

// Protorpc 新建一个 rpc 连接
func Protorpc(address string, max int, timeOut int) string {
	var err error
	if poolMap[address] != nil {
		return "ok"
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
		return "新建连接错误:" + err.Error()
	}
	poolMap[address] = pool
	return "ok"
}

// ProtorpcCall 请求服务器
func ProtorpcCall(address string, serviceMethod string, data string) string {
	pool, ok := poolMap[address]
	if !ok {
		return "无法找到可用连接"
	}
	fmt.Println(pool)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(data)))
	replay := make([]byte, 0, 1500)
	err := pool.Call(serviceMethod, []byte(data), &replay)
	fmt.Println("pool.Call:", err)
	fmt.Println(replay)
	return string(replay)
}
