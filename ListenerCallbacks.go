package gumjs

type IListenerCallback interface {
	GetOnEnter() func(context *InvocationContext)
	GetOnLeave() func(context *InvocationContext)
}

type InvocationListenerCallbacks struct {
	OnEnter func(context *InvocationContext)
	OnLeave func(context *InvocationContext)
}

func (i *InvocationListenerCallbacks) GetOnEnter() func(context *InvocationContext) {
	if i.OnEnter != nil {
		return i.OnEnter
	}
	return nil
}
func (i *InvocationListenerCallbacks) GetOnLeave() func(context *InvocationContext) {
	if i.OnLeave != nil {
		return i.OnLeave
	}
	return nil
}

type InstructionProbeCallback func(context *InvocationContext)

func (i InstructionProbeCallback) GetOnEnter() func(context *InvocationContext) {
	return i
}

func (i InstructionProbeCallback) GetOnLeave() func(context *InvocationContext) {
	return nil
}
