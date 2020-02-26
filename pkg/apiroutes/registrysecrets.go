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

package apiroutes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/eclipse/codewind-installer/pkg/connections"
	"github.com/eclipse/codewind-installer/pkg/sechttp"
	"github.com/eclipse/codewind-installer/pkg/utils"
)

type (
	// RegistryResponse : The registry information
	RegistryResponse struct {
		Address  string `json:"address"`
		Username string `json:"username"`
	}

	// Registry details: The request structure to set the log level
	RegistryDetails struct {
		Level string `json:"level"`
	}
)

// GetLogLevel : Get the current log level for the PFE container
func GetRegistrySecrets(conInfo *connections.Connection, conURL string, httpClient utils.HTTPClient) ([]RegistryResponse, error) {
	req, err := http.NewRequest("GET", conURL+"/api/v1/registrysecrets", nil)
	if err != nil {
		return nil, err
	}

	resp, httpSecError := sechttp.DispatchHTTPRequest(httpClient, req, conInfo)
	if httpSecError != nil {
		return nil, httpSecError
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var registrySecrets []RegistryResponse
	err = json.Unmarshal(byteArray, &registrySecrets)
	if err != nil {
		return nil, err
	}

	return registrySecrets, nil
}

// TODO - Add POST, DELETE calls.
