package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
    "bufio"
    "os"
)

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer inputFile.Close()

    var columnOne []int
    var columnTwo []int
    var columnThree []int

    scanner := bufio.NewScanner(inputFile) 
    for scanner.Scan() {
        line := strings.Fields(strings.TrimSpace(scanner.Text()))
        a, err := strconv.Atoi(line[0])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        b, err := strconv.Atoi(line[1])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        c, err := strconv.Atoi(line[2])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        columnOne = append(columnOne, a)
        columnTwo = append(columnTwo, b)
        columnThree = append(columnThree, c)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
    var possibleTriangles []int
    possibleTriangles = append(possibleTriangles, columnOne...)
    possibleTriangles = append(possibleTriangles, columnTwo...)
    possibleTriangles = append(possibleTriangles, columnThree...)
    actualTriangles := 0
    for i := 0; i < len(possibleTriangles); i += 3 {

        a := possibleTriangles[i]
        b := possibleTriangles[i + 1]
        c := possibleTriangles[i + 2]

        if a + b > c && b + c > a && a + c > b {
            actualTriangles += 1
        }
    }
    
    fmt.Println(actualTriangles)

}