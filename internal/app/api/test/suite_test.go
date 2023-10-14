package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"

	"go-template/internal/app/api/component"
	"go-template/internal/app/api/route"
	"go-template/internal/pkg/config"
)

var (
	testRouter *gin.Engine
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupSuite() {
	fmt.Println("============ Start running tests for package `api`... ============")

	// Load config
	err := config.Load("../../../pkg/config/file/test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Init components
	err = component.InitComponents()
	if err != nil {
		log.Fatal(err)
	}
	// HTTP mocks
	//httpmock.ActivateNonDefault(...)

	// setup router
	if testRouter == nil {
		gin.SetMode(config.Instance().Http.Mode)
		testRouter = gin.New()
		route.Register(testRouter)
	}
}

func (suite *TestSuite) SetupTest() {
	// Load config
	err := config.Load("../config/file/test.yaml")
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
