package elevator

import (
	"container/heap"
	"fmt"
	"math"
)

type pickup struct {
	floorNumber int
	direction   int
}

type Elevator struct {
	ElevatorId  int
	FloorNumber int
	Direction   int
	Open        bool
	floorsUp    *floorsQueue
	floorsDown  *floorsQueue
}

func NewElevator(id int) *Elevator {
	return &Elevator{
		ElevatorId: id,
		floorsUp:   newFloorsQueue(true),
		floorsDown: newFloorsQueue(false),
	}
}

func (e *Elevator) Str() string {
	direction := "none"
	if e.Direction > 0 {
		direction = "up"
	} else if e.Direction < 0 {
		direction = "down"
	}

	return fmt.Sprintf("id:[%d], floor:[%d], open:[%t] direction:[%s], upq:[%d], downq:[%d]", e.ElevatorId,
		e.FloorNumber, e.Open, direction, e.floorsUp.Len(), e.floorsDown.Len())
}

func (e *Elevator) update(goToFloor int) {
	if e.FloorNumber > goToFloor {
		e.floorsDown.Push(goToFloor)
	} else if e.FloorNumber < goToFloor {
		e.floorsUp.Push(goToFloor)
	}
	e.updateDirection()
}

func (e *Elevator) step() {
	move := func(fq *floorsQueue) {
		l := len(fq.floors)
		if l > 0 && fq.floors[l-1] == e.FloorNumber {
			e.FloorNumber = heap.Pop(fq).(int)
			e.Open = true
		} else {
			e.Open = false
		}
	}

	if e.Direction > 0 {
		e.FloorNumber = e.FloorNumber + 1
		move(e.floorsUp)
	} else if e.Direction < 0 {
		e.FloorNumber = e.FloorNumber - 1
		move(e.floorsDown)
	}
	e.updateDirection()
}

func (e *Elevator) distanceToFloor(floorId int) int {
	return int(math.Abs(float64(e.FloorNumber) - float64(floorId)))
}

func (e *Elevator) updateDirection() {
	if e.floorsUp.Len() == 0 && e.floorsDown.Len() > 0 {
		e.Direction = -1
	} else if e.floorsUp.Len() > 0 && e.floorsDown.Len() == 0 {
		e.Direction = 1
	} else {
		e.Direction = 0
	}
}

type ElevatorById []*Elevator

func (a ElevatorById) Len() int           { return len(a) }
func (a ElevatorById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ElevatorById) Less(i, j int) bool { return a[i].ElevatorId < a[j].ElevatorId }

type ElevatorByDistance struct {
	Elevators []*Elevator
	Floor     int
	Direction int
}

func (a ElevatorByDistance) Len() int { return len(a.Elevators) }
func (a ElevatorByDistance) Swap(i, j int) {
	a.Elevators[i], a.Elevators[j] = a.Elevators[j], a.Elevators[i]
}
func (a ElevatorByDistance) Less(i, j int) bool {
	return a.Elevators[i].distanceToFloor(a.Floor) < a.Elevators[j].distanceToFloor(a.Floor)
}
