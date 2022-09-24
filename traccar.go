package traccar

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/budimanlai/go-traccar/test/helpers"
	"github.com/eqto/go-json"
	"github.com/valyala/fasthttp"
)

type Traccar struct {
	ServerUrl string
	Token     string
	Username  string
	Password  string
}

func NewTraccar(url string, username string, password string, token string) *Traccar {
	return &Traccar{
		ServerUrl: url,
		Token:     token,
		Username:  username,
		Password:  password,
	}
}

func (t *Traccar) get(endpoint string) (*fasthttp.Response, error) {
	url := t.ServerUrl + "/" + endpoint

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.DisableNormalizing()

	basic := base64.StdEncoding.EncodeToString([]byte(t.Username + ":" + t.Password))

	// req.Header.SetContentType(`application/json`)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set(`Authorization`, `basic `+basic)

	respClone := &fasthttp.Response{}
	e1 := fasthttp.DoTimeout(req, resp, 60*time.Second)
	resp.CopyTo(respClone)

	fmt.Println("Call:", url)
	fmt.Println(string(respClone.Body()))
	return respClone, e1
}

func (t *Traccar) GetSession() (string, error) {
	params := url.Values{}
	params.Add("token", t.Token)

	url := t.ServerUrl + "/session" + "?" + params.Encode()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.DisableNormalizing()

	// req.Header.SetContentType(`application/json`)
	req.Header.SetMethod(fasthttp.MethodGet)

	respClone := &fasthttp.Response{}
	e1 := fasthttp.DoTimeout(req, resp, 60*time.Second)
	resp.CopyTo(respClone)

	fmt.Println("Call:", url)
	fmt.Println(string(respClone.Body()))

	if e1 != nil {
		return "", e1
	}

	cookie := resp.Header.PeekCookie("JSESSIONID")

	return string(cookie), nil
}

/*
 * Get device trips
 */
func (t *Traccar) Trips(device_id []string, from string, to string, page int, start int, limit int) ([]json.Object, error) {
	params := url.Values{}
	params.Add("token", t.Token)
	params.Add("from", helpers.NormalizeDateFormat(from))
	params.Add("to", helpers.NormalizeDateFormat(to))
	params.Add("page", strconv.Itoa(page))
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	for _, d := range device_id {
		params.Add("deviceId", d)
	}

	resp, e := t.get(`reports/trips` + "?" + params.Encode())
	if e != nil {
		return nil, e
	}

	jsonBody, e := json.ParseArray(resp.Body())
	if e != nil {
		return nil, e
	}
	return jsonBody, nil
}
