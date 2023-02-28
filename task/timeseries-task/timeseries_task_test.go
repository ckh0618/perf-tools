package timeseries_task

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"han-mongodb/multitx"
	"testing"
)

func TestTimeSeriesTask_Execute(t *testing.T) {

	conn, err := multitx.GetConnection(context.Background(),
		"mongodb://localhost:27017")

	assert.NoError(t, err)

	t1 := NewTimeSeriesTask(conn, GenData1, 100, false, "timeseries", "timeseriesTest")
	t1.Prepare()
	t1.Execute(0, 0)
}

func TestGenerateData(t *testing.T) {

	c := GenData1(0, 100, 0)
	c2 := GenData1(0, 100, 100)

	for _, data := range c {
		fmt.Println(data)
	}

	for _, data := range c2 {
		fmt.Println(data)
	}

}
