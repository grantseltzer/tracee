package selftest

type KernelVersion struct {
	Major, Minor, Patch uint 
}

type Selftester interface {
	MinimumKernel() KernelVersion
	Start() error
	Stop() error 
	Success() bool
}
