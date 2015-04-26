# elevator-control
Simple elevator control system

## Setup and run

1. Download repo with git client or with go tools ```go get github.com/mmigacz/elevator-control```
2. Build application with ```make build```
3. Run application 
```./elevator-control [number_of_lifts]```. The application starts by default with 2 elvators.
Alternatively you can run app directly with go with the following command ```go run main.go [number_of_lifts]```


After start the application will prompt you a command line. Type anything + Enter to get list of available commands.

* *status* - displays list of elevators

  ```
  #id - elevator's id
  #floor - current flor
  #direction - stop|up|down - current state of the elevator
  #upq - lenght of the queue with stops above
  #downq - lenght of the queue with stops below
  id:[0], floor:[0], direction:[stop], upq:[0], downq:[0]
  id:[1], floor:[0], direction:[stop], upq:[0], downq:[0]
  ```
* *step* - next step of the simulation
* *update elevator_id floor_id* - add destination floor_id to the elevator_id. Reflects pressing a button in the elevator.
* *pickup floor_id direction* - pickup any lift on floor_id in to the direction (-1 down, +1 up). Reflects pressing a button on the corridor
* *exit* - exits application

## How it works

