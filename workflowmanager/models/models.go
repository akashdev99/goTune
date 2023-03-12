package models

type SetupPlugin interface {
	Description()
	Setup(map[string]interface{}, string)
}

type StressPlugin interface {
	Description()
	Stress()
}

type CleanUpPlugin interface {
	Description()
	Clean()
}
