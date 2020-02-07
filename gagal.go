package main
import "fmt"
func main() {
    var a[3]int
    var b[3]int
    var skorA, skorB int
    skorA = 0
    skorB = 0
    i := 0
    for i < 3 {
        fmt.Scan(&a[i])
        i++
    }
    for i<3 {
        fmt.Scan(&b[i])
        i++
    }
    for i < 3 {
        if a[i] >= b[i] {
            skorA = skorA + 1
        }else {
            skorB = skorB + 1
        }
        i++
    }
    fmt.Printf("%d %d\n", skorA, skorB)
}