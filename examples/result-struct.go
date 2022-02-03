package main

import (
    "fmt"
    "strconv"

    "github.com/alex-bechanko/go-monad-examples/pkg/struct/result"
)

func stringToInt(s string) result.Result[int64, error] {
    n, err := strconv.ParseInt(s, 10, 64)

    if err != nil {
        return result.ToErr[int64, error](err)
    }

    return result.ToOk[int64, error](n)
}


func double(n int64) int64 {
    return n * 2
}


func inverse(n int64) result.Result[float64, error] {
    if n == 0 {
        return result.ToErr[float64, error](fmt.Errorf("unable to divide by %d", n))
    }

    return result.ToOk[float64, error](1/float64(n))
}

func main() {

    res := result.AndThen(
        inverse,
        result.Fmap(double, stringToInt("32")),
    )

    fmt.Println(res)

}
