package elevator

import (
	"sort"
)

type ElevatorControlSystem interface {
	Status() []*Elevator
	Update(elevatorId, floor int)
	Pickup(floor, direction int)
	Step()
}

type ElevatorController struct {
	Elevators map[int]*Elevator
	MinFloor  int
	MaxFloor  int
}

//todo add floor range to config
func NewElevatorController(numberOfElevators int) (ec *ElevatorController) {
	ec = new(ElevatorController)
	ec.Elevators = map[int]*Elevator{}

	for i := 0; i < numberOfElevators; i++ {
		ec.Elevators[i] = NewElevator(i)
	}
	return
}

func (ec *ElevatorController) Status() (elevators []*Elevator) {
	for _, v := range ec.Elevators {
		elevators = append(elevators, v)
	}
	sort.Sort(ElevatorById(elevators))
	return
}

func (ec *ElevatorController) Update(elevatorId, floor int) {
	ec.Elevators[elevatorId].update(floor)
}

func (ec *ElevatorController) Pickup(floor, direction int) {
	// find the nearest elevator which stays
	elevators := make([]*Elevator, 0)

	isOnPath := func(fq *floorsQueue) bool {
		l := len(fq.floors)
		return l > 0 && ((fq.up && fq.floors[l-1] <= floor) || (!fq.up && fq.floors[l-1] >= floor))
	}

	for _, v := range ec.Elevators {
		//if any elevator is on path, than just take it
		if (v.Direction > 0 && direction > 0 && isOnPath(v.floorsUp)) ||
			(v.Direction < 0 && direction < 0 && isOnPath(v.floorsDown)) {
			v.update(floor)

			return
		}

		if v.Direction == 0 {
			elevators = append(elevators, v)
		}
	}

	// if no lift on path, take a free one
	if len(elevators) > 0 {
		sort.Sort(ElevatorByDistance{Elevators: elevators, Floor: floor})
		elevators[0].update(floor)
	} else {
		//otherwise take any
		ec.Update(0, floor)
	}
}

func (ec *ElevatorController) Step() {
	for _, v := range ec.Elevators {
		v.step()
	}
}
