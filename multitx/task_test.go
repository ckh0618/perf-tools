package multitx

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"han-mongodb/log"
	"testing"
)

func TestTask(t *testing.T) {

	log.Info("An Information Sections")
}

func TestGetConnection(t *testing.T) {
	uri := "mongodb://localhost:27017"

	c, err := GetConnection(context.Background(), uri)
	assert.NoError(t, err)

	log.Info("Connection Success !!")

	c.Database("test")
}
