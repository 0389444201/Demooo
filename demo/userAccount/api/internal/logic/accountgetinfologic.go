package logic

import (
	"context"
	"time"

	"demo/userAccount/api/internal/middlerware"
	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountGetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountGetInfoLogic {
	return &AccountGetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountGetInfoLogic) AccountGetInfo(req *types.InfoRequest) (resp *types.InfoResponse, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("Process login account &v", req.Name)
	value := time.Now().Format(time.RFC3339)
	middlerware.AccessMiddleware(l.ctx, req.Name+" was find", value)
	info, err := l.svcCtx.UserModel.FindByName(l.ctx, req.Name)
	if err != nil {
		l.Logger.Errorf("error: not found")
	}
	return &types.InfoResponse{
		Name:   info.Name,
		Email:  info.Email,
		Gender: info.Gender,
	}, nil
}
