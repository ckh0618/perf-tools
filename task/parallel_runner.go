package task

import (
	"fmt"
	"han-mongodb/timerstat"
	"sync"
	"time"
)

type executionResult struct {
	ThreadID int
	Duration time.Duration
}

type ParallelRunner struct {
	TotalCount int
	Tasks      []Task

	myTimer         timerstat.TimeStore
	countPerThread  int
	resultChan      chan executionResult
	waitGroup       sync.WaitGroup
	reportWaitGroup sync.WaitGroup
}

func NewParallelRunner(tasks []Task, totalCount int) *ParallelRunner {

	taskLen := len(tasks)
	countPerTask := totalCount / taskLen

	return &ParallelRunner{
		TotalCount:      totalCount,
		Tasks:           tasks,
		countPerThread:  countPerTask,
		resultChan:      make(chan executionResult, 8192*100),
		waitGroup:       sync.WaitGroup{},
		reportWaitGroup: sync.WaitGroup{},
	}

}

func (p *ParallelRunner) Run() error {

	p.myTimer.Initialize()

	p.reportWaitGroup.Add(1)
	go p.receiveResult()

	if err := p.Tasks[0].SetUp(); err != nil {
		return err
	}

	for i, task := range p.Tasks {
		p.waitGroup.Add(1)
		go p.RunTask(i, task)
	}
	p.waitGroup.Wait()
	close(p.resultChan)
	p.reportWaitGroup.Wait()

	if err := p.Tasks[0].TearDown(); err != nil {
		return err
	}

	return nil
}

func (p *ParallelRunner) RunTask(thread int, task Task) error {

	start := thread * p.countPerThread
	end := start + p.countPerThread

	task.Prepare()
	for i := start; i < end; i++ {
		start := time.Now()
		task.Execute(thread, i)
		elapsed := time.Now().Sub(start)
		p.resultChan <- executionResult{
			ThreadID: thread,
			Duration: elapsed,
		}
	}
	task.Done()
	p.waitGroup.Done()
	return nil
}

func (p *ParallelRunner) receiveResult() {

	for d := range p.resultChan {
		p.myTimer.Add(d.ThreadID, d.Duration)
	}
	p.generateReport()
}

func (p *ParallelRunner) generateReport() {
	defer p.reportWaitGroup.Done()
	totalTps := float64(0)
	for i, _ := range p.Tasks {
		stat, err := p.myTimer.Stat(i)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("ID[%3d] Min[%10d] Max[%10d] Avg[%10.2f] Tps[%10.2f] 99P [%10d]\n",
			stat.ID, stat.Min, stat.Max, stat.Avg, stat.Tps, stat.P99)
		totalTps = totalTps + stat.Tps
	}
	fmt.Println("Total TPs : ", totalTps)
}
