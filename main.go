package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mmigacz/elevator-control/elevator"
	"os"
	"strconv"
	"strings"
)

func main() {
	var elevators int
	flag.IntVar(&elevators, "elevators", 2, "Number of elevators")
	ec := elevator.NewElevatorController(elevators)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("cmd> ")

		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(strings.ToLower(text), "\n")

		switch {
		case text == "status":
			for _, e := range ec.Status() {
				fmt.Println(e.Str())
			}
		case text == "step":
			ec.Step()
			for _, e := range ec.Status() {
				fmt.Println(e.Str())
			}
		case strings.HasPrefix(text, "update"):
			updates := strings.Split(text, " ")
			if len(updates) == 3 {
				var err error
				var elevatorId, floorId int64

				if elevatorId, err = strconv.ParseInt(updates[1], 10, 32); err != nil {
					fmt.Println("Elevator id must be int, i.e. update 1 2")
					continue
				}
				if floorId, err = strconv.ParseInt(updates[2], 10, 32); err != nil {
					fmt.Println("Floor id must be int, i.e. update 1 2")
					continue
				}
				ec.Update(int(elevatorId), int(floorId))
			} else {
				fmt.Println("usage: update elevator_id floor_id")
			}

		case strings.HasPrefix(text, "pickup"):
			updates := strings.Split(text, " ")
			if len(updates) == 3 {
				var err error
				var floorId, direction int64

				if floorId, err = strconv.ParseInt(updates[1], 10, 32); err != nil {
					fmt.Println("Elevator id must be int, i.e. pickup 1 2")
					continue
				}
				if direction, err = strconv.ParseInt(updates[2], 10, 32); err != nil {
					fmt.Println("Floor id must be int, i.e. update 1 2")
					continue
				}
				ec.Pickup(int(floorId), int(direction))
			} else {
				fmt.Println("usage: pickup floor_id direction")
			}

		case text == "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Use one of the following: status, step, update, pickup, exit")
		}
	}

}
