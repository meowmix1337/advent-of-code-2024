# Acuity Go Advent of Code

## How to Run
```
go run main.go --day=X

example:
go run main.go --day=1
```

## Structure
* **cmd**
  * Cobra command line creation
  * You shouldn't need to modify this file
* **input**
  * Your advent of code input
  * Follows dayXX.txt naming format. i.e. `day01.txt`, `day02.txt`, `day12.txt`, etc
* **internal**
  * Business logic for parser and running your day solutions
  * You shouldn't need to modify this code
* **solutions**
  * This is where you will implement the day's solutions
  * Must implement Part1() and Part2() functions
  * `day_factory`
    * Add additional days solvers here
* **util**
  * Helper functions that may potentially help with parsing
  * Add more functions here to be used across your solutions