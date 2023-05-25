package officer

import (
	"context"
	v1 "go-sim/api/backend/v1"
)

func (s *OfficerService) Info(ctx context.Context, req *v1.InfoReq) (*v1.OfficerInfo, error) {
	officer, err := s.uc.Info(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	response := &v1.OfficerInfo{
		Id:        officer.Id,
		Username:  officer.Username,
		Name:      officer.Name,
		Mobile:    officer.Mobile,
		Email:     officer.Email,
		Avatar:    officer.Avatar,
		Status:    officer.Status,
		CreatedAt: officer.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: officer.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return response, nil
}
