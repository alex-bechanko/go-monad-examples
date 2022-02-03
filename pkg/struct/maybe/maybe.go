package maybe

import (
	"fmt"
)

// Maybe is an minimal interface definition for the Just and Nothing types to share
type Maybe[T any] interface {
	// Helper method to be used for the monadic functions and extracting values from the Maybe interface
	Unwrap() (T, error)
}

// Just is a struct that contains a value of type T
type Just[T any] struct {
	just T
}

func (j Just[T]) Unwrap() (T, error) {
	return j.just, nil
}



// Nothing is a struct that does not contain a value of type T
type Nothing[T any] struct{}

func (j Nothing[T]) Unwrap() (t T, err error) {
	err = fmt.Errorf("nothing!")
	return
}



// fmap :: (a -> b) -> m a -> m b
// Fmap is a function that applies a function to values within a Maybe interface.
// This functoin defines the plumbing to extract a value from Maybe, apply the function,
// and then wrap the value again in a Maybe
func Fmap[A, B any](f func(A) B, ma Maybe[A]) Maybe[B] {
	if u,err := ma.Unwrap(); err == nil {
		return Just[B]{f(u)}
	}

	return Nothing[B]{}
}

// andThen :: (a -> m b) -> m a -> m b
// AndThen is a function that lets you chain functions that return Maybe values.
// Another name for this commonly used is `bind` (in the Haskell world)
// This function defines the plumbing for extracting a value, and then applying the function.
// Since the function returns a Maybe, no further action is required.
func AndThen[A, B any](f func(a A) Maybe[B], ma Maybe[A]) Maybe[B] {
	if u,err := ma.Unwrap(); err == nil {
		return f(u)
	}

	return Nothing[B]{}
}

// pure :: a -> m a
// Pure is a simple constructor for making a Maybe interface value with Just.
// You could use typical struct construction for this too.
// This is defined here to emphasize that this is a necessary condition to making Maybe a monad.
// This name comes from the Haskell world, where your lifting a pure value
// (meaning no side-effects associated to the value) into your Monad/Functor type.
func Pure[A any](a A) Maybe[A] {
	return Just[A]{a}
}


// Join is a function used for collapsing structure down.
// 
func Join[A any](m Maybe[Maybe[A]]) Maybe[A] {
	if u, err := m.Unwrap(); err == nil {
		return u
	}

	return Nothing[A]{}
}
