package evmredemptionfulfillmentgrants

import "github.com/pkg/errors"

// StatusT is the lifecycle of a reusable, narrow redemption fulfillment grant.
type StatusT string

const (
	StatusActive     StatusT = "active"
	StatusRevoked    StatusT = "revoked"
	StatusSuperseded StatusT = "superseded"
)

func AllStatusT() []StatusT {
	return []StatusT{StatusActive, StatusRevoked, StatusSuperseded}
}

func (value StatusT) IsValid() error {
	switch value {
	case StatusActive, StatusRevoked, StatusSuperseded:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (value StatusT) String() string { return string(value) }
