package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "regexp"
    "strings"
    "strconv"
    "sort"
)

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

func getKeys(m map[int]string) []int {
    keys := make([]int, len(m))
    i := 0
    for k := range m {
        keys[i] = k
        i++
    }
    return keys
}

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer inputFile.Close()

    var possibleRooms []string
    sumSectorIDs := 0
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        possibleRooms = append(possibleRooms, scanner.Text())
    }

    re := regexp.MustCompile("(([a-z]*-)*)([0-9]*)\\[([a-z]*)\\]")
    for _, encryptedRoom := range possibleRooms {
        fields := re.FindAllStringSubmatch(encryptedRoom, -1)[0]
        encryptedName := strings.Replace(fields[1], "-", "", -1)
        sectorID, err := strconv.Atoi(fields[3])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        checksum := fields[4]

        characterCount := make(map[int]string)
        var uniqueCharacters sort.StringSlice
        uniqueCharacters = getUniqueCharacters(encryptedName)
        uniqueCharacters.Sort()
        for _, character := range uniqueCharacters {
            count := strings.Count(encryptedName, character)
            if _, present := characterCount[count]; present {
                characterCount[count] = characterCount[count] + character
            } else {
                characterCount[count] = character
            }
        }

        possibleChecksum := ""
        keys := getKeys(characterCount)
        sort.Sort(sort.Reverse(sort.IntSlice(keys)))
        for _, key := range keys {
            possibleChecksum = possibleChecksum + characterCount[key]
        }

        if possibleChecksum[:5] == checksum {
            sumSectorIDs += sectorID
        }
    }
    fmt.Println(sumSectorIDs)
}