package containers

import (
	"context"
	"fmt"
	"io"
	"os"

	apps_repository "vessl/modules/apps-repository"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

var (
	ContainerListOptions   types.ContainerListOptions
	ContainerRemoveOptions types.ContainerRemoveOptions
)

func GetContainers(networkName string) []types.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	networkKeyValuePair := filters.KeyValuePair{Key: "network", Value: networkName}
	networkFilter := filters.NewArgs(networkKeyValuePair)

	ContainerListOptions.Filters = networkFilter
	ContainerListOptions.All = true
	containers, err := cli.ContainerList(ctx, ContainerListOptions)
	if err != nil {
		panic(err)
	}

	return containers
}

func InstallContainer(AppTemplate apps_repository.Template) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, AppTemplate.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
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

	fmt.Println(resp.ID)
	return fmt.Sprintf("App ID: %s installed successfuly", resp.ID[:10])
}

func StartContainer(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Print("Starting container ", Id[:10], "... ")
	if err := cli.ContainerStart(ctx, Id, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Println("Success")
	return "App successfully started"
}

func StopContainer(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Print("Stopping container ", Id[:10], "... ")
	if err := cli.ContainerStop(ctx, Id, nil); err != nil {
		panic(err)
	}
	fmt.Println("Success")
	return "App successfully stopped"
}

func RestartContainer(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Print("Restarting container ", Id[:10], "... ")
	if err := cli.ContainerRestart(ctx, Id, nil); err != nil {
		panic(err)
	}
	fmt.Println("Success")
	return "App successfully restarted"
}

func RemoveContainer(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	ContainerRemoveOptions.Force = true
	ContainerRemoveOptions.RemoveVolumes = true

	fmt.Print("Removing container ", Id[:10], "... ")
	if err := cli.ContainerRemove(ctx, Id, ContainerRemoveOptions); err != nil {
		panic(err)
	}
	fmt.Println("Success")
	return "App successfully removed"
}

func Logs(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, Id, options)
	if err != nil {
		panic(err)
	}

	if b, err := io.ReadAll(out); err == nil {
		return string(b)
	}
	//io.Copy(os.Stdout, out)
	return "no text to show"
}

func GetDockerServerInfo() types.Info {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	info, err := cli.Info(ctx)
	if err != nil {
		panic(err)
	}

	return info
}
