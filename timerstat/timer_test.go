package timerstat

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/model"
	"testing"
)

func TestTimer(t *testing.T) {
	var store TimeStore
	store.Initialize()
	store.Start(1)
	store.End(1)
}

func TestTimer2(t *testing.T) {
	var store TimeStore

	store.Initialize()

	for j := 0; j < 10000; j++ {
		for i := 0; i < 100; i++ {
			store.Start(i)
		}

		for i := 0; i < 100; i++ {
			store.End(i)
		}
	}

	for i := 0; i < 100; i++ {
		stat, err := store.Stat(i)

		_, err = store.Stat(i)
		if err != nil {
			t.Fail()
		}

		fmt.Println(stat)
	}
}

func Test2(t *testing.T) {
	tr := bus.GetBus()
	channel := "hello-world"

	// create new channel
	tr.GetChannelManager().CreateChannel(channel)

	// listen for a single request on 'hello-world'
	requestHandler, _ := tr.ListenRequestStream(channel)

	// define request handler logic
	requestHandler.Handle(
		func(msg *model.Message) {
			resp := msg.Payload.(string)
			fmt.Printf("\\nHello: %s\\n", resp)

			// send a response back.
			tr.SendResponseMessage(channel, resp+" Doodly", msg.DestinationId)
		},
		func(err error) {
			// something went wrong...
			fmt.Println(err)
		},
	)

	// send a request to 'hello-world' and handle a single response
	responseHandler, _ := tr.RequestOnce(channel, "Howdy")

	// define response handler logic
	responseHandler.Handle(
		func(msg *model.Message) {
			fmt.Printf("World: %s\\n", msg.Payload.(string))
		},
		func(err error) {
			// something went wrong...
		})

	// fire the request.
	responseHandler.Fire()
}
