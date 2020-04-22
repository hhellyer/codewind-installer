/*******************************************************************************
 * Copyright (c) 2019 IBM Corporation and others.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 *     IBM Corporation - initial API and implementation
 *******************************************************************************/

package docker

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

// DockerClient requires all the functions called on the docker client
type DockerClient interface {
	ImagePull(ctx context.Context, image string, imagePullOptions types.ImagePullOptions) (io.ReadCloser, error)
	ImageList(ctx context.Context, imageListOptions types.ImageListOptions) ([]types.ImageSummary, error)
	ContainerList(ctx context.Context, containerListOptions types.ContainerListOptions) ([]types.Container, error)
	ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
	ContainerStop(ctx context.Context, containerID string, timeout *time.Duration) error
	ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error
	DaemonHost() string
	DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registry.DistributionInspect, error)
}

// NewDockerClient creates a new client for the docker API
func NewDockerClient() (DockerClient, *DockerError) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.30"))
	if err != nil {
		return nil, &DockerError{errOpClientCreate, err, err.Error()}
	}
	// Use to confirm client.WithVersion("1.30") overrides client.FromEnv version.
	// fmt.Printf("Client version is %s\n", dockerClient.ClientVersion())
	// fmt.Printf("DaemonHost is %s\n", dockerClient.DaemonHost())
	return dockerClient, nil
}

// UsingLocalDockerHost returns true if we are using the default local docker host.
func UsingLocalDockerHost(dockerClient DockerClient) bool {
	return dockerClient.DaemonHost() == client.DefaultDockerHost
}
