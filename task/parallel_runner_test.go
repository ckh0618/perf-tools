package task

import (
	"context"
	"github.com/stretchr/testify/assert"
	"han-mongodb/task/timeseries-task"
	"testing"
)

func TestNewTimeSeriesTask(t *testing.T) {

	conn, err := GetConnection(context.Background())
	assert.NoError(t, err)

	ordered := false
	collection := "timeseriesTest"
	database := "timeseries"

	tasks := []Task{
		timeseries_task.NewTimeSeriesTask(conn, timeseries_task.GenData1, 10000, ordered, database, collection),
		timeseries_task.NewTimeSeriesTask(conn, timeseries_task.GenData1, 10000, ordered, database, collection),
		timeseries_task.NewTimeSeriesTask(conn, timeseries_task.GenData1, 10000, ordered, database, collection),
		timeseries_task.NewTimeSeriesTask(conn, timeseries_task.GenData1, 10000, ordered, database, collection),
		timeseries_task.NewTimeSeriesTask(conn, timeseries_task.GenData1, 10000, ordered, database, collection),
	}

	runner := NewParallelRunner(tasks, 1000)

	runner.Run()

}
