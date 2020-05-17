package transport

import (
	"github.com/guzhongzhi/go-micro-client/options"
)

const TransportGrpc = "grpc"

type Grpc struct {
	meta map[string]interface{}
}

func (s *Grpc) Init(opts []options.OptionValue) {

}

func (s *Grpc) Do(serviceName string, params interface{}) (interface{}, error) {
	return nil, nil
}

func (s *Grpc) Type() string {
	return TransportGrpc
}

func (s *Grpc) SetMetas(v map[string]interface{}) {
	s.meta = v
}
