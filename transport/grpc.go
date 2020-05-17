package transport

import (
	"github.com/guzhongzhi/go-micro-client/options"
)

type Grpc struct {

}

func (s *Grpc) Init(opts []options.OptionValue) {

}

func (s *Grpc) Do() interface{}{
	return nil
}


func (s *Grpc) Type() string  {
	return "grpc"
}


func (s *Grpc) SetMetas(v map[string]interface{}) {
	panic("implement me")
}