package task

type Task interface {
	SetUp() error
	Prepare() error
	Execute(thread, loopCounter int) error
	Done() error
	TearDown() error
}
