package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"
)

type Point2d struct {
	x int
	y int
}

const numberOfWorkers = 8

var r = regexp.MustCompile(`\((\d*),(\d*)\)`)

func findArea(input <-chan string, resultChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range input {
		var points []Point2d
		for _, match := range r.FindAllStringSubmatch(line, -1) {
			x, y := match[1], match[2]
			xi, _ := strconv.Atoi(x)
			yi, _ := strconv.Atoi(y)
			points = append(points, Point2d{xi, yi})
		}

		area := 0.0
		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y - b.x*a.y)
		}
		resultChan <- math.Abs(area) / 2.0
	}
}

func splitToPoint(inputCh chan string) error {
	absPath, err := filepath.Abs("./threadpool/polygons.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer close(inputCh)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)  //increase buffer size to 1MB
	for scanner.Scan() {
		inputCh <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {

	inputChan := make(chan string, 1000)
	resultChan := make(chan float64)
	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go findArea(inputChan, resultChan, &wg)
	}

	go func() {
		err := splitToPoint(inputChan)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for result := range resultChan {
		fmt.Println(result)
	}
	timeTaken := time.Since(start)
	fmt.Printf("Time taken: %s\n", timeTaken)
}
