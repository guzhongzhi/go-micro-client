package transport

import "github.com/guzhongzhi/go-micro-client/client"

type Transport interface {
	Init(client *client.Client)
	Do() interface{}
}
