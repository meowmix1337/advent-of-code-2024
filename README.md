# Acuity Go Advent of Code

## How to Run
Run the following make command:
1. `make run ${DAY}` where ${DAY} is just the day number. i.e. 1, 2, 3, 4, 10, 12, etc
   1. i.e. `make run 3` will run day 3 of advent

## Create a new day
Run the following command to create a new day solver:
1. `make new-day ${DAY}` where ${DAY} is just the day number. i.e. 1, 2, 3, 4, 10, 12, etc
   1. i.e. `make new-day 3` will generate `day03` under `./solutions`
2. Update `./solutions/day_factory/day_factory.go` to include the new day solver you created.
3. Start solving!

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