package ext

const (
	LifecycleConfigure = LifecyclePhase(0)
	LifecycleEnable    = LifecyclePhase(1)
	LifecycleDisable   = LifecyclePhase(2)
)

type LifecyclePhase int
