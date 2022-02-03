package result

import "fmt"

type Result[OkT,ErrT any] interface {
    Unwrap() (OkT, error)
    UnwrapError() (ErrT, error)
}


type Ok[OkT, ErrT any] struct {
    val OkT
}

func (ok Ok[OkT, ErrT]) Unwrap() (OkT, error) {
    return ok.val, nil
}

func (ok Ok[OkT, ErrT]) UnwrapError() (ErrT, error) {
    var err ErrT
    return err, fmt.Errorf("cannot unwrap error, not an error: %v", ok)
}

func ToOk[OkT, ErrT any](v OkT) Ok[OkT, ErrT] {
    return Ok[OkT, ErrT]{v}
}



type Err[OkT,ErrT any] struct {
    val ErrT
}

func (err Err[OkT, ErrT]) Unwrap() (OkT, error) {
    var ok OkT
    return ok, fmt.Errorf("cannot unwrap ok, not an ok: %v", err)
}

func (err Err[OkT, ErrT]) UnwrapError() (ErrT, error) {
    return err.val, nil
}

func ToErr[OkT, ErrT any](e ErrT) Err[OkT, ErrT] {
    return Err[OkT, ErrT]{e}
}



func Fmap[OkT1,OkT2,ErrT any](f func(OkT1) OkT2, res Result[OkT1,ErrT]) Result[OkT2,ErrT] {
    if v, err := res.Unwrap(); err == nil {
        return Ok[OkT2, ErrT]{f(v)}
    }

    e, _ := res.UnwrapError()

    return ToErr[OkT2, ErrT](e)
}


func AndThen[OkT1, OkT2, ErrT any](f func(OkT1) Result[OkT2, ErrT], res Result[OkT1, ErrT]) Result[OkT2, ErrT] {
    if v, err := res.Unwrap(); err == nil {
        return f(v)
    }

    e, _ := res.UnwrapError()

    return ToErr[OkT2, ErrT](e)
}

func Pure[OkT, ErrT any](v OkT) Result[OkT, ErrT] {
    return ToOk[OkT, ErrT](v)
}

func Join[OkT, ErrT any](res Result[Result[OkT, ErrT], ErrT]) Result[OkT, ErrT] {
    if r, err := res.Unwrap(); err == nil {
        return r
    }

    e, _ := res.UnwrapError()

    return ToErr[OkT, ErrT](e)
}

