package logic

import (
	"context"

	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"

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
	var arr []types.ResponseGetAll
	result, err := l.svcCtx.UserModel.GetAll(l.ctx)

	if err != nil {
		l.Logger.Infof("error occured while get info")
	}
	for i := 0; i < len(result); i++ {
		rs := types.ResponseGetAll{
			Name:  result[i].Name,
			Email: result[i].Email,
		}
		arr = append(arr, rs)
	}
	msg := types.ResponseGetAll{
		Error: "DB is empty",
	}
	if arr == nil {
		return &[]types.ResponseGetAll{
			msg,
		}, nil
	}

	return &arr, nil
}
