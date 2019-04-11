package utils

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type anotherError struct {
	msg string
}

func (a anotherError) Error() string {
	return a.msg
}

func TestIsSameError(t *testing.T) {
	type args struct {
		a error
		b error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"different error message and same type", args{fmt.Errorf("a"), fmt.Errorf("b")}, false},
		{"same error message and same type", args{fmt.Errorf("a"), fmt.Errorf("a")}, true},
		{"same error message and different type", args{fmt.Errorf("a"), anotherError{"a"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameError(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsSameError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrapResult(t *testing.T) {
	type args struct {
		job func()
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"throw data", args{func() {
			ThrowIfError(nil)
			ThrowData("1")
		}}, "1", false},
		{"throw error", args{func() {
			ThrowIfError(errors.New("error"))
			ThrowData("1")
		}}, nil, true},
		{"throw result", args{func() {
			ThrowResult(Result{Data: 1, Err: errors.New("error")})
		}}, 1, true},
		{"throw nothing", args{func() {
		}}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnwrapResultFromJob(tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnwrapResultFromJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapResultFromJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
