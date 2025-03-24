package blue_green_state_monitor_go

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestStateFromShort(t *testing.T) {
	assertions := require.New(t)
	tests := map[string]State{
		"a": StateActive,
		"i": StateIdle,
		"c": StateCandidate,
		"l": StateLegacy,
	}
	for short, s := range tests {
		state, err := StateFromShort(short)
		assertions.NoError(err)
		assertions.Equal(s, state)
		shortName := state.ShortString()
		assertions.Equal(short, shortName)
	}
}

func TestBlueGreenStateString(t *testing.T) {
	assertions := require.New(t)
	bgState := BlueGreenState{
		Current:    NamespaceVersion{Namespace: "test-origin", Version: &Version{Value: "v1", IntValue: 1}, State: StateActive},
		Sibling:    &NamespaceVersion{Namespace: "test-peer", Version: &Version{Value: "v2", IntValue: 2}, State: StateCandidate},
		UpdateTime: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	str := bgState.String()
	assertions.Equal("Current: {Namespace: test-origin, State: ACTIVE, Version: v1}, Sibling: {Namespace: test-peer, State: CANDIDATE, Version: v2}, UpdateTime: 2000-01-01 00:00:00 +0000 UTC", str)
}

func TestBlueGreenStateEqual(t *testing.T) {
	assertions := require.New(t)
	bgState1 := &BlueGreenState{
		Current:    NamespaceVersion{Namespace: "test-origin", Version: &Version{Value: "v1", IntValue: 1}, State: StateActive},
		Sibling:    &NamespaceVersion{Namespace: "test-peer", Version: &Version{Value: "v2", IntValue: 2}, State: StateCandidate},
		UpdateTime: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	bgState2 := &BlueGreenState{
		Current:    NamespaceVersion{Namespace: "test-origin", Version: &Version{Value: "v1", IntValue: 1}, State: StateActive},
		Sibling:    &NamespaceVersion{Namespace: "test-peer", Version: &Version{Value: "v2", IntValue: 2}, State: StateCandidate},
		UpdateTime: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	assertions.True(bgState1.Equal(bgState2))
}
