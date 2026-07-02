package common

import "sync"

type Train struct {
	Id          int
	TrainLength int
	Front       int
}

type Intersection struct {
	Id       int
	LockedBy int
	Mutex    sync.Mutex
}

type Crossing struct {
	Position     int
	Intersection *Intersection
}
