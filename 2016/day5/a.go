package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "regexp"
    "strings"
    "strconv"
    "crypto/md5"
    "encoding/hex"
)

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer inputFile.Close()

    var doorIDs []string
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        doorIDs = append(doorIDs, strings.TrimSpace(scanner.Text()))
    }

    password := ""
    for _, id := range doorIDs {
        index := 0
        for len(password) < 8 {
            data := []byte(id + strconv.Itoa(index))
            re := regexp.MustCompile("^00000")
            md5sum := md5.Sum(data)
            md5sumString := hex.EncodeToString(md5sum[:16])
            if re.MatchString(md5sumString) {
                password = password + md5sumString[5:6]
            }
            index++
        }
    }
    fmt.Println(password)
}