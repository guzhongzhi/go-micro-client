package transport

import (
	"github.com/guzhongzhi/go-micro-client/options"
)

type Http struct {

}

func (s *Http) SetMetas(v map[string]interface{}) {
	panic("implement me")
}

func (s *Http) Init(opts []options.OptionValue)  {

}

func (s *Http) Do() interface{}{
	return nil
}

func (s *Http) Type() string  {
	return "http"
}