package client

import (
	"fmt"
	"github.com/guzhongzhi/go-micro-client/options"
	"github.com/guzhongzhi/go-micro-client/transport"
)

type Client struct {
	name string
	transport  transport.Transport
	meta map[string]interface{}
}

func (s *Client) getTrans() transport.Transport{

	return s.transport
}

func (s *Client) Call(serviceName string,params interface{}) interface{} {
	return s.transport.Do(serviceName,params)
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
		case "meta":
			c.meta = opt.Value.(map[string]interface{})
		case "transport":
			transportType := fmt.Sprintf("%v",opt.Value)

			switch transportType  {
			case "http":
				c.transport = &transport.Http{}
				c.transport.Init(opt.Children)
			case "grpc":
				c.transport = &transport.Grpc{}
				c.transport.Init(opt.Children)
			default:
				panic("invalid transport")
			}
		}
	}
	c.transport.SetMetas(c.meta)

	return c
}