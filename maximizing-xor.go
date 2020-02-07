package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the maximizingXor function below.
func maximizingXor(l int, r int) int {
    var maks int = l ^ l

    for i := l; i <= r; i++ {
        for j := l; j <= r; j++ {
            if i ^ j > maks {
                maks = i ^ j
            }
        }
    }
    return maks
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    lTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    l := int(lTemp)

    rTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    r := int(rTemp)

    result := maximizingXor(l, r)

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

