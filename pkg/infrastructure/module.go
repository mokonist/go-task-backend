package infrastructure

import "go.uber.org/fx"

var Module = fx.Module("infrastructure", fx.Provide(newDB, newGrpcServer))
