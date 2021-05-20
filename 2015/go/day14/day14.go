package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	//readAndPrepare("/home/slash3b/Projects/aoc/2015/go/day14/input")
	readAndPrepareSecondPart("/home/slash3b/Projects/aoc/2015/go/day14/input")
}

func readAndPrepare(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	//maxDistance := 0

	for _, v := range inputs {
		reindeerDistance := 0
		line := strings.Split(v, " ")
		name := line[0]
		speed, _ := strconv.Atoi(line[3])
		runtime, _ := strconv.Atoi(line[6])
		resttime, _ := strconv.Atoi(line[13])

		fmt.Println("---------------------")
		fmt.Printf("%s speed=%d runtime=%d rest=%d \n", name, speed, runtime, resttime)

		hop := runtime + resttime

		time := 2503
		//time := 1000

		hops := time / hop
		partialHopTime := time % hop

		fmt.Println("did ", hops, " hops")
		fmt.Println("partial time is  ", partialHopTime)

		reindeerDistance = hops * runtime * speed

		if partialHopTime > runtime {
			reindeerDistance += runtime * speed
		} else {
			reindeerDistance += partialHopTime * speed
		}

		fmt.Printf("ANSWER: %s has run for %d \n", name, reindeerDistance)

	}

	// after 2503 seconds
}

type conf struct {
	speed         int
	runTimeLimit  int
	restTimeLimit int
}

type deer struct {
	points      int
	name        string
	isRunning   bool
	runCounter  int
	restCounter int
	mileage     int
	c           conf
}

func (d *deer) tick() {
	if d.isRunning {
		d.mileage += d.c.speed
		d.runCounter++
		if d.runCounter == d.c.runTimeLimit {
			d.isRunning = false
			d.runCounter = 0
		}
	} else {
		d.restCounter++
		if d.restCounter == d.c.restTimeLimit {
			d.isRunning = true
			d.restCounter = 0
		}
	}
}

func readAndPrepareSecondPart(s string) {
	res, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(string(res), "\n")

	deers := make(map[string]*deer)

	for _, v := range inputs {
		line := strings.Split(v, " ")
		nameValue := line[0]
		speedValue, _ := strconv.Atoi(line[3])
		runTimeValue, _ := strconv.Atoi(line[6])
		restTimeValue, _ := strconv.Atoi(line[13])

		deers[nameValue] = &deer{
			name:      nameValue,
			isRunning: true,
			c: conf{
				speed:         speedValue,
				runTimeLimit:  runTimeValue,
				restTimeLimit: restTimeValue,
			}}
	}

	for _, v := range deers {
		fmt.Println(*v)
	}

	seconds := 2503
	//seconds := 1000

	for i := 0; i < seconds; i++ {
		score := make(map[int][]string)

		for _, d := range deers {
			d.tick()
			_, exists := score[d.mileage]
			if !exists {
				score[d.mileage] = []string{d.name}
				continue
			}
			score[d.mileage] = append(score[d.mileage], d.name)
		}

		// now find maximum mileage at that point
		currentMax := 0
		for k := range score {
			if k > currentMax {
				currentMax = k
			}
		}

		for _, v := range score[currentMax] {
			deers[v].points++
		}
	}

	maxMileage := 0
	maxPoints := 0
	for _, d := range deers {
		if d.mileage > maxMileage {
			maxMileage = d.mileage
		}

		if d.points > maxPoints {
			maxPoints = d.points
		}
	}

	fmt.Println("^^^^ mileage", maxMileage)
	fmt.Println("^^^^ points", maxPoints)
	// after 2503 seconds
}
