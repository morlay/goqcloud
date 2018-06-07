package signature

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/morlay/goqcloud/transform"
)

type Credential struct {
	SecretId  string
	SecretKey string
}

func NewAutoSignTransport(credential *Credential) transform.Transport {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &AutoSignRoundTripper{
			Credential:       credential,
			NextRoundTripper: rt,
		}
	}
}

type AutoSignRoundTripper struct {
	*Credential
	NextRoundTripper http.RoundTripper
}

var _ interface {
	http.RoundTripper
} = (*AutoSignRoundTripper)(nil)

func (t *AutoSignRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	method := SHA256

	query := request.URL.Query()

	query.Set("SecretId", t.SecretId)
	query.Set("Timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	query.Set("Nonce", strconv.Itoa(rand.Int()))
	query.Set("SignatureMethod", string(method))

	query.Del("Signature")

	keys := make([]string, 0)
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	urlPath := request.URL.Path

	if len(urlPath) == 0 || urlPath[len(urlPath)-1] != '/' {
		urlPath += "/"
	}

	text := request.Method + request.URL.Host + urlPath + "?"

	for i := range keys {
		if i > 0 {
			text += "&"
		}
		k := keys[i]
		value := query.Get(k)
		if value == "" {
			continue
		}
		text += fmt.Sprintf("%v=%v", strings.Replace(k, "_", ".", -1), value)
	}

	query.Set("Signature", t.Signature(text, method))

	request.URL.RawQuery = query.Encode()

	return t.NextRoundTripper.RoundTrip(request)
}

func (t *AutoSignRoundTripper) Signature(s string, method SignatureMethod) string {
	hashed := hmac.New(sha1.New, []byte(t.SecretKey))
	if method == SHA256 {
		hashed = hmac.New(sha256.New, []byte(t.SecretKey))
	}
	hashed.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}

type SignatureMethod string

const (
	SHA256 SignatureMethod = "HmacSHA256"
	SHA1   SignatureMethod = "HmacSHA1"
)
