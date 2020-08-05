package librajaongkir

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/valyala/fasthttp"
)

const (
	Starter TypeAcount = "Starter"
	Basic   TypeAcount = "Basic"
	Pro     TypeAcount = "Pro"

	URIStarter string = "https://api.rajaongkir.com/starter"
	URIBasic   string = "https://api.rajaongkir.com/basic"
	URIPro     string = "https://pro.rajaongkir.com/api"
)

type (
	TypeAcount string

	// Client is used to ...
	Client struct {
		TypeAcount TypeAcount
		APIKey     string
	}
)

func NewClient() *Client {
	return &Client{}
}

// NewRequest test is ...
func (c *Client) NewRequest(method, fullPath string, headers map[string]string, body io.Reader) (*fasthttp.Request, error) {
	log.Info(fullPath)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fullPath)
	req.Header.SetMethod(method)
	if method == fasthttp.MethodPost {
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)
		req.SetBody(buf.Bytes())
	}
	if headers != nil {
		for k, vv := range headers {
			req.Header.Set(k, vv)
		}
	}
	return req, nil
}

// ExecuteRequest is ...
func (c *Client) ExecuteRequest(req *fasthttp.Request, v interface{}) error {
	start := time.Now()
	resp := fasthttp.AcquireResponse()
	httpClient := &fasthttp.Client{}
	err := httpClient.Do(req, resp)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("Completed in ", time.Since(start))

	if v != nil && resp.StatusCode() == 200 {
		if err = json.Unmarshal(resp.Body(), v); err != nil {
			log.Error(err)
			return err
		}
		return nil
	}
	var respErr ResponseError
	if err = json.Unmarshal(resp.Body(), &respErr); err != nil {
		return err
	}
	return errors.New(respErr.Rajaongkir.Status.Description)
}

// Call is ...
func (c *Client) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, header, body)
	if err != nil {
		return err
	}
	return c.ExecuteRequest(req, v)
}
