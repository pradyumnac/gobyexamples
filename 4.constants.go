package main

import (
    "math"
    "fmt"
)

const s string = "Pradyumna"

func main(){
    fmt.Println(s)

    const n = 50000
    fmt.Println(n)

    const m = 5e20/n
    fmt.Println(m)
    fmt.Println(int64(m))

    fmt.Println(math.Sin(m))
}
