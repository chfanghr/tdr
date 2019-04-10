package utils

import "fmt"

type Result struct {
	Err  error
	Data interface{}
}

var InvalidResult = Result{Err: fmt.Errorf("invalid result ")}

func (r Result) IsError() bool { return r.Err != nil }

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

func UnwrapResult(job func()) (interface{}, error) {
	chRes := make(chan Result)
	func() {
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
	return res.Data, res.Err
}
