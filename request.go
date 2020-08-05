package librajaongkir

type (
	// CostRequest is a struct of cost request
	CostRequest struct {
		origin          string
		originType      string
		destination     string
		destinationType string
		weight          int
		courier         string
		length          int
		width           int
		height          int
		diameter        int
	}

	// WaybillRequest is a struct
	WaybillRequest struct {
		waybill string
		courier string
	}
)
