package utils

import (
	"fmt"
	"reflect"
)

type Result struct {
	Err  error
	Data interface{}
}

func (r Result) HasError() bool { return r.Err == nil }

func (r Result) HasData() bool { return r.Data == nil }

func (r Result) Throw() { ThrowResult(r) }

type ErrorWrapped struct {
	error
	AdditionalMessage string
}

func (e ErrorWrapped) Error() string {
	return fmt.Sprintf("%s: %v", e.AdditionalMessage, e.error)
}

var EmptyResult = Result{}

var InvalidResult = Result{Err: fmt.Errorf("invalid result ")}

func ThrowError(err error) { panic(Result{Err: err}) }

func ThrowIfError(err error) {
	if err != nil {
		ThrowError(err)
	}
}

func ThrowData(data interface{}) {
	panic(Result{Data: data})
}

func ThrowEmptyResult() { ThrowResult(EmptyResult) }

func ThrowResult(res Result) {
	panic(res)
}

func ResultFromJob(job func()) Result {
	chRes := make(chan Result)
	go func() {
		defer func() {
			if res := recover(); res == nil {
				chRes <- EmptyResult
			} else if result, ok := res.(Result); !ok {
				chRes <- InvalidResult
			} else {
				chRes <- result
			}
		}()
		job()
	}()
	return <-chRes
}

func UnwrapResultFromJob(job func()) (interface{}, error) {
	return UnwrapResult(ResultFromJob(job))
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

func WrapError(msg string, ori error) error {
	return ErrorWrapped{ori, msg}
}

func WrapAndThrowError(msg string, ori error) {
	ThrowError(WrapError(msg, ori))
}
func WrapAndThrowIfError(msg string, ori error) {
	if ori != nil {
		ThrowError(WrapError(msg, ori))
	}
}
