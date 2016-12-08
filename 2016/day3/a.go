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

    var possibleTriangles [][]string
    actualTriangles := 0

    scanner := bufio.NewScanner(inputFile) 
    for scanner.Scan() {
        triangle := strings.Fields(strings.TrimSpace(scanner.Text()))
        possibleTriangles = append(possibleTriangles, triangle)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for _, triangle := range possibleTriangles {
        a, err := strconv.Atoi(triangle[0])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        b, err := strconv.Atoi(triangle[1])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        c, err := strconv.Atoi(triangle[2])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        if a + b > c && b + c > a && a + c > b {
            actualTriangles = actualTriangles + 1
        }
    }

    fmt.Println(actualTriangles)
}