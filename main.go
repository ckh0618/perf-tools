package main

import (
	"context"
	"flag"
	"fmt"
	"han-mongodb/task"
	"han-mongodb/task/timeseries-task"
)

type CommandArgs struct {
	TotalRun         int
	ArrayPerEachCall int
	Threads          int
	Ordered          bool
	Collection       string
	Database         string
	GenFunction      string
}

var (
	mapFunc = map[string]func(thread int, n int, start int) []any{
		"GenData":              timeseries_task.GenData1,
		"GenData2-Sensor100":   timeseries_task.GenFunctions(100),
		"GenData2-Sensor1000":  timeseries_task.GenFunctions(1000),
		"GenData2-Sensor10000": timeseries_task.GenFunctions(10000),
	}
)

func main() {

	collection := flag.String("collection", "timeseriesData", "collection name")
	database := flag.String("database", "timeseries", "database name")
	totalRun := flag.Int("totalrun", 1000, "total test count")
	threads := flag.Int("thread", 4, "number_of_threads")
	arrayPerEachCall := flag.Int("array", 10000, "array  per each call")
	ordered := flag.Bool("ordered", false, "ordered")
	genFunction := flag.String("genfunc", "GenData", "arraydata")

	flag.Parse()

	c := CommandArgs{
		TotalRun:         *totalRun,
		ArrayPerEachCall: *arrayPerEachCall,
		Threads:          *threads,
		Ordered:          *ordered,
		Collection:       *collection,
		Database:         *database,
		GenFunction:      *genFunction,
	}

	fmt.Println(c)

	conn, err := task.GetConnection(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	var tasks []task.Task

	for i := 0; i < c.Threads; i++ {
		t := timeseries_task.NewTimeSeriesTask(conn, mapFunc[c.GenFunction], c.ArrayPerEachCall, c.Ordered, c.Database, c.Collection)
		tasks = append(tasks, t)
	}

	runner := task.NewParallelRunner(tasks, c.TotalRun)

	if err := runner.Run(); err != nil {
		fmt.Println(err)
		return
	}

	return

}
