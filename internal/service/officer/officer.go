package officer

import (
	v1 "go-sim/api/backend/v1"
	"go-sim/internal/biz"
)

// OfficerService is a greeter service.
type OfficerService struct {
	v1.UnimplementedOfficerServer

	uc *biz.OfficerUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.OfficerUsecase) *OfficerService {
	return &OfficerService{uc: uc}
}
