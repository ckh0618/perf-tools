package timerstat

import (
	"errors"
	"sort"
	"time"
)

type timerList []time.Duration

// TimeStore is a kind of StopWatch class.
type TimeStore struct {
	data      map[int]timerList
	startTime map[int]time.Time
}

// TimeStat is Time statistics report per thread
type TimeStat struct {
	ID  int     `json:"id"`
	Min int64   `json:"min"`
	Max int64   `json:"max"`
	Avg float64 `json:"avg"`
	Tps float64 `json:"tps"`
	P99 int64   `json:"p99"`
}

// Initialize TimerStore
func (t *TimeStore) Initialize() {
	t.data = make(map[int]timerList)
	t.startTime = make(map[int]time.Time)
}

// Start Timer
func (t *TimeStore) Start(id int) {
	t.startTime[id] = time.Now()
}

// End evaluate and store duration
func (t *TimeStore) End(id int) time.Duration {
	endTime := time.Now()
	elapsed := endTime.Sub(t.startTime[id])

	t.Add(id, elapsed)

	return elapsed
}

// Add new point
func (t *TimeStore) Add(id int, val time.Duration) {
	t.data[id] = append(t.data[id], val)
}

func (t timerList) Len() int {
	return len(t)
}

func (t timerList) Less(i, j int) bool {
	return t[i].Nanoseconds() < t[j].Nanoseconds()
}

func (t timerList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Stat method
func (t *TimeStore) Stat(id int) (TimeStat, error) {

	var total float64
	var sReturn TimeStat

	data, ok := t.data[id]
	totalLen := data.Len()

	if !ok {
		return sReturn, errors.New("Empty Result")
	}

	sort.Sort(t.data[id])

	sReturn.ID = id
	sReturn.Min = data[0].Nanoseconds() / 1000
	sReturn.Max = data[totalLen-1].Nanoseconds() / 1000

	for _, val := range data {
		total = total + ((float64)(val.Nanoseconds()) / 1000)
	}
	sReturn.Avg = total / (float64)(totalLen)

	s99Index := totalLen - (totalLen / 100)
	s999Index := totalLen - (totalLen / 1000)

	if s99Index >= totalLen {
		s99Index = totalLen - 1
	}

	if s999Index >= totalLen {
		s999Index = totalLen - 1
	}

	sReturn.P99 = data[s99Index-1].Nanoseconds() / 1000

	sReturn.Tps = 1000000 / sReturn.Avg

	return sReturn, nil
}
