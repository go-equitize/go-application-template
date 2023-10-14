package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildAPIServer_Successfully(t *testing.T) {
	name := "TestBuildAPIServer_Successfully"
	t.Log(name)

	server := NewAPIServer()
	assert.NotNil(t, server)
}
