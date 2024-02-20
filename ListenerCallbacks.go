package gumjs

import "sync"

type IListenerCallback interface {
	GetVal(key string) any
	SetVal(key string, val any)
	GetOnEnter() func(cb IListenerCallback, context *InvocationContext)
	GetOnLeave() func(cb IListenerCallback, context *InvocationContext)
}

type InvocationListenerCallbacks struct {
	val     sync.Map
	OnEnter func(cb IListenerCallback, context *InvocationContext)
	OnLeave func(cb IListenerCallback, context *InvocationContext)
}

func (i *InvocationListenerCallbacks) GetVal(key string) any {
	v, _ := i.val.Load(key)
	return v
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
