package task

import (
	"context"
	"github.com/stretchr/testify/assert"
	"han-mongodb/log"
	"testing"
)

func TestGetConnection(t *testing.T) {

	c, err := GetConnection(context.Background())
	assert.NoError(t, err)

	log.Info("Connection Success !!")

	c.Database("test")
}
