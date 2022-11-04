package volumes

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func GetVolumes() volume.VolumeListOKBody {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	filter := filters.Args{}
	volumes, err := cli.VolumeList(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer ctx.Done()
	defer cli.Close()
	return volumes
}

func RemoveVolume(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	force := false

	fmt.Print("Removing volume ", Id, "... ")
	if err := cli.VolumeRemove(ctx, Id, force); err != nil {
		return err.Error()
	}
	fmt.Println("Success")

	defer ctx.Done()
	defer cli.Close()
	return "Volume successfully removed"
}
