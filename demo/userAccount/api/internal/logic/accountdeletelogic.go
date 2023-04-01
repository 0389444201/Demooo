package logic

import (
	"context"
	"demo/userAccount/api/internal/middlerware"
	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountDeleteLogic {
	return &AccountDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountDeleteLogic) AccountDelete(req *types.DeleteRequest) error {
	// todo: add your logic here and delete this line
	name := l.ctx.Value("name").(string)
	value := time.Now().Format(time.RFC3339)
	middlerware.AccessMiddleware(l.ctx, name+" was deleted", value)
	err := l.svcCtx.UserModel.DeleteByName(l.ctx, name)
	if err != nil {
		l.Logger.Errorf("Failed while deleting account, error: %v", err)
		return err
	}

	return nil
}
