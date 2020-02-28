/*******************************************************************************
 * Copyright (c) 2020 IBM Corporation and others.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 *     IBM Corporation - initial API and implementation
 *******************************************************************************/

package actions

import (
	"net/http"
	"os"
	"strings"

	"github.com/eclipse/codewind-installer/pkg/apiroutes"
	"github.com/eclipse/codewind-installer/pkg/config"
	"github.com/eclipse/codewind-installer/pkg/connections"
	"github.com/eclipse/codewind-installer/pkg/utils"
	"github.com/urfave/cli"
)

// GetRegistrySecrets : Retrieve docker registry secrets.
func GetRegistrySecrets(c *cli.Context) {
	conInfo, conURL := getConnectionDetailsOrExit(c)

	registrySecrets, err := apiroutes.GetRegistrySecrets(conInfo, conURL, http.DefaultClient)
	if err != nil {
		registryErr := &RegistryError{errOpListRegistries, err, err.Error()}
		HandleRegistryError(registryErr)
		os.Exit(1)
	}
	utils.PrettyPrintJSON(registrySecrets)
}

// AddRegistrySecret : Set a docker registry secret.
func AddRegistrySecret(c *cli.Context) {
	conInfo, conURL := getConnectionDetailsOrExit(c)

	address := strings.TrimSpace(c.String("address"))
	username := strings.TrimSpace(c.String("username"))
	password := strings.TrimSpace(c.String("password"))

	registrySecrets, err := apiroutes.AddRegistrySecret(conInfo, conURL, http.DefaultClient, address, username, password)
	if err != nil {
		registryErr := &RegistryError{errOpAddRegistry, err, err.Error()}
		HandleRegistryError(registryErr)
		os.Exit(1)
	}

	// if local connection....
	fmt.Printf("conInfo.ID = %v\n", conInfo.ID)
	if conInfo.ID == "local" {
		// TODO
		// - Do docker login here.
		dockerClient, dockerErr := utils.NewDockerClient()
		if dockerErr != nil {
			HandleDockerError(dockerErr)
			os.Exit(1)
		}
		dockerErr = utils.DockerLogin(dockerClient, address, username, password)
		// TODO
		// - Write to keychain.
		fmt.Println("Local connection, adding secret to keychain entry.")

	}

	utils.PrettyPrintJSON(registrySecrets)
}

// RemoveRegistrySecret : Delete a docker registry secret.
func RemoveRegistrySecret(c *cli.Context) {
	conInfo, conURL := getConnectionDetailsOrExit(c)

	address := strings.TrimSpace(c.String("address"))

	registrySecrets, err := apiroutes.RemoveRegistrySecret(conInfo, conURL, http.DefaultClient, address)
	if err != nil {
		registryErr := &RegistryError{errOpRemoveRegistry, err, err.Error()}
		HandleRegistryError(registryErr)
		os.Exit(1)
	}

	// TODO - Remove secret from our keychain entry.
	// (But don't logout of docker.)
	utils.PrettyPrintJSON(registrySecrets)
}

func getConnectionDetailsOrExit(c *cli.Context) (*connections.Connection, string) {
	connectionID := strings.TrimSpace(strings.ToLower(c.String("conid")))

	conInfo, conInfoErr := connections.GetConnectionByID(connectionID)
	if conInfoErr != nil {
		HandleConnectionError(conInfoErr)
		os.Exit(1)
	}

	conURL, conErr := config.PFEOriginFromConnection(conInfo)
	if conErr != nil {
		HandleConfigError(conErr)
		os.Exit(1)
	}
	return conInfo, conURL
}
