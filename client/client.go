package client

import (
	"fmt"
	"github.com/guzhongzhi/go-micro-client/options"
	"github.com/guzhongzhi/go-micro-client/transport"
)

type Client struct {
	name string
	transport  string
	trans transport.Transport
}

func (s *Client) getTrans() transport.Transport{

	return s.trans
}

func (s *Client) Call(serviceName string,params interface{}) interface{} {
	return s.trans.Do()
}


func NewClient(name string,opts ...options.Option) *Client{

	nameOption := options.NewNameOption(name)
	opts = append(opts,nameOption)

	optValues := options.OptionValues(opts...)
	c := &Client{}
	for _,opt := range optValues {
		switch opt.Name {
		case "name":
			c.name = fmt.Sprintf("%v",opt.Value)
		case "transport":
			c.transport = fmt.Sprintf("%v",opt.Value)
		}
	}
	switch c.transport  {
	case "http":
		c.trans = &transport.Http{}
		c.trans.Init(c)
	case "grpc":
		c.trans = &transport.Grpc{}
		c.trans.Init(c)
	}
	return c
}