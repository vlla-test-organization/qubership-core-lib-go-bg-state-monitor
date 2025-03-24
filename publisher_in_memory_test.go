package blue_green_state_monitor_go

import (
	"context"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestInMemoryPublisherSubscribeUnsubscribe(t *testing.T) {
	assertions := require.New(t)
	state1 := BGStateForCurrent(NamespaceVersion{Namespace: "ns", Version: NewVersionMust("v1"), State: StateActive})
	state2 := BGStateForCurrent(NamespaceVersion{Namespace: "ns", Version: NewVersionMust("v2"), State: StateActive})
	publisher, err := NewInMemoryPublisher(state1)
	assertions.NoError(err)
	assertions.NotNil(publisher)

	assertions.Equal(state1, publisher.GetState())

	ctx1, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	publisher.Subscribe(ctx1, func(s BlueGreenState) {
		if s == state1 {
			wg.Done()
		} else if s == state2 {
			wg2.Done()
		}
	})
	assertions.True(waitWG(3*time.Second, wg))

	cancel() // since cancel is async operation, need to add subscribe one more time to make sure that un-subscription has taken place

	wg3 := &sync.WaitGroup{}
	wg3.Add(1)
	wg4 := &sync.WaitGroup{}
	wg4.Add(1)
	ctx2, cancel := context.WithCancel(context.Background())
	defer cancel()

	publisher.Subscribe(ctx2, func(s BlueGreenState) {
		if s == state1 {
			wg3.Done()
		} else if s == state2 {
			wg4.Done()
		}
	})
	assertions.True(waitWG(3*time.Second, wg3))

	publisher.SetState(state2)

	assertions.Equal(state2, publisher.GetState())

	assertions.True(waitWG(3*time.Second, wg4))
	assertions.False(waitWG(1*time.Second, wg2))
}
