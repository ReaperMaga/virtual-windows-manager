package vw

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/volume"
	docker "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"regexp"
	"strconv"
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
	containerPort := "8006"
	hostPort := strconv.Itoa(virtualWindows.Port)
	hostIP := "0.0.0.0"
	resp, err := Client.ContainerCreate(Context, &container.Config{
		Image: "dockurr/windows",
		Tty:   false,
		ExposedPorts: map[nat.Port]struct{}{
			nat.Port(containerPort + "/tcp"): {},
		},
		Env:     []string{"VERSION=" + virtualWindows.OS},
		Volumes: map[string]struct{}{virtualWindows.Id + ":/storage": {}},
	}, &container.HostConfig{
		CapAdd:     []string{"NET_ADMIN"},
		Privileged: true,
		PortBindings: nat.PortMap{
			nat.Port(containerPort + "/tcp"): []nat.PortBinding{{
				HostIP:   hostIP,
				HostPort: hostPort,
			}},
		},
		Binds: []string{virtualWindows.Id + ":/storage"},
	}, nil, nil, virtualWindows.Id)
	if err != nil {
		return err
	}
	if err := Client.ContainerStart(Context, resp.ID, container.StartOptions{}); err != nil {
		return err
	}
	return nil
}

func GetVWLogs(virtualWindows *VirtualWindows) (string, error) {
	reader, err := Client.ContainerLogs(Context, virtualWindows.Id, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       "40",
	})
	if err != nil {
		return "", err
	}
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			return
		}
	}(reader)
	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s-;#_!?*(){}'\"`$@+=,.<>/&%]")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(string(content), ""), nil
}

func randPort() int {
	return rand.Intn(9999-1000) + 1000
}

func CreateVW(name string, os string) (*VirtualWindows, error) {
	virtualWindows := &VirtualWindows{
		Id:   uuid.NewString(),
		Name: name,
		OS:   os,
		Port: randPort(),
	}
	err := Repository.Create(virtualWindows)
	if err != nil {
		return nil, err
	}
	err = createVolume(virtualWindows)
	if err != nil {
		fmt.Println("There was an error while trying to create a volume: ", err)
		return nil, err
	}
	return virtualWindows, nil
}

func createVolume(windows *VirtualWindows) error {
	_, err := Client.VolumeCreate(Context, volume.CreateOptions{
		Name: windows.Id,
	})
	if err != nil {
		return err
	}
	return nil
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
