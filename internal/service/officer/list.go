package officer

import (
	"context"
	v1 "go-sim/api/backend/v1"
)

func (s *OfficerService) List(ctx context.Context, req *v1.ListReq) (*v1.ListRep, error) {
	//if err := validate.ValidateStructCN(req); err != nil {
	//	return nil, v1.ErrorParamsValidation(err.Error())
	//}
	params := make(map[string]interface{})
	params["username"] = req.GetUsername()
	params["name"] = req.Name
	params["mobile"] = req.Mobile
	params["status"] = req.Status
	officers, count, err := s.uc.List(ctx, 0, 10, params)
	if err != nil {
		return nil, err
	}
	response := &v1.ListRep{Total: count}
	for _, v := range officers {
		response.List = append(response.List, &v1.OfficerInfo{
			Id:        v.Id,
			Username:  v.Username,
			Name:      v.Name,
			Mobile:    v.Mobile,
			Email:     v.Email,
			Avatar:    v.Avatar,
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return response, nil
}
