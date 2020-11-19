package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func workerDispacher(services map[string]bool) map[string]StatusResponse {

	results := map[string]StatusResponse{}

	if services["router"] {
		results["router"] = pingRouter()
	}

	if services["gateway"] {
		results["gateway"] = pingGateway()
	}

	return results
}

func getNumberOfJobs(m map[string]bool) int {
	c := 0
	for _, v := range m {
		if v {
			c++
		}
	}
	return c
}

func pingRouter() StatusResponse {
	const endpoint = "http://192.168.86.1/api/v1/status"

	resp, err := http.Get(endpoint)

	if err != nil {
		return StatusResponse{status: "FAIL", message: err.Error()}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return StatusResponse{status: "FAIL", message: err.Error()}
	}

	var w = new(RouterStatusResponse)
	err = json.Unmarshal(body, &w)
	if err != nil {
		return StatusResponse{status: "FAIL", message: err.Error()}
	}

	if w.Wan.Online {
		return StatusResponse{status: "OK", message: fmt.Sprintf("Online\n\tGateway Ip Addr: %v\n\tUsing DNS Servers: %v", w.Wan.GatewayIpAddress, w.Wan.NameServers)}
	}
	return StatusResponse{status: "WARN", message: fmt.Sprintf("Router Cannot Connect to Internet")}

}

func pingGateway() StatusResponse {
	const endpoint = "http://10.2.6.1/graphs/iface/vlan103/"

	resp, err := http.Get(endpoint)

	if err != nil {
		return StatusResponse{status: "FAIL", message: fmt.Sprintf("Could not connect to gateway.\n\t%v", err)}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return StatusResponse{status: "WARN", message: fmt.Sprintf("Got unexpected StatusCode = %v,\n\t%v", resp.StatusCode, http.StatusText(resp.StatusCode))}
	}

	return StatusResponse{status: "OK", message: "Pinnged Gateway Successfully"}
}

func pingDNS() string {
	return "pingDNS is WIP"
}

func pingDiscordBot() string {
	return "pingDiscordBot is WIP"
}

func pingRaspberryPi() string {
	return "pingRaspberryPi is WIP"
}
