package bt

import (
	"errors"
)

func FallbackFunc[T any](s ...func() (T, error)) (T, error) {
	var errlog error
	for _, f := range s {
		v, err := f()
		if err != nil {
			errors.Join(errlog, err)
			continue
		}
		return v, nil

	}
	var temp T
	return temp, errlog
}
func FallbackArg[Tin any, Tout any](f func(Tin) (Tout, error), in ...Tin) (Tout, error) {
	var errlog error
	for _, i := range in {
		v, err := f(i)
		if err != nil {
			errors.Join(errlog, err)
			continue
		}
		return v, nil

	}
	var temp Tout
	return temp, errlog
}

func SequenceNode[TOut any](f func() (TOut, error), preerr error) (TOut, error) {
	var temp TOut
	if preerr != nil {

		return temp, preerr
	}
	temp, err := f()

	return temp, err
}
func FallbackNode[T any](prerr error, f func() (T, error)) (T, error) {
	var temp T
	if prerr != nil {
		res, err := f()
		return res, errors.Join(prerr, err)
	}

	return temp, prerr
}
