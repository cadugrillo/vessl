package images

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetImages() []types.ImageSummary {

	var ImageListOptions types.ImageListOptions

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return []types.ImageSummary{}
	}

	ImageListOptions.All = false

	images, err := cli.ImageList(ctx, ImageListOptions)
	if err != nil {
		return []types.ImageSummary{}
	}

	defer ctx.Done()
	defer cli.Close()
	return images
}

func RemoveImage(Id string) string {

	var ImageRemoveOptions types.ImageRemoveOptions

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

	defer ctx.Done()
	defer cli.Close()
	fmt.Println("Success")
	return "Image successfully removed"
}

func InspectImage(Id string) (types.ImageInspect, error) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return types.ImageInspect{}, err
	}
	ii, _, err := cli.ImageInspectWithRaw(ctx, Id)
	if err != nil {
		return types.ImageInspect{}, err
	}
	defer ctx.Done()
	defer cli.Close()

	return ii, nil
}
