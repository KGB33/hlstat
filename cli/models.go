package cli

type StatusResponse struct {
	status  string `One of: OK, WARN, FAIL`
	message string
}

type RouterStatusResponse struct {
	Wan WanResponse `json:"wan"`
}

type WanResponse struct {
	GatewayIpAddress string   `json:"gatewayIpAddress"`
	Online           bool     `json:"online"`
	NameServers      []string `json:"nameServers"`
}
