package domain

import (
	"errors"
	"sync"
	"time"
)

type CallStatus string

const (
	CallRinging CallStatus = "ringing"
	CallActive  CallStatus = "active"
	CallEnded   CallStatus = "ended"
)

var (
	ErrInvalidTransition = errors.New("invalid call status transition")
	ErrCallEnded         = errors.New("call already ended")
	ErrVersionMismatch   = errors.New("version mismatch - concurrent update")
)

type Call struct {
	mu        sync.RWMutex
	ID        string
	ChannelID string
	Status    CallStatus
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCall(id, channelID string) *Call {
	now := time.Now()
	return &Call{
		ID:        id,
		ChannelID: channelID,
		Status:    CallRinging,
		Version:   1,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (c *Call) CanTransition(to CallStatus) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	switch c.Status {
	case CallRinging:
		return to == CallActive || to == CallEnded
	case CallActive:
		return to == CallEnded
	case CallEnded:
		return false
	default:
		return false
	}
}

func (c *Call) Transition(to CallStatus, expectedVersion int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Version != expectedVersion {
		return ErrVersionMismatch
	}
	if !c.canTransitionUnsafe(to) {
		if c.Status == CallEnded {
			return ErrCallEnded
		}
		return ErrInvalidTransition
	}

	c.Status = to
	c.Version++
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Call) canTransitionUnsafe(to CallStatus) bool {
	switch c.Status {
	case CallRinging:
		return to == CallActive || to == CallEnded
	case CallActive:
		return to == CallEnded
	case CallEnded:
		return false
	default:
		return false
	}
}

func (c *Call) GetStatus() CallStatus {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Status
}

func (c *Call) GetVersion() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Version
}

func (c *Call) IsActive() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Status == CallActive || c.Status == CallRinging
}

func (c *Call) Snapshot() CallSnapshot {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return CallSnapshot{
		ID:        c.ID,
		ChannelID: c.ChannelID,
		Status:    c.Status,
		Version:   c.Version,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

type CallSnapshot struct {
	ID        string
	ChannelID string
	Status    CallStatus
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
