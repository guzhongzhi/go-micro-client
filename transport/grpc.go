package transport

import (
	"github.com/guzhongzhi/go-micro-client/options"
)

type Grpc struct {

}

func (s *Grpc) Init(opts []options.OptionValue) {

}

func (s *Grpc) Do(serviceName string,params interface{}) (interface{},error){
	return nil,nil
}


func (s *Grpc) Type() string  {
	return "grpc"
}


func (s *Grpc) SetMetas(v map[string]interface{}) {
	panic("implement me")
}