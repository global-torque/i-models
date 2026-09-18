package evmredemptiongrantcommands

import "github.com/pkg/errors"

type KindT string

const (
	KindPrepare       KindT = "prepare"
	KindConfirm       KindT = "confirm"
	KindRevokePrepare KindT = "revoke_prepare"
	KindRevokeConfirm KindT = "revoke_confirm"
)

func AllKindT() []KindT {
	return []KindT{KindPrepare, KindConfirm, KindRevokePrepare, KindRevokeConfirm}
}
func (value KindT) IsValid() error {
	for _, candidate := range AllKindT() {
		if value == candidate {
			return nil
		}
	}
	return errors.New("enum is not valid")
}
func (value KindT) String() string { return string(value) }

type StateT string

const (
	StatePrepared       StateT = "prepared"
	StateExecuting      StateT = "executing"
	StateReconciling    StateT = "reconciling"
	StateCompleted      StateT = "completed"
	StateFailedTerminal StateT = "failed_terminal"
)

func AllStateT() []StateT {
	return []StateT{StatePrepared, StateExecuting, StateReconciling, StateCompleted, StateFailedTerminal}
}
func (value StateT) IsValid() error {
	for _, candidate := range AllStateT() {
		if value == candidate {
			return nil
		}
	}
	return errors.New("enum is not valid")
}
func (value StateT) String() string { return string(value) }
