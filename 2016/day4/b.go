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
    "math"
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

func decrypt(cipherText string, sectorID int) string {
    letterMap := map[string]int{
        "a": 0,
        "b": 1,
        "c": 2,
        "d": 3,
        "e": 4,
        "f": 5,
        "g": 6,
        "h": 7,
        "i": 8,
        "j": 9,
        "k": 10,
        "l": 11,
        "m": 12,
        "n": 13,
        "o": 14,
        "p": 15,
        "q": 16,
        "r": 17,
        "s": 18,
        "t": 19,
        "u": 20,
        "v": 21,
        "w": 22,
        "x": 23,
        "y": 24,
        "z": 25,
    }

    floatMap := make(map[float64]string)
    for key, value := range letterMap {
        floatMap[float64(value)] = key
    }

    plainText := ""
    for _, letter := range strings.Split(cipherText, "") {
        if letter == "-" {
            plainText = plainText + " "
        } else {
            x := float64(letterMap[letter])
            n := float64(sectorID)
            m := float64(26)
            f := math.Abs(math.Mod(x + n, m))
            plainText = plainText + floatMap[f]
        }
    }

    return plainText
}

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer inputFile.Close()

    var possibleRooms []string
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        possibleRooms = append(possibleRooms, scanner.Text())
    }

    re := regexp.MustCompile("(([a-z]*-)*)([0-9]*)\\[([a-z]*)\\]")
    for _, encryptedRoom := range possibleRooms {
        fields := re.FindAllStringSubmatch(encryptedRoom, -1)[0]
        cipherText := fields[1]
        strippedCipherText := strings.Replace(cipherText, "-", "", -1)
        sectorID, err := strconv.Atoi(fields[3])
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        checksum := fields[4]

        characterCount := make(map[int]string)
        var uniqueCharacters sort.StringSlice
        uniqueCharacters = getUniqueCharacters(strippedCipherText)
        uniqueCharacters.Sort()
        for _, character := range uniqueCharacters {
            count := strings.Count(strippedCipherText, character)
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
            plainText := decrypt(cipherText, sectorID)
            fmt.Println(plainText, sectorID)
        }
    }
}