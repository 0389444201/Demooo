package logic

import (
	"context"

	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"
	"demo/userAccount/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountGetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountGetAllLogic {
	return &AccountGetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountGetAllLogic) AccountGetAll(req *types.GetAll) (*[]types.ResponseGetAll, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.UserModel.GetAll(l.ctx)
	msg := types.ResponseGetAll{
		Error: "Error occured while getting DB",
	}
	if err != nil {
		return &[]types.ResponseGetAll{
			msg,
		}, err
	}
	arr := SetValues(result)
	return &arr, nil
}

func SetValues(result []models.UserTable) (arr []types.ResponseGetAll) {
	for i := 0; i < len(result); i++ {
		setValues := types.ResponseGetAll{
			Id:     result[i].Id,
			Name:   result[i].Name,
			Email:  result[i].Email,
			Gender: result[i].Gender,
		}
		arr = append(arr, setValues)
	}
	return arr
}
