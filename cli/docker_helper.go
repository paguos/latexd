package main

import (
	"bufio"
	"context"
	"fmt"
	// "io"
	"os"
	
	log "github.com/sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)


func runContainer(cmd Command, imageName string) {
	log.Info("Executing '", cmd.name, "' ...")
	log.SetLevel(log.DebugLevel)
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	_ = reader

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Creating container ...")
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   cmd.args,
		Tty:   true,
		AttachStdout: true,
        AttachStderr: true,
	}, &container.HostConfig{
        Mounts: []mount.Mount{
            {
                Type:   mount.TypeBind,
                Source: pwd,
                Target: "/docs",
            },
        },
    }, nil, nil, "")
	if err != nil {
		panic(err)
	}

	log.Info("Running container ...")

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	go func() {
		reader, err := cli.ContainerLogs(context.Background(), resp.ID, types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     true,
			Timestamps: false,
		})
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	log.Info("'", cmd.name, "' done!")
}