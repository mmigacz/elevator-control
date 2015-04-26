package elevator

type floorsQueue struct {
	up     bool
	floors []int
}

func newFloorsQueue(up bool) *floorsQueue {
	return &floorsQueue{up: up}
}

func (f *floorsQueue) containsFloor(floorId int) bool {
	for _, nr := range f.floors {
		if nr == floorId {
			return true
		}
	}
	return false
}

// floorsQueue implements heap.Interface and holds floors.
func (f floorsQueue) Len() int { return len(f.floors) }

func (f floorsQueue) Less(i, j int) bool {
	if f.up {
		return f.floors[i] < f.floors[j]
	} else {
		return f.floors[i] > f.floors[j]
	}
}

func (f floorsQueue) Swap(i, j int) {
	f.floors[i], f.floors[j] = f.floors[j], f.floors[i]
}

func (f *floorsQueue) Push(x interface{}) {
	item := x.(int)
	if f.containsFloor(item) {
		return
	}
	f.floors = append(f.floors, item)
}

func (f *floorsQueue) Pop() interface{} {
	old := f.floors
	n := len(old)
	x := old[n-1]
	f.floors = old[0 : n-1]
	return x
}
