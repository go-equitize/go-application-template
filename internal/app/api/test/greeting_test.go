package test

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func (suite *TestSuite) TestAPI_Greeting_Successfully() {
	name := "TestAPI_Greeting_Successfully"
	suite.T().Log(name)

	// call API
	method := "GET"
	endpoint := "/api/v1/greetings"
	req, _ := http.NewRequest(method, endpoint, nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)

	// assert response
	suite.Equal(`{"code":0,"message":"Successfully","data":"Hello World"}`, string(responseData))
}
