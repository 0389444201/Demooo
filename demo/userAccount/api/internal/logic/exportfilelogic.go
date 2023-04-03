package logic

import (
	"context"

	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"
	"demo/userAccount/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportFileLogic {
	return &ExportFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportFileLogic) ExportFile() (*types.ExportFile, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserModel.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	arr := SetValues(data)

	err = utils.ExportFile(arr)
	if err != nil {
		return nil, err
	}

	return &types.ExportFile{
		Errror: "Successfull",
	}, nil
}
