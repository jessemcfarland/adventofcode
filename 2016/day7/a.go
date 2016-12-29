package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "regexp"
)

func MatchAbbaSequence(s string) bool {
    for i := 0; i < len(s) - 1; i++ {
        first := string(s[i])
        second := string(s[i+1])
        if first != second {
            re := regexp.MustCompile(first + second + second + first)
            if re.MatchString(s) {
                return true
            }
        }
    } 
    return false
}

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer inputFile.Close()

    var rawIPAddresses []string
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        rawIPAddresses = append(rawIPAddresses, strings.TrimSpace(scanner.Text()))
    }


    var processedIPAddresses []string
    re := regexp.MustCompile("\\[[a-z]*\\]")
    for _, ip := range rawIPAddresses {
        hypernetSequences := re.FindAllStringSubmatch(ip, -1)
        processedIP := ip
        for _, sequence := range hypernetSequences {
            processedSequence := sequence[0]
            processedSequence = strings.Replace(processedSequence, "[", "", -1)
            processedSequence = strings.Replace(processedSequence, "]", "", -1)
            //for i := 0; i < len(processedSequence) - 1; i++ {
            //    first := string(processedSequence[i])
            //    second := string(processedSequence[i+1])
            //    abba := regexp.MustCompile(first + second + second + first)
            //    if !abba.MatchString(processedSequence) {
            //        processedIP = strings.Replace(processedIP, sequence[0], "", -1) 
            //    } else {
            //        processedIP = ""
            //        break
            //    }
            if !MatchAbbaSequence(processedSequence) {
                processedIP = strings.Replace(processedIP, sequence[0], "", -1) 
            } else {
                fmt.Println(processedSequence)
                processedIP = ""
                break
            }
        }
        if processedIP != "" {
            processedIPAddresses = append(processedIPAddresses, processedIP)
        }
    }

    tls := 0
    for _, ip := range processedIPAddresses {
        if MatchAbbaSequence(ip) {
            fmt.Println(ip)
            tls++
        }

    }
    fmt.Println(tls)
}