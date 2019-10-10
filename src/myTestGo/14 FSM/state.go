package main

import (
	"reflect"
)

type State interface {
	Name() string
	EnableSameTransit() bool
	OnBegin()
	OnEnd()
	CanTransitTo(name string) bool
}

func StateName(s State) string {
	if s == nil {
		return "none"
	}

	// 使用反射获取状态的名称
	return reflect.TypeOf(s).Elem().Name()
}
