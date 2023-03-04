package timeseries_task

import (
	"context"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimeSeriesTask struct {
	client *mongo.Client

	collection string
	database   string

	dataGenF func(thread int, n int, start int) []any
	coll     *mongo.Collection

	ordered    bool
	arrayCount int
}

func NewTimeSeriesTask(client *mongo.Client,
	dataGenF func(thread, n int, start int) []any,
	arrayCount int, ordered bool, database, collection string) *TimeSeriesTask {
	return &TimeSeriesTask{client: client, dataGenF: dataGenF, arrayCount: arrayCount, ordered: ordered, database: database, collection: collection}
}

func (t *TimeSeriesTask) Prepare() error {
	t.coll = t.client.Database(t.database).Collection(t.collection)

	return nil
}

func (t *TimeSeriesTask) Execute(thread int, loopCounter int) error {

	ordered := false

	data := t.dataGenF(thread, t.arrayCount, loopCounter)

	_, err := t.coll.InsertMany(context.Background(), data, &options.InsertManyOptions{
		Ordered: &ordered,
	})

	if err != nil {
		return err
	}

	return nil
}

func (t *TimeSeriesTask) Done() error {
	return nil
}

func (t *TimeSeriesTask) SetUp() error {

	//metaField := "sensorid"
	//granularity := "hours"
	//
	//err := t.client.Database(t.database).Collection(t.collection).Drop(context.Background())
	//if err != nil {
	//	return err
	//}
	//err = t.client.Database(t.database).CreateCollection(context.Background(), t.collection, &options.CreateCollectionOptions{
	//	TimeSeriesOptions: &options.TimeSeriesOptions{
	//		TimeField:   "timestamp",
	//		MetaField:   &metaField,
	//		Granularity: &granularity,
	//	},
	//})
	//
	//if err != nil {
	//	return err
	//}
	return nil
}

func (t *TimeSeriesTask) TearDown() error {
	return nil
}

func GenData1(thread, n int, start int) []any {

	type Measurement struct {
		Timestamp time.Time
		SensorId  int
		Temp      int
		Col1      int
		Col2      int
		Col3      int
		Col4      int
		Col5      int
		Col6      int
	}

	var arr []any
	//startPoint := time.Now()
	for i := 0; i < n; i++ {

		//	t := startPoint.Add(time.Second * time.Duration(i*start))

		o := Measurement{time.Now(),
			thread,
			rand.Intn(100),
			rand.Intn(100),
			rand.Intn(100),
			rand.Intn(100),
			rand.Intn(100),
			rand.Intn(100),
			rand.Intn(100),
		}
		arr = append(arr, o)
	}
	return arr
}

func GenFunctions(sensorCount int) func(thread, n int, start int) []any {

	return func(thread, n int, start int) []any {

		type Measurement struct {
			Timestamp time.Time
			SensorId  int
			Temp      int
			Col1      int
			Col2      int
			Col3      int
			Col4      int
			Col5      int
			Col6      int
		}

		var arr []any
		//startPoint := time.Now()
		for i := 0; i < n; i++ {

			//	t := startPoint.Add(time.Second * time.Duration(i*start))

			o := Measurement{time.Now(),
				rand.Intn(sensorCount),
				rand.Intn(100),
				rand.Intn(100),
				rand.Intn(100),
				rand.Intn(100),
				rand.Intn(100),
				rand.Intn(100),
				rand.Intn(100),
			}
			arr = append(arr, o)
		}
		return arr
	}
}
