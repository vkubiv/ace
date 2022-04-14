/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hyperledger/aries-framework-go/pkg/common/log"

	"github.com/trustbloc/ace/pkg/restapi/handler"
)

var logger = log.New("healthcheck")

// API endpoints.
const (
	healthCheckEndpoint = "/healthcheck"
)

type healthCheckResp struct {
	Status      string    `json:"status"`
	CurrentTime time.Time `json:"currentTime"`
}

// New returns CreateCredential instance.
func New() *Operation {
	return &Operation{}
}

// Operation defines handlers for rp operations.
type Operation struct{}

// GetRESTHandlers get all controller API handler available for this service.
func (o *Operation) GetRESTHandlers() []handler.Handler {
	return []handler.Handler{
		handler.NewHTTPHandler(healthCheckEndpoint, http.MethodGet, o.healthCheckHandler),
	}
}

func (o *Operation) healthCheckHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)

	err := json.NewEncoder(rw).Encode(&healthCheckResp{
		Status:      "success",
		CurrentTime: time.Now(),
	})
	if err != nil {
		logger.Errorf("healthcheck response failure, %s", err)
	}
}
