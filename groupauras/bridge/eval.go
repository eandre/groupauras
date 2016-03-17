package bridge

import "errors"

func EvalEventBool(src string) (func(event string, args []interface{}) bool, error) {
	return nil, errors.New("NYI")
}

func EvalEvent(src string) (func(event string, args []interface{}), error) {
	return nil, errors.New("NYI")
}

func EvalUpdate(src string) (func(dt float32), error) {
	return nil, errors.New("NYI")
}
