package librajaongkir

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/valyala/fasthttp"
)

// Core is core of rajaonkir
type Core struct {
	*Client
}

// CallPro is used to postpaid
func (cr *Core) CallPro(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}
	path = URIPro + path
	return cr.Client.Call(method, path, header, body, v)
}

// ProvinceList is used to get provice
func (cr *Core) ProvinceList() (res ProvinceListResponse, err error) {
	urlPath := "province"
	headers := map[string]string{
		"key": cr.Client.APIKey,
	}
	err = cr.CallPro(fasthttp.MethodGet, urlPath, headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// CityList is a func to get all city
func (cr *Core) CityList(provinceID int) (res CityListResponse, err error) {
	urlPath := fmt.Sprintf("city?province=%v", provinceID)
	headers := map[string]string{
		"key": cr.Client.APIKey,
	}
	err = cr.CallPro(fasthttp.MethodGet, urlPath, headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// SubdistrictList is a func to get all city
func (cr *Core) SubdistrictList(cityID int) (res SubdiscrictListResponse, err error) {
	urlPath := fmt.Sprintf("subdistrict?city=%v", cityID)
	headers := map[string]string{
		"key": cr.Client.APIKey,
	}
	err = cr.CallPro(fasthttp.MethodGet, urlPath, headers, nil, &res)
	if err != nil {
		return
	}
	return
}

// Cost is a func is used to get shipping cost.
func (cr *Core) Cost(req CostRequest) (res CostResponse, err error) {
	urlPath := "cost"
	headers := map[string]string{
		"key":          cr.Client.APIKey,
		"Content-Type": "application/x-www-form-urlencoded",
	}
	err = cr.CallPro(fasthttp.MethodPost, urlPath, headers, bytes.NewBufferString(StructToMap(&req).Encode()), &res)
	if err != nil {
		return
	}
	return
}

// Waybill is a func  used to get tracking shipping
func (cr *Core) Waybill(req WaybillRequest) (res WaybillResponse, err error) {
	urlPath := "waybill"
	headers := map[string]string{
		"key":          cr.Client.APIKey,
		"Content-Type": "application/x-www-form-urlencoded",
	}
	err = cr.CallPro(fasthttp.MethodPost, urlPath, headers, bytes.NewBufferString(StructToMap(&req).Encode()), &res)
	if err != nil {
		return
	}
	return
}
