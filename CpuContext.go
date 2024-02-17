package gumjs

type ICpuContextBase struct {
}

func (I *ICpuContextBase) GetRegValue(reg string) string {
	return ""
}

type ICpuContext interface {
	GetRegValue(reg string) string
}

func ICpuContextWithPtr(anyPtr any) ICpuContext {
	return &ICpuContextBase{}
}
