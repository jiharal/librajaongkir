package librajaongkir

import "testing"

func initClient() *Core {
	client := NewClient()
	client.APIKey = ""
	client.TypeAcount = Pro
	return &Core{client}
}

func TestProvince(t *testing.T) {
	core := initClient()
	resp, err := core.ProvinceList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
func TestCityList(t *testing.T) {
	core := initClient()
	resp, err := core.CityList(20)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
func TestSubdistrictList(t *testing.T) {
	core := initClient()
	resp, err := core.SubdistrictList(100)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
func TestCost(t *testing.T) {
	core := initClient()
	req := CostRequest{
		origin:          "501",
		originType:      "city",
		destination:     "574",
		destinationType: "subdistrict",
		weight:          17000,
		courier:         "jne",
	}
	resp, err := core.Cost(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
func TestWaibill(t *testing.T) {
	core := initClient()
	req := WaybillRequest{
		waybill: "18009898426",
		courier: "pos",
	}
	resp, err := core.Waybill(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
