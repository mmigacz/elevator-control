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
	for _, v := range ec.Elevators {
		if v.Direction == 0 {
			elevators = append(elevators, v)
		}
	}

	// and start it
	if len(elevators) > 0 {
		sort.Sort(ElevatorByDistance{Elevators: elevators, Floor: floor})

		elevators[0].update(floor)
	}
}

func (ec *ElevatorController) Step() {
	for _, v := range ec.Elevators {
		v.step()
	}
}
