package librajaongkir

type (
	// ResponseError is a struct response error
	ResponseError struct {
		Rajaongkir struct {
			Status struct {
				Code        int    `json:"status"`
				Description string `json:"description"`
			} `json:"status"`
		} `json:"rajaongkir"`
	}
	// ProvinceListResponse is struct province response
	ProvinceListResponse struct {
		Rajaongkir struct {
			Results []struct {
				ProvinceID   string `json:"province_id"`
				ProvinceName string `json:"province"`
			} `json:"results"`
		} `json:"rajaongkir"`
	}
	// CityListResponse is struct city response
	CityListResponse struct {
		Rajaongkir struct {
			Results []struct {
				CityID       string `json:"city_id"`
				ProvinceID   string `json:"province_id"`
				ProvinceName string `json:"province"`
				Type         string `json:"type"`
				CityName     string `json:"city_name"`
				PostalCode   string `json:"postal_code"`
			} `json:"results"`
		} `json:"rajaongkir"`
	}

	// SubdiscrictListResponse is a model
	SubdiscrictListResponse struct {
		Rajaongkir struct {
			Results []struct {
				SubdistrictID   string `json:"subdistrict_id"`
				SubdistrictName string `json:"subdistrict_name"`
				CityID          string `json:"city_id"`
				ProvinceID      string `json:"province_id"`
				ProvinceName    string `json:"province"`
				Type            string `json:"type"`
				CityName        string `json:"city"`
				PostalCode      string `json:"postal_code"`
			} `json:"results"`
		} `json:"rajaongkir"`
	}

	// CostResponse is ...
	CostResponse struct {
		Rajaongkir struct {
			OriginDetails struct {
				CityID     string `json:"city_id"`
				ProvinceID string `json:"province_id"`
				Province   string `json:"province"`
				Type       string `json:"type"`
				CityName   string `json:"city_name"`
				PostalCode string `json:"postal_code"`
			} `json:"origin_details"`
			DestinationDetails struct {
				SubdistrictID   string `json:"subdistrict_id"`
				ProvinceID      string `json:"province_id"`
				Province        string `json:"province"`
				CityID          string `json:"city_id"`
				City            string `json:"city"`
				Type            string `json:"type"`
				SubdistrictName string `json:"subdistrict_name"`
			} `json:"destination_details"`
			Results []struct {
				Code  string `json:"code"`
				Name  string `json:"name"`
				Costs []struct {
					Service     string `json:"service"`
					Description string `json:"description"`
					Cost        []struct {
						Value int    `json:"value"`
						Etd   string `json:"etd"`
						Note  string `json:"note"`
					} `json:"cost"`
				}
			} `json:"results"`
		} `json:"rajaongkir"`
	}

	// WaybillResponse is used to ...
	WaybillResponse struct {
		Rajaongkir struct {
			Result struct {
				Delivered bool `json:"delivered"`
				Summary   struct {
					CourierCode   string `json:"courier_code"`
					CourierName   string `json:"courier_name"`
					WaybillNumber string `json:"waybill_number"`
					ServiceCode   string `json:"service_code"`
					WaybillDate   string `json:"waybill_date"`
					ShipperName   string `json:"shipper_name"`
					ReceiverName  string `json:"receiver_name"`
					Origin        string `json:"origin"`
					Destination   string `json:"destination"`
					Status        string `json:"status"`
				} `json:"summary"`
				Details struct {
					WaybillNumber    string `json:"waybill_number"`
					WaybillDate      string `json:"waybill_date"`
					WaybillTime      string `json:"waybill_time"`
					Weight           string `json:"weight"`
					Origin           string `json:"origin"`
					Destination      string `json:"destination"`
					ShippperName     string `json:"shippper_name"`
					ShipperAddress1  string `json:"shipper_address1"`
					ShipperAddress2  string `json:"shipper_address2"`
					ShipperAddress3  string `json:"shipper_address3"`
					ShipperCity      string `json:"shipper_city"`
					ReceiverName     string `json:"receiver_name"`
					ReceiverAddress1 string `json:"receiver_address1"`
					ReceiverAddress2 string `json:"receiver_address2"`
					ReceiverAddress3 string `json:"receiver_address3"`
					ReceiverCity     string `json:"receiver_city"`
				} `json:"details"`
				DeliveryStatus struct {
					Status      string `json:"status"`
					PodReceiver string `json:"pod_receiver"`
					PodDate     string `json:"pod_date"`
					PodTime     string `json:"pod_time"`
				} `json:"delivery_status"`
				Manifest []struct {
					ManifestCode        string `json:"manifest_code"`
					ManifestDescription string `json:"manifest_description"`
					ManifestDate        string `json:"manifest_date"`
					ManifesTime         string `json:"manifest_time"`
					CityName            string `json:"city_name"`
				} `json:"manifest"`
			} `json:"result"`
		} `json:"rajaongkir"`
	}
)
