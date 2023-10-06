package api

import (
	"net/http"
	"net/http/httptest"
)

func (suite *TestSuite) TestAPI_HealthLive() {
	suite.T().Log("TestAPI_HealthLive")

	method := "GET"
	endpoint := "/health/live"
	req, _ := http.NewRequest(method, endpoint, nil)
	w := httptest.NewRecorder()
	testServer.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TestSuite) TestAPI_HealthReady() {
	suite.T().Log("TestAPI_HealthReady")

	method := "GET"
	endpoint := "/health/ready"
	req, _ := http.NewRequest(method, endpoint, nil)
	w := httptest.NewRecorder()
	testServer.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}
