package logic

import (
	"context"
	"demo/userAccount/api/internal/helpers"
	"demo/userAccount/api/internal/middlerware"
	"demo/userAccount/models"
	"time"

	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountUpdateLogic {
	return &AccountUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountUpdateLogic) AccountUpdate(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	name := l.ctx.Value("name").(string)
	hashpassword := helpers.HashPassword(req.Password)
	value := time.Now().Format(time.RFC3339)
	middlerware.AccessMiddleware(l.ctx, name+" was updated", value)
	l.svcCtx.UserModel.Update(l.ctx, &models.UserTable{
		Name:     name,
		Password: hashpassword,
		Email:    req.Email,
	})
	return &types.UpdateResponse{
		Name:  name,
		Email: req.Email,
	}, nil
}
