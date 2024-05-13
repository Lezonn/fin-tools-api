package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	assert.NotNil(t, db)
}
