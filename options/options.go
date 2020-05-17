package options

const TransportHttp = "http"
const TransportGrpc = "grpc"

type Option func() OptionValue

type OptionValue struct {
	Value interface{}
	Name  string
	Children []OptionValue
}

func NewMetas(v map[string]interface{}) Option {
	return func() OptionValue {
		return OptionValue{
			Value: v,
			Name:  "meta",
		}
	}
}

func NewHttpTransportOption(apiUrl string) Option {
	return func() OptionValue {
		return OptionValue{
			Name:"transport",
			Value:"http",
			Children:[]OptionValue{
				OptionValue{
					Value: apiUrl,
					Name:  "api",
				},
			},
		}
	}
}

func NewGrpcTransportOption(registry string, address string) Option {
	return func() OptionValue {
		return OptionValue{
			Name:"transport",
			Value:"grpc",
			Children:[]OptionValue{
				OptionValue{
					Value: registry,
					Name:  "registry",
				},
				OptionValue{
					Value: address,
					Name:  "address",
				},
			},
		}
	}
}

func NewNameOption(v string) Option {
	return func() OptionValue {
		return OptionValue{
			Value: v,
			Name:  "name",
		}
	}
}

func OptionValues(opts ...Option) []OptionValue{
	v := make([]OptionValue,0)
	for _,pts := range opts {
		v = append(v,pts())
	}
	return v
}
