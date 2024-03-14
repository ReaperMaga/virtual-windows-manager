package vw

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"io"
)

func Initialize() {
	Repository = NewMongoVirtualWindowsRepository()
	err := connectDockerClient()
	if err != nil {
		panic(err)
	}
}

var Client *docker.Client

var Context = context.Background()

func connectDockerClient() error {
	client, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	reader, err := client.ImagePull(Context, "docker.io/dockurr/windows", types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	io.Copy(io.Discard, reader)
	Client = client
	return nil
}

func StartVW(virtualWindows *VirtualWindows) error {
	containerPort := "8006" // Port to be exposed inside the container
	hostPort := "8006"      // Port to be exposed on the host
	hostIP := "0.0.0.0"     // Host IP address to bind the port
	resp, err := Client.ContainerCreate(Context, &container.Config{
		Image: "dockurr/windows",
		Tty:   false,
		ExposedPorts: map[nat.Port]struct{}{
			nat.Port(containerPort + "/tcp"): {},
		},
		Env: []string{"VERSION=" + virtualWindows.OS},
	}, &container.HostConfig{
		CapAdd:     []string{"NET_ADMIN"},
		Privileged: true,
		PortBindings: nat.PortMap{
			nat.Port(containerPort + "/tcp"): []nat.PortBinding{{
				HostIP:   hostIP,
				HostPort: hostPort,
			}},
		},
	}, nil, nil, virtualWindows.Id)
	if err != nil {
		return err
	}
	if err := Client.ContainerStart(Context, resp.ID, container.StartOptions{}); err != nil {
		return err
	}
	return nil
}

func CreateVW(name string, os string) (*VirtualWindows, error) {
	virtualWindows := &VirtualWindows{
		Id:   uuid.NewString(),
		Name: name,
		OS:   os,
	}
	err := Repository.Create(virtualWindows)
	if err != nil {
		return nil, err
	}
	return virtualWindows, nil
}

func IsVWRunning(windows *VirtualWindows) bool {
	c, err := Client.ContainerInspect(Context, windows.Id)
	if err != nil {
		return false
	}
	return c.State.Running == true
}

func StopVW(windows *VirtualWindows) error {
	err := Client.ContainerStop(Context, windows.Id, container.StopOptions{})
	if err != nil {
		return err
	}
	err = Client.ContainerRemove(Context, windows.Id, container.RemoveOptions{})
	if err != nil {
		return err
	}
	return err
}
