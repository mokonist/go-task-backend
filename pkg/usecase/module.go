package usecase

import "go.uber.org/fx"

var Module = fx.Module("usecase",
	fx.Provide(
		newGetAllTasksUseCase,
		newGetTaskUseCase,
		newCreateTaskUseCase,
	))
