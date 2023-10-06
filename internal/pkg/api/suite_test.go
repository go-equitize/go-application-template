package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"testing"

	"go-template/internal/pkg/components"
	"go-template/internal/pkg/config"
)

var testServer *gin.Engine

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupSuite() {
	fmt.Println("-------------- Init components before test --------------")

	err := config.Load("../config/files/test.yaml")
	if err != nil {
		panic(err)
	}

	_ = components.InitComponents()

	if testServer == nil {
		gin.SetMode(config.Instance().Http.Mode)
		testServer = gin.New()
	}
	AddHealthRouter(testServer)
}

func (suite *TestSuite) SetupTest() {
	// Load config
	err := config.Load("../config/files/test.yaml")
	if err != nil {
		panic(err)
	}

	httpmock.Reset()
}

func (suite *TestSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
