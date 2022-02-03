package main

import (
    "fmt"
    "strconv"

    "github.com/alex-bechanko/go-monad-examples/pkg/struct/maybe"
)

func stringToInt(s string) maybe.Maybe[int64] {
    n, err := strconv.ParseInt(s, 10, 64)

    if err != nil {
        return maybe.Nothing[int64]{}
    }

    return maybe.Pure(n)
}


func double(n int64) int64 {
    return n * 2
}


func inverse(n int64) maybe.Maybe[float64] {
    if n == 0 {
        return maybe.Nothing[float64]{}
    }

    return maybe.Pure(1/float64(n))
}

func main() {

    res := maybe.AndThen(
        inverse,
        maybe.Fmap(double, stringToInt("32")),
    )

    fmt.Println(res)

}
