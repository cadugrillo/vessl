package containers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"math"
	"os"

	apps_repository "vessl/modules/apps-repository"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type ContainerStats struct {
	Id       string  `json:"Id"`
	Name     string  `json:"Name"`
	CpuPct   float64 `json:"CpuPct"`
	MemUsage float64 `json:"MemUsage"`
	MemLimit float64 `json:"MemLimit"`
	MemPct   float64 `json:"MemPct"`
}

func GetContainers(networkName string) []types.Container {

	var ContainerListOptions types.ContainerListOptions

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return []types.Container{}
	}

	if networkName == "All" {
		ContainerListOptions = types.ContainerListOptions{}
	} else {
		networkKeyValuePair := filters.KeyValuePair{Key: "network", Value: networkName}
		networkFilter := filters.NewArgs(networkKeyValuePair)
		ContainerListOptions.Filters = networkFilter
	}

	ContainerListOptions.All = true
	containers, err := cli.ContainerList(ctx, ContainerListOptions)
	if err != nil {
		return []types.Container{}
	}
	defer ctx.Done()
	defer cli.Close()

	return containers
}

func InstallContainer(AppTemplate apps_repository.Template) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	out, err := cli.ImagePull(ctx, AppTemplate.Image, types.ImagePullOptions{})
	if err != nil {
		return err.Error()
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	_, portBindings, err := nat.ParsePortSpecs(AppTemplate.Ports)
	if err != nil {
		fmt.Println("Unable to create docker port")
		return "Unable to create docker port"
	}

	// Configured hostConfig:
	HostConfig := &container.HostConfig{
		Binds:        AppTemplate.Volumes,
		NetworkMode:  container.NetworkMode(AppTemplate.Network),
		PortBindings: portBindings,
		RestartPolicy: container.RestartPolicy{
			Name: AppTemplate.Restart_policy,
		},
		LogConfig: container.LogConfig{
			Type:   "json-file",
			Config: map[string]string{},
		},
	}

	//////NETWORK CONFIGURATION////////
	NetworkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: AppTemplate.Network,
	}
	NetworkConfig.EndpointsConfig[AppTemplate.Network] = gatewayConfig

	// Configuration
	ContainerConfig := &container.Config{
		Image:    AppTemplate.Image,
		Hostname: AppTemplate.Hostname,
		Env:      AppTemplate.Env,
	}

	resp, err := cli.ContainerCreate(ctx, ContainerConfig, HostConfig, NetworkConfig, nil, AppTemplate.Hostname)
	if err != nil {
		return err.Error()
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err.Error()
	}

	defer ctx.Done()
	defer cli.Close()
	fmt.Println(resp.ID)
	return fmt.Sprintf("App ID: %s installed successfuly", resp.ID[:10])
}

func StartContainer(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	fmt.Print("Starting container ", Id[:10], "... ")
	if err := cli.ContainerStart(ctx, Id, types.ContainerStartOptions{}); err != nil {
		return err.Error()
	}

	defer ctx.Done()
	defer cli.Close()
	fmt.Println("Success")
	return "App successfully started"
}

func StopContainer(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	fmt.Print("Stopping container ", Id[:10], "... ")
	if err := cli.ContainerStop(ctx, Id, nil); err != nil {
		return err.Error()
	}

	defer ctx.Done()
	defer cli.Close()
	fmt.Println("Success")
	return "App successfully stopped"
}

func RestartContainer(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	fmt.Print("Restarting container ", Id[:10], "... ")
	if err := cli.ContainerRestart(ctx, Id, nil); err != nil {
		return err.Error()
	}

	defer ctx.Done()
	defer cli.Close()
	fmt.Println("Success")
	return "App successfully restarted"
}

func RemoveContainer(Id string) string {

	var ContainerRemoveOptions types.ContainerRemoveOptions

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	ContainerRemoveOptions.Force = true
	ContainerRemoveOptions.RemoveVolumes = true

	fmt.Print("Removing container ", Id[:10], "... ")
	if err := cli.ContainerRemove(ctx, Id, ContainerRemoveOptions); err != nil {
		return err.Error()
	}

	defer ctx.Done()
	defer cli.Close()
	fmt.Println("Success")
	return "App successfully removed"
}

func Logs(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	options := types.ContainerLogsOptions{ShowStdout: true, Timestamps: true}
	out, err := cli.ContainerLogs(ctx, Id, options)
	if err != nil {
		return err.Error()
	}

	b, err := io.ReadAll(out)
	if err != nil {
		return err.Error()
	}

	defer out.Close()
	defer ctx.Done()
	defer cli.Close()
	return string(b)
}

func GetDockerServerInfo() types.Info {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return types.Info{}
	}

	info, err := cli.Info(ctx)
	if err != nil {
		return types.Info{}
	}

	defer ctx.Done()
	defer cli.Close()
	return info
}

func GetContainerStats(Id string) ContainerStats {

	var cstats types.Stats

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return ContainerStats{Name: "Internal Error 1"}
	}

	stats, err := cli.ContainerStats(ctx, Id, false)
	if err != nil {
		return ContainerStats{Name: "Internal Error 2"}
	}
	statsBody, err := io.ReadAll(stats.Body)
	if err != nil {
		return ContainerStats{Name: "Internal Error 3"}
	}
	defer stats.Body.Close()
	err = json.Unmarshal(statsBody, &cstats)
	if err != nil {
		return ContainerStats{Name: "Internal Error 4"}
	}

	defer ctx.Done()
	defer cli.Close()
	return calculateStats(cstats)
}

func GetCompleteStats() []ContainerStats {

	var totalStats []ContainerStats
	var containerStats ContainerStats
	var cstats types.Stats

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return []ContainerStats{{Name: "Internal Error 1"}}
	}

	runningContainers := GetContainers("")

	for i := 0; i < len(runningContainers); i++ {

		stats, err := cli.ContainerStats(ctx, runningContainers[i].ID, false)
		if err != nil {
			return []ContainerStats{{Name: "Internal Error 2"}}
		}

		statsBody, err := io.ReadAll(stats.Body)
		if err != nil {
			return []ContainerStats{{Name: "Internal Error 3"}}
		}
		defer stats.Body.Close()

		err = json.Unmarshal(statsBody, &cstats)
		if err != nil {
			return []ContainerStats{{Name: "Internal Error 4"}}
		}

		containerStats = calculateStats(cstats)
		containerStats.Id = runningContainers[i].ID[:10]
		containerStats.Name = runningContainers[i].Names[0]

		totalStats = append(totalStats, containerStats)

	}

	defer ctx.Done()
	defer cli.Close()
	return totalStats
}

func calculateStats(cstats types.Stats) ContainerStats {

	var containerStats ContainerStats

	cpu_delta := float64(cstats.CPUStats.CPUUsage.TotalUsage - cstats.PreCPUStats.CPUUsage.TotalUsage)
	system_cpu_delta := float64(cstats.CPUStats.SystemUsage - cstats.PreCPUStats.SystemUsage)
	number_cpus := float64(cstats.CPUStats.OnlineCPUs)

	//fmt.Println(cstats.MemoryStats.Usage, cstats.MemoryStats.Stats["cache"])
	containerStats.CpuPct = math.Round(((cpu_delta/system_cpu_delta)*number_cpus*100)*100) / 100
	containerStats.MemUsage = math.Round(byteToMegabyte(cstats.MemoryStats.Usage-cstats.MemoryStats.Stats["cache"])*100) / 100
	containerStats.MemLimit = math.Round(byteToMegabyte(cstats.MemoryStats.Limit)*100) / 100
	containerStats.MemPct = math.Round(((containerStats.MemUsage/containerStats.MemLimit)*100)*100) / 100

	return containerStats
}

func byteToMegabyte(Byte uint64) float64 {
	return (float64(Byte) / 1024) / 1024
}
