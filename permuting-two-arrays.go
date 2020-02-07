package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the twoArrays function below.
func twoArrays(k int32, A []int32, B []int32) string {

    for i, j := 0, 0; i < len(A) || j < len(B); i, j = i + 1, j + 1 {
        if i > 0 && i < len(A) {
            if A[i] < A[i-1] {
                A[i], A[i-1] = A[i-1], A[i]
                i -= 2
            }
        }
        if j > 0 && j < len(B) {
            if B[j] > B[j-1] {
                B[j], B[j-1] = B[j-1], B[j]
                j -= 2
            }
        }
    }

    for i := 0; i < (len(A) + len(B)) / 2; i++ {
        if A[i] + B[i] < k {
            return "NO"
        }
    }

    return "YES"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        nk := strings.Split(readLine(reader), " ")

        nTemp, err := strconv.ParseInt(nk[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        kTemp, err := strconv.ParseInt(nk[1], 10, 64)
        checkError(err)
        k := int32(kTemp)

        ATemp := strings.Split(readLine(reader), " ")

        var A []int32

        for i := 0; i < int(n); i++ {
            AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
            checkError(err)
            AItem := int32(AItemTemp)
            A = append(A, AItem)
        }

        BTemp := strings.Split(readLine(reader), " ")

        var B []int32

        for i := 0; i < int(n); i++ {
            BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
            checkError(err)
            BItem := int32(BItemTemp)
            B = append(B, BItem)
        }

        result := twoArrays(k, A, B)

        fmt.Fprintf(writer, "%s\n", result)
    }

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
