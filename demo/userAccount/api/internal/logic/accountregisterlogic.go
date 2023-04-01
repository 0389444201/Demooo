package logic

import (
	"context"
	"demo/userAccount/api/internal/helpers"
	"demo/userAccount/api/internal/middlerware"
	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"
	"demo/userAccount/models"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountRegisterLogic {
	return &AccountRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountRegisterLogic) AccountRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("Process register %v", req)
	password := helpers.HashPassword(req.Password)
	req.Password = password
	user := models.UserTable{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Gender:   req.Gender,
	}
	value := time.Now().Format(time.RFC3339)
	middlerware.AccessMiddleware(l.ctx, req.Name+"was registered", value)
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &user)

	if err != nil {
		l.Logger.Errorf("failed to register account")
	}
	return &types.RegisterResponse{
		Name:   req.Name,
		Email:  req.Email,
		Gender: req.Gender,
	}, nil
}
