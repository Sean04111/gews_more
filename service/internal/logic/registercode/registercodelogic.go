package registercode

import (
	"context"

	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistercodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegistercodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistercodeLogic {
	return &RegistercodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistercodeLogic) Registercode(req *types.Registercodereque) (resp *types.Registercoderespo, err error) {
	// todo: add your logic here and delete this line
	
	return
}
