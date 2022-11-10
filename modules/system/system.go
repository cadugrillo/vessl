package system

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"time"

	ni "github.com/hilt0n/netif"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
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

	url := "http://vessl-host-control:4383/host/restart"

	req, _ := http.NewRequest("POST", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err.Error()
	}

	defer res.Body.Close()
	return ""
}

func ShutDownHost() string {

	url := "http://vessl-host-control:4383/host/shutdown"

	req, _ := http.NewRequest("POST", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err.Error()
	}

	defer res.Body.Close()
	return ""
}

func GetHostStats() HostStats {

	var hostStats HostStats

	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		fmt.Println(err.Error())
		return HostStats{}
	}
	hostStats.CpuUsage = cpuPercent

	for i := 0; i < len(hostStats.CpuUsage); i++ {
		hostStats.CpuUsage[i] = math.Round(hostStats.CpuUsage[i]*100) / 100
	}

	m, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err.Error())
		return HostStats{}
	}
	hostStats.RamTotal = getReadableSize(m.Total)
	hostStats.RamUsed = getReadableSize(m.Used)
	hostStats.RamUsedPct = math.Round(m.UsedPercent*100) / 100
	hostStats.RamAvailable = getReadableSize(m.Available)
	hostStats.RamFree = getReadableSize(m.Free)

	diskUsage, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err.Error())
		return HostStats{}
	}

	hostStats.DiskUsage = math.Round(diskUsage.UsedPercent*100) / 100
	hostStats.DiskAvailable = getReadableSize(diskUsage.Free)
	hostStats.DiskTotal = getReadableSize(diskUsage.Total)

	return hostStats
}

func getReadableSize(sizeInBytes uint64) (readableSize float64) {
	var (
		units = []string{"KB", "MB"} //, "GB", "TB", "PB}
		size  = float64(sizeInBytes)
		i     = 0
	)
	for ; i < len(units); i++ { //&& size >= 1000
		size = size / 1024
	}

	return math.Round(size*100) / 100
}
