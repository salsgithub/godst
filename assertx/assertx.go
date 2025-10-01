package assertx

import (
	"errors"
	"reflect"
	"testing"
)

func Equal[T any](t *testing.T, got, want T) {
	t.Helper()
	if !IsEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func NotEqual[T any](t *testing.T, got, want T) {
	t.Helper()
	if IsEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func True(t *testing.T, condition bool) {
	t.Helper()
	if !condition {
		t.Error("got: false, want: true")
	}
}

func False(t *testing.T, condition bool) {
	t.Helper()
	if condition {
		t.Error("got: true, want: false")
	}
}

func ErrorIs(t *testing.T, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func Nil(t *testing.T, v any) {
	t.Helper()
	if !isNil(v) {
		t.Errorf("got: %v, want: nil", v)
	}
}

func NotNil(t *testing.T, v any) {
	t.Helper()
	if isNil(v) {
		t.Error("got: nil, want: non-nil")
	}
}

func IsEqual[T any](a, b T) bool {
	if isNil(a) && isNil(b) {
		return true
	}
	if equatable, ok := any(b).(interface{ Equal(T) bool }); ok {
		return equatable.Equal(a)
	}
	return reflect.DeepEqual(a, b)
}

func isNil(v any) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Chan,
		reflect.Interface,
		reflect.Func,
		reflect.Map,
		reflect.Slice:
		return value.IsNil()
	}
	return false
}
