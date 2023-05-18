package service

import (
	"github.com/google/wire"
	"go-sim/internal/service/officer"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(officer.NewGreeterService)
