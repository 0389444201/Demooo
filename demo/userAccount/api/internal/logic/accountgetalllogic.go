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

func (l *AccountGetAllLogic) AccountGetAll(req *types.GetAll) ([]models.Table, error) {
	// todo: add your logic here and delete this line

	result, err := l.svcCtx.UserModel.GetAll(l.ctx)
	if err != nil {
		logx.Infof("", err)
	}
	logx.Info(result)
	return result, nil
}
