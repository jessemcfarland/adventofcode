package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
)

type CharacterCount struct {
    character string
    count int
}

func getUniqueCharacters(s string) []string {
    m := make(map[rune]uint, len(s))
    unique := ""
    for _, r := range s {
        m[r]++
        if m[r] == 1 {
            unique = unique + string(r)
        }
    }
    return strings.Split(unique, "")
}

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer inputFile.Close()

    var rows []string
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        rows = append(rows, strings.TrimSpace(scanner.Text()))
    }

    streams := make(map[int]string)
    for _, row := range rows {
        columns := strings.Split(row, "")
        for column, value := range columns {
            streams[column] = streams[column] + value
        } 
    }

    message := make([]string, len(streams))
    for stream := 0; stream < len(streams); stream++ {
        uniqueCharacters := getUniqueCharacters(streams[stream])
        var characterCount CharacterCount
        for _, character := range uniqueCharacters {
            count := strings.Count(streams[stream], character)
            if count > characterCount.count {
                characterCount.character = character
                characterCount.count = count
                message[stream] = characterCount.character 
            }
        }
    }

    fmt.Println(strings.Join(message, ""))
}