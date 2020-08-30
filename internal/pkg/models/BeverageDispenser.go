package models

import (
	"sync"
)

// BeverageDispenser corresponds to an beverage outlet,
// outlets are created as per the configuration machine.outlets.count_n
// a dispenser can only serve one beverage at a time
// it is locked for a beverage request and is released once the beverage is served
type BeverageDispenser struct {
	id            int
	isIdle        bool
	statusChecker *sync.Mutex // mutex lock to lock, release dispenser
}

func NewBeverageDispenser(id int) BeverageDispenser {
	return BeverageDispenser{
		id:            id,
		isIdle:        true,
		statusChecker: &sync.Mutex{},
	}
}

func (bd *BeverageDispenser) GetId() int {
	return bd.id
}

// acquire an idle dispenser
func (bd *BeverageDispenser) Acquire() bool {
	bd.statusChecker.Lock()
	defer bd.statusChecker.Unlock()
	if bd.isIdle {
		bd.isIdle = false
		return true
	}
	return false
}

// release the dispenser
func (bd *BeverageDispenser) Release() {
	bd.statusChecker.Lock()
	defer bd.statusChecker.Unlock()
	bd.isIdle = true
}
