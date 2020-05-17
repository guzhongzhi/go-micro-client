package transport

import (
	"github.com/guzhongzhi/go-micro-client/options"
)

type Transport interface {
	Type() string
	Init(opts []options.OptionValue)
	SetMetas(v map[string]interface{})
	Do() interface{}
}
