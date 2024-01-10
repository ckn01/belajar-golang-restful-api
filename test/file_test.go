package test

import (
	"testing"

	"github.com/ckn01/belajar-golang-restful-api/simple"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databaes")
	assert.NotNil(t, connection)

	cleanup()
}
