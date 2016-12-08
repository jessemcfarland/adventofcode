package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
    "bufio"
    "os"
)

func nextButton(button int, instruction string) int {
    var newButton int
    switch instruction {
    case "U":
        switch button {
        case 1, 2, 3:
            newButton = button 
        default:
            newButton = button - 3
        }
    case "D":
        switch button {
        case 7, 8, 9:
            newButton = button 
        default:
            newButton = button + 3
        }
    case "L":
        switch button {
        case 1, 4, 7:
            newButton = button 
        default:
            newButton = button - 1
        }
    case "R":
        switch button {
        case 3, 6, 9:
            newButton = button 
        default:
            newButton = button + 1
        }
    }
    return newButton
}

func main() {
    button := 5
    inputFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    var instructions [][]string
    var code []string

    scanner := bufio.NewScanner(inputFile) 
    for scanner.Scan() {
        instructionSet := strings.Split(strings.TrimSpace(scanner.Text()), "")
        instructions = append(instructions, instructionSet)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for _, instructionSet := range instructions {
        for _, instruction := range instructionSet {
            button = nextButton(button, instruction)
        }
        code = append(code, strconv.Itoa(button))
    }
    fmt.Println(strings.Join(code, ""))
}