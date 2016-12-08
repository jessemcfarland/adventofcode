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
        case 1, 2, 4, 5, 9:
            newButton = button 
        case 3, 13:
            newButton = button - 2
        default:
            newButton = button - 4
        }
    case "D":
        switch button {
        case 5, 9, 10, 12, 13:
            newButton = button 
        case 1, 11:
            newButton = button + 2
        default:
            newButton = button + 4
        }
    case "L":
        switch button {
        case 1, 2, 5, 10, 13:
            newButton = button 
        default:
            newButton = button - 1
        }
    case "R":
        switch button {
        case 1, 4, 9, 12, 13:
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
        if button < 1 || button > 9 {
            var hexButton string
            switch button {
            case 10:
                hexButton = "A"
            case 11:
                hexButton = "B"
            case 12:
                hexButton = "C"
            case 13:
                hexButton = "D"
            default:
                log.Fatal("Button is out of range")
            }
            code = append(code, hexButton)
        } else {
            code = append(code, strconv.Itoa(button))
        }
    }
    fmt.Println(strings.Join(code, ""))
}