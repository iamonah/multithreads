package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	matrixSize = 256
)

var (
	matrixA   = [matrixSize][matrixSize]int{}
	matrixB   = [matrixSize][matrixSize]int{}
	result    = [matrixSize][matrixSize]int{}
	rwLock    = sync.RWMutex{}
	// cond coordinates worker goroutines with the main goroutine.
	// Each worker calls Wait to sleep until main broadcasts that fresh
	// matrix data is ready, and Wait temporarily releases the read lock
	// while the worker is blocked so the writer can proceed.
	cond      = sync.NewCond(rwLock.RLocker())
	waitGroup = sync.WaitGroup{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func WorkoutRow(row int) {
	// Wait must be called while holding the condition variable's locker.
	// The call atomically unlocks, blocks until Broadcast wakes the worker,
	// then re-locks before returning so the row computation sees a consistent
	// snapshot of matrixA and matrixB for that round.
	rwLock.RLock() //acuiring the reader portion of the lock
	for {
		waitGroup.Done()
		cond.Wait()
		for column := 0; column < matrixSize; column++ {
			for i := 0; i < matrixSize; i++ {
				result[row][column] += matrixA[row][i] * matrixB[i][column]
			}
		}
	}
}

func main() {
	fmt.Println("Working...")
	waitGroup.Add(matrixSize)
	for i := 0; i < matrixSize; i++ {
		go WorkoutRow(i)
	}

	//multiplication of 100 pairs of matrices
	start := time.Now()
	for i := 0; i < 100; i++ {
		waitGroup.Wait()
		rwLock.Lock()
		go generateRandomMatrix(&matrixA)
		go generateRandomMatrix(&matrixB)
		waitGroup.Add(matrixSize)
		rwLock.Unlock()
		cond.Broadcast()
	}
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
