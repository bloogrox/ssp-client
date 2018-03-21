package sspclient

type DSP struct {
	BidURL    string   `json:"bid_url"`
	Name      string   `json:"name"`
	ExtFields []string `json:"ext_fields"`
}

type Request struct {
	DSP        *DSP              `json:"dsp"`
	Subscriber map[string]string `json:"subscriber"`
}
