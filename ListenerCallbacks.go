package gumjs

import "sync"

type IListenerCallback interface {
	MustGetVal(key string) any
	GetVal(key string) (any, bool)
	SetVal(key string, val any)
	HasVal(key string) bool
	GetOnEnter() func(cb IListenerCallback, context *InvocationContext)
	GetOnLeave() func(cb IListenerCallback, context *InvocationContext)
}

type InvocationListenerCallbacks struct {
	val     sync.Map
	OnEnter func(cb IListenerCallback, context *InvocationContext)
	OnLeave func(cb IListenerCallback, context *InvocationContext)
}

func (i *InvocationListenerCallbacks) HasVal(key string) bool {
	_, b := i.val.Load(key)
	return b
}
func (i *InvocationListenerCallbacks) MustGetVal(key string) any {
	v, _ := i.val.Load(key)
	return v
}
func (i *InvocationListenerCallbacks) GetVal(key string) (any, bool) {
	v, b := i.val.Load(key)
	return v, b
}
func (i *InvocationListenerCallbacks) SetVal(key string, val any) {
	i.val.Store(key, val)
}
func (i *InvocationListenerCallbacks) GetOnEnter() func(cb IListenerCallback, context *InvocationContext) {
	if i.OnEnter != nil {
		return i.OnEnter
	}
	return nil
}
func (i *InvocationListenerCallbacks) GetOnLeave() func(cb IListenerCallback, context *InvocationContext) {
	if i.OnLeave != nil {
		return i.OnLeave
	}
	return nil
}
