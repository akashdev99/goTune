package models

type SetupPlugin interface {
	Description()
	Setup(map[string]interface{}, string) error
}

type StressPlugin interface {
	Description()
	Stress()
}

type CleanUpPlugin interface {
	Description()
	Clean()
}
