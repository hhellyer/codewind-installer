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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/eclipse/codewind-installer/pkg/connections"
	"github.com/eclipse/codewind-installer/pkg/security"
	"github.com/stretchr/testify/assert"
)

// func Test_GetRegistrySecrets(t *testing.T) {
// 	allLevels := []string{"error", "warn", "info", "debug", "trace"}
// 	loggingResponse := LoggingResponse{
// 		CurrentLevel: "debug",
// 		DefaultLevel: "info",
// 		AllLevels:    allLevels,
// 	}
// 	mockConnection := connections.Connection{ID: "local"}
// 	t.Run("success case - correct logging levels are returned", func(t *testing.T) {
// 		jsonResponse, err := json.Marshal(loggingResponse)
// 		if err != nil {
// 			t.Fail()
// 		}
// 		body := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
// 		mockClient := &security.ClientMockAuthenticate{StatusCode: http.StatusOK, Body: body}
// 		gotSecrets, err := GetRegistrySecrets(&mockConnection, "mockURL", mockClient)
// 		if err != nil {
// 			t.Fail()
// 		}
// 		assert.Equal(t, gotSecrets, loggingResponse)
// 	})
// 	t.Run("error case - returns error when PFE status code non 200", func(t *testing.T) {
// 		mockClientFalse := &security.ClientMockAuthenticate{StatusCode: http.StatusNotFound, Body: nil}
// 		_, err := GetIgnoredPaths(mockClientFalse, &mockConnection, "nodejs", "dummyurl")
// 		assert.Error(t, err)
// 	})
// }

func Test_SetRegistrySecret(t *testing.T) {

	t.Run("success case - returns nil error when PFE status code 200", func(t *testing.T) {
		expectedRegistrySecrets := []RegistryResponse{RegistryResponse{Address: "testdockerregistry", Username: "testuser"}}
		jsonResponse, err := json.Marshal(expectedRegistrySecrets)
		if err != nil {
			t.Fail()
		}
		body := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
		mockClient := &security.ClientMockAuthenticate{StatusCode: http.StatusOK, Body: body}
		mockConnection := connections.Connection{ID: "local"}
		actualRegistrySecrets, err := SetRegistrySecret(&mockConnection, "mockURL", mockClient, "testdockerregistry", "testuser", "testpassword")
		assert.Nil(t, err)
		assert.Equal(t, expectedRegistrySecrets, actualRegistrySecrets)
	})
	// t.Run("error case - returns error when PFE status code non 200", func(t *testing.T) {
	// 	mockConnection := connections.Connection{ID: "local"}
	// 	mockClientFalse := &security.ClientMockAuthenticate{StatusCode: http.StatusNotFound, Body: nil}
	// 	_, err := GetIgnoredPaths(mockClientFalse, &mockConnection, "nodejs", "dummyurl")
	// 	assert.Error(t, err)
	// })
}
