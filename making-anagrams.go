package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    // "math"
    "strings"
)

// Complete the makingAnagrams function below.
func makingAnagrams(s1 string, s2 string) int {
    var res [26]int
    var n int

    for _, val := range s1 {
        res[val%26]++
    }
    for _, val := range s2 {
        res[val%26]--
    }
    for _, val := range res {
        if val != 0 {
            if val < 0 {
                n += (val * -1)
            } else {
                n += val
            }
        }
    }

    return n

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s1 := readLine(reader)

    s2 := readLine(reader)

    result := makingAnagrams(s1, s2)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

