package goqcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/morlay/goqcloud/signature"
	"github.com/morlay/goqcloud/transform"
	"github.com/sirupsen/logrus"
)

type Client interface {
	Request(service string, action string, version string) RequestSender
}

type RequestSender interface {
	Do(req interface{}, resp interface{}) error
}

type MaybeErrorResponse interface {
	GetError() error
}

func NewClientWithCredential(secretId, secretKey string, opts ...*ClientOption) *ClientWithCredential {
	return &ClientWithCredential{
		ClientOption: ComposeClientOptions(opts...),
		Credential: signature.Credential{
			SecretId:  secretId,
			SecretKey: secretKey,
		},
	}
}

func ComposeClientOptions(opts ...*ClientOption) *ClientOption {
	if len(opts) == 0 {
		return nil
	}

	opt := &ClientOption{}

	for i := range opts {
		o := opts[i]
		if o.Timeout != 0 {
			opt.Timeout = o.Timeout
		}
		if o.Transports != nil {
			opt.Transports = o.Transports
		}
	}

	return opt
}

func ClientOptionWithTimeout(timeout time.Duration) *ClientOption {
	return &ClientOption{
		Timeout: timeout,
	}
}

func ClientOptionWithTransports(transports ...transform.Transport) *ClientOption {
	return &ClientOption{
		Transports: transports,
	}
}

type ClientOption struct {
	Timeout    time.Duration
	Transports []transform.Transport
}

type ClientWithCredential struct {
	*ClientOption
	signature.Credential
}

func (c *ClientWithCredential) Request(service string, action string, version string) RequestSender {
	opt := c.ClientOption
	if opt == nil {
		opt = &ClientOption{}
	}

	opt.Transports = append(opt.Transports, signature.NewAutoSignTransport(&c.Credential))

	return &httpRequestSender{
		service:      service,
		action:       action,
		version:      version,
		ClientOption: *opt,
	}
}

type httpRequestSender struct {
	service string
	version string
	action  string
	ClientOption
}

func (r *httpRequestSender) Do(req interface{}, resp interface{}) error {
	request, errForTransform := r.transformRequest(req)
	if errForTransform != nil {
		return NewTencentCloudError(
			"SDK.TransformRequestFailed",
			errForTransform.Error(),
		)
	}

	if r.Timeout == 0 {
		r.Timeout = 60 * time.Second
	}

	client := &http.Client{}
	client.Timeout = r.Timeout
	client.Transport = transform.ComposeTransports(r.Transports...)(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   client.Timeout,
			KeepAlive: 0,
		}).DialContext,
		DisableKeepAlives: true,
	})

	response, errForDo := client.Do(request)
	if errForDo != nil {
		return NewTencentCloudError(
			"SDK.DoRequestFailed",
			errForDo.Error(),
		)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return NewTencentCloudError(
			"SDK.ResponseBodyReadFailed",
			err.Error(),
		)
	}

	if err := json.Unmarshal(data, &TencentCloudResponse{Response: resp}); err != nil {
		return NewTencentCloudError(
			"SDK.JSONUnmarshalFailed",
			err.Error(),
		)
	}

	if mayError, ok := resp.(MaybeErrorResponse); ok {
		return mayError.GetError()
	}
	return nil
}

func (r *httpRequestSender) transformRequest(req interface{}) (*http.Request, error) {
	u, _ := url.Parse(fmt.Sprintf("https://%s.tencentcloudapi.com", r.service))

	s := transform.NewParameterScanner()
	if err := s.Scan(reflect.ValueOf(req)); err != nil {
		return nil, err
	}
	params := s.Params()

	params.Set("Action", r.action)
	params.Set("Version", r.version)

	u.RawQuery = params.Encode()

	return http.NewRequest("GET", u.String(), nil)
}

type TencentCloudResponse struct {
	Response interface{} `json:"Response"`
}

type TencentCloudBaseResponse struct {
	RequestId string             `json:"RequestId"`
	Error     *TencentCloudError `json:"Error,omitempty"`
}

func (resp *TencentCloudBaseResponse) GetError() error {
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func NewTencentCloudError(code string, msg string) *TencentCloudError {
	return &TencentCloudError{
		Code:    code,
		Message: msg,
	}
}

type TencentCloudError struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

func (e *TencentCloudError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewLogTransport() transform.Transport {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &LogTransport{
			NextRoundTripper: rt,
		}
	}
}

type LogTransport struct {
	NextRoundTripper http.RoundTripper
}

func (t *LogTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	logrus.Infof("%s %s", request.Method, request.URL.String())
	return t.NextRoundTripper.RoundTrip(request)
}
