package transport

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guzhongzhi/go-micro-client/options"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Http struct {
	meta map[string]interface{}
	api  string
	timeout time.Duration
}

func (s *Http) SetMetas(v map[string]interface{}) {
	panic("implement me")
}

func (s *Http) Init(opts []options.OptionValue) {
	for _, opt := range opts {
		switch opt.Name {
		case "api":
			s.api = opt.Value.(string)
			if !strings.HasSuffix(s.api, "/") {
				s.api = s.api + "/"
			}
		case "timeout":
			s.timeout = opt.Value.(time.Duration)
		}
	}
}

func (s *Http) Do(serviceName string, params interface{}) (interface{},error) {
	servicePath := strings.Replace(serviceName, ".", "/", -1)
	fullUrl := s.api + servicePath
	postJS := ""
	switch params.(type) {
	case string:
		postJS = params.(string)
	case io.Reader:
		d,err := ioutil.ReadAll(params.(io.Reader))
		if err != nil {
			return nil,err
		}
		postJS = string(d)
	default:
		d,err := json.Marshal(params)
		if err != nil {
			return nil,err
		}
		postJS = string(d)
	}

	r := bytes.NewBuffer([]byte(postJS))
	req, err := http.NewRequest("POST", fullUrl, r)
	if err != nil {
		return nil,err
	}
	for k,v := range s.meta {
		req.Header.Set(k,fmt.Sprintf("%v",v))
	}
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{
		Timeout:s.timeout,
	}
	rsp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode == 200 {
		body, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return nil,err
		}
		return string(body),nil
	}

	return nil,errors.New(rsp.Status)
}

func (s *Http) Type() string {
	return "http"
}
