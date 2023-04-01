package logic

import (
	"context"
	"demo/userAccount/api/internal/helpers"
	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"
	"demo/userAccount/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLoginLogic {
	return &AccountLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLoginLogic) AccountLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("Process login account &v", req)
	user, err := l.svcCtx.UserModel.FindByName(l.ctx, req.Name)
	if err != nil {
		return nil, models.ErrNotFound
	}
	passwordIsvalid, msg := helpers.VerifyPassword(user.Password, helpers.HashPassword(req.Password))
	if !passwordIsvalid {
		return &types.LoginResponse{
			Errror: msg,
		}, nil
	}
	tokenString, err := helpers.GenToken(req.Name, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		l.Logger.Errorf("failed to get token")
	}
	return &types.LoginResponse{
		AccessToken: tokenString,
		Errror:      "nil",
	}, nil

}
