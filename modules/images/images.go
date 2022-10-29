package images

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	ImageListOptions   types.ImageListOptions
	ImageRemoveOptions types.ImageRemoveOptions
)

func GetImages() []types.ImageSummary {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	ImageListOptions.All = false

	images, err := cli.ImageList(ctx, ImageListOptions)
	if err != nil {
		panic(err)
	}

	return images
}

func RemoveImage(Id string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err.Error()
	}

	ImageRemoveOptions.Force = false
	ImageRemoveOptions.PruneChildren = false

	fmt.Print("Removing image ", Id[:10], "... ")
	if _, err := cli.ImageRemove(ctx, Id, ImageRemoveOptions); err != nil {
		return err.Error()
	}
	fmt.Println("Success")
	return "Image successfully removed"
}
