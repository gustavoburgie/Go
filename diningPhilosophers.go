package main

import (
    "fmt"
    "time"
)

const timesToEat = 3

// Philosopher represents a philosopher.
type Philosopher struct {
    id              int
    leftFork, rightFork chan bool
}

// run simulates the philosopher thinking and eating.
func (p *Philosopher) run(done chan<- bool) {
    for i := 0; i < timesToEat; i++ {
        fmt.Printf("Philosopher %d is thinking.\n", p.id)
        time.Sleep(time.Second)

        <-p.leftFork // pick up left fork
        <-p.rightFork // pick up right fork

        fmt.Printf("Philosopher %d is eating.\n", p.id)
        time.Sleep(time.Second)

        p.leftFork <- true // put down left fork
        p.rightFork <- true // put down right fork
    }
    done <- true // Signal that this philosopher is done eating
}

func main() {
    const numPhilosophers = 5

    startTime := time.Now() // Capture start time

    forks := make([]chan bool, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        forks[i] = make(chan bool, 1)
        forks[i] <- true // put a fork on the table
    }

    done := make(chan bool, numPhilosophers) // Channel to signal completion
    philosophers := make([]Philosopher, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        philosophers[i] = Philosopher{
            id: i + 1,
            leftFork: forks[i],
            rightFork: forks[(i+1)%numPhilosophers],
        }
        go philosophers[i].run(done)
    }

    // Wait for all philosophers to signal they are done
    for i := 0; i < numPhilosophers; i++ {
        <-done
    }

    executionTime := time.Since(startTime) // Calculate total execution time
    fmt.Printf("Total execution time: %s\n", executionTime)
}
