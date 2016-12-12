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

    password := make(map[int]string)
    for _, id := range doorIDs {
        index := 0
        for len(password) < 8 {
            data := []byte(id + strconv.Itoa(index))
            re := regexp.MustCompile("^00000")
            md5sum := md5.Sum(data)
            md5sumString := hex.EncodeToString(md5sum[:16])
            if re.MatchString(md5sumString) {
                n, err := strconv.Atoi(md5sumString[5:6])
                if err == nil {
                    if n >= 0 && n <= 7 {
                        if _, present := password[n]; !present {
                            password[n] = md5sumString[6:7]
                        }
                    }
                }
            }
            index++
        }
    }

    passwordString := ""
    for p := 0; p <= 7; p++ {
        passwordString = passwordString + password[p]
    }
    
    fmt.Println(passwordString)
}