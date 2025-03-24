# blue-green-state-monitor-go

## Usage with fixed consul token

```go
package main

import (
	"context"
	"fmt"
	"github.com/netcracker/qubership-core-lib-go-bg-state-monitor/v2"
)

func demo(ctx context.Context, consulUrl string, consulToken string) {
	publisher, err := blue_green_state_monitor_go.NewPublisher(ctx,
		consulUrl, "test-namespace", func(ctx context.Context) (string, error) { return consulToken, nil })
	if err != nil {
		panic(err)
	}
	ctxS, cancel := context.WithCancel(ctx)
	publisher.Subscribe(ctxS, func(state blue_green_state_monitor_go.BlueGreenState) {
		fmt.Printf("state=%s\n", state)
	})
	state := publisher.GetState()
	fmt.Printf("state=%s\n", state.String())
	cancel()
}
```

## Usage with m2m consul token provider
```go
package main

import (
	"context"
	"fmt"
	bgState "github.com/netcracker/qubership-core-lib-go-bg-state-monitor/v2"
)

func demo(ctx context.Context, consulUrl string, consulToken string, namespace string) {
	publisher, err := bgState.NewPublisher(ctx,
		consulUrl, namespace, func(ctx context.Context) (string, error) { return consulToken, nil })
	if err != nil {
		panic(err)
	}
	ctxS, cancel := context.WithCancel(ctx)
	publisher.Subscribe(ctxS, func(state bgState.BlueGreenState) {
		fmt.Printf("state=%s\n", state)
	})
	state := publisher.GetState()
	fmt.Printf("state=%s\n", state.String())
	cancel()
}
```