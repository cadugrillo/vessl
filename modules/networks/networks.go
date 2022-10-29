package networks

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	NetworkListOptions   types.NetworkListOptions
	NetworkCreateOptions types.NetworkCreate
)

func GetNetworks() []types.NetworkResource {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	networks, err := cli.NetworkList(ctx, NetworkListOptions)
	if err != nil {
		panic(err)
	}

	return networks
}

func CreateNetwork(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	NetworkCreateOptions.CheckDuplicate = true

	resp, err := cli.NetworkCreate(ctx, Id, NetworkCreateOptions)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Network ID: %s created successfuly", resp.ID[:10])
}

func RemoveNetwork(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	fmt.Print("Removing network ", Id[:10], "... ")
	if err := cli.NetworkRemove(ctx, Id); err != nil {
		return err.Error()
	}
	fmt.Println("Success")
	return "Network successfully removed"
}
