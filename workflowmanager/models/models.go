package models

type SetupPlugin interface {
	Description()
	Setup()
}

type StressPlugin interface {
	Description()
	Stress()
}

type CleanUpPlugin interface {
	Description()
	Clean()
}
