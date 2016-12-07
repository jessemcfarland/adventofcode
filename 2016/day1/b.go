package main

import (
    "fmt"
    "math"
    "strings"
    "strconv"
    "log"
    "encoding/csv"
    "bufio"
    "os"
)

type coordinate struct {
    X, Y float64
}

func process_instruction(input []string) (string, float64) {
    integer, err := strconv.Atoi(input[1])
    if err != nil {
        log.Fatal(err)
    }
    return input[0], float64(integer)
}

func (position coordinate) move(facing string, distance float64) coordinate {
    var movement coordinate 
    if facing == "North" {
        movement = coordinate{0, distance}
    }
    if facing == "South" {
        movement = coordinate{0, distance * -1}
    }
    if facing == "East" {
        movement = coordinate{distance, 0}
    }
    if facing == "West" {
        movement = coordinate{distance * -1, 0}
    }
    return coordinate{position.X + movement.X, position.Y + movement.Y}
}

func turn(facing, turn string) string {
    var nowFacing string
    if (facing == "North" && turn == "L") || (facing == "South" && turn == "R") {
        nowFacing = "West"
    }
    if (facing == "North" && turn == "R") || (facing == "South" && turn == "L") {
        nowFacing = "East"
    }
    if (facing == "East" && turn == "L") || (facing == "West" && turn == "R") {
        nowFacing = "North"
    }
    if (facing == "East" && turn == "R") || (facing == "West" && turn == "L") {
        nowFacing = "South"
    }
    return nowFacing
}

func distance(s, e coordinate) float64 {
    return math.Abs(e.X - s.X) + math.Abs(e.Y - s.X)
}

func main() {
    initialPosition := coordinate{0, 0}
    startFacing := "North"
    inputFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    csvInput := csv.NewReader(bufio.NewReader(inputFile))
    input, err := csvInput.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    instructions := input[0]
    position := initialPosition
    facing := startFacing
    path := map[coordinate]bool{}
    Walk:
    for _, value := range instructions {
        instruction := strings.SplitN(strings.TrimSpace(value), "", 2)
        turnDirection, blocksToWalk := process_instruction(instruction)
        nowFacing := turn(facing, turnDirection)
        facing = nowFacing
        for block := float64(1); block <= blocksToWalk; block++ {
            newPosition := position.move(facing, float64(1))
            position = newPosition
            if _, present := path[position]; present {
                break Walk
            }
            path[position] = true
        }
    }
    distanceToHQ := distance(initialPosition, position)
    fmt.Println(fmt.Sprintf("Distance to HQ: %f", distanceToHQ))

}
