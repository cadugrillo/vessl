package system

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	ni "github.com/hilt0n/netif"
)

type InterfaceSet struct {
	InterfacePath string    `json:"InterfacePath"`
	Adapters      []Adapter `json:"Adapters"`
}

type Adapter struct {
	AddrFamily int    `json:"AddrFamily"`
	AddrSource int    `json:"AddrSource"`
	Address    string `json:"Address"`
	Auto       bool   `json:"Auto"`
	Broadcast  string `json:"Broadcast"`
	Gateway    string `json:"Gateway"`
	Hotplug    bool   `json:"Hotplug"`
	Name       string `json:"Name"`
	Netmask    string `json:"Netmask"`
	Network    string `json:"Network"`
}

type HostStats struct {
	CpuUsage      []float64 `json:"CpuUsage"`
	RamTotal      float64   `json:"RamTotal"`
	RamUsed       float64   `json:"RamUsed"`
	RamUsedPct    float64   `json:"RamUsedPct"`
	RamAvailable  float64   `json:"RamAvailable"`
	RamFree       float64   `json:"RamFree"`
	DiskUsage     float64   `json:"DiskUsage"`
	DiskAvailable float64   `json:"DiskAvailable"`
	DiskTotal     float64   `json:"DiskTotal"`
}

var (
	hostStats HostStats
)

func GetNetworkInfo() *ni.InterfaceSet {
	is := ni.Parse(ni.Path("/etc/network/interfaces"))
	return is
}

func SetNetworkInfo(InterfaceSet InterfaceSet) string {

	is := ni.Parse(ni.Path("/etc/network/interfaces"))

	for i := 0; i < len(is.Adapters); i++ {

		switch InterfaceSet.Adapters[i].AddrSource {
		case 1:
			is.Adapters[i].AddrSource = ni.DHCP //{1 - "dhcp", 2 - "static", 3 - "loopback", 4 - "manual"}
		case 2:
			is.Adapters[i].AddrSource = ni.STATIC //{1 - "dhcp", 2 - "static", 3 - "loopback", 4 - "manual"}
		case 3:
			is.Adapters[i].AddrSource = ni.LOOPBACK //{1 - "dhcp", 2 - "static", 3 - "loopback", 4 - "manual"}
		case 4:
			is.Adapters[i].AddrSource = ni.MANUAL //{1 - "dhcp", 2 - "static", 3 - "loopback", 4 - "manual"}
		}

		is.Adapters[i].Address = net.ParseIP(InterfaceSet.Adapters[i].Address)
		is.Adapters[i].Netmask = net.ParseIP(InterfaceSet.Adapters[i].Netmask)
		is.Adapters[i].Gateway = net.ParseIP(InterfaceSet.Adapters[i].Gateway)

	}

	is.Write(ni.Path("/etc/network/interfaces"))
	return "Network Settings updated successfully! (You should restart the system to apply new settings)"

}

func RestartHost() string {

	url := "http://host.docker.internal:4383/host/restart"

	req, _ := http.NewRequest("POST", url, nil)

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		return err.Error()
	}
	return ""
}

func ShutDownHost() string {

	url := "http://host.docker.internal:4383/host/shutdown"

	req, _ := http.NewRequest("POST", url, nil)

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		return err.Error()
	}
	return ""
}

func GetHostStats() HostStats {

	url := "http://host.docker.internal:4383/host/stats"
	//url := "http://localhost:4383/host/stats"

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return hostStats
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println(err.Error())
		return hostStats
	}

	json.Unmarshal(body, &hostStats)
	return hostStats

}
