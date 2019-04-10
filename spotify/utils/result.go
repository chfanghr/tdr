package utils

import (
	"fmt"
	"reflect"
)

type Result struct {
	Err  error
	Data interface{}
}

var InvalidResult = Result{Err: fmt.Errorf("invalid result ")}

func ThrowIfError(err error) {
	if err != nil {
		panic(Result{Err: err})
	}
}

func ThrowData(data interface{}) {
	panic(Result{Data: data})
}

func ThrowResult(res Result) {
	panic(res)
}

func UnwrapResultFromJob(job func()) (interface{}, error) {
	chRes := make(chan Result)
	go func() {
		defer func() {
			if res := recover(); res == nil {
				chRes <- InvalidResult
			} else if result, ok := res.(Result); !ok {
				chRes <- InvalidResult
			} else {
				chRes <- result
			}
		}()
		job()
	}()
	res := <-chRes
	return UnwrapResult(res)
}

func UnwrapResult(res Result) (interface{}, error) {
	return res.Data, res.Err
}

func AsyncDo(job func()) <-chan Result {
	chRes := make(chan Result)
	go func() {
		data, err := UnwrapResultFromJob(job)
		chRes <- Result{Err: err, Data: data}
	}()
	return chRes
}

func IsSameError(a, b error) bool { return reflect.DeepEqual(a, b) }
