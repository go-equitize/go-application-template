package test

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func (suite *TestSuite) TestAPI_Healthcheck_Successfully() {
	name := "TestAPI_Healthcheck_Successfully"
	suite.T().Log(name)

	// call API
	method := "GET"
	endpoint := "/health"
	req, _ := http.NewRequest(method, endpoint, nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)

	// assert response
	suite.Equal(`"OK"`, string(responseData))
}
