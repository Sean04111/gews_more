package login

import (
	"context"

	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
//使用明文首先
func (l *LoginLogic) Login(req *types.Reque) (resp *types.Respo, err error) {
	// todo: add your logic here and delete this line
	real,err:=l.svcCtx.UserModel.FindOne(l.ctx,req.Email)
	if err==sqlc.ErrNotFound{
		return &types.Respo{
			Error_code: 2,
		},err
	}else if req.Password!=real.Password{
		return &types.Respo{
			Error_code: 1,
		},nil
	}
	return &types.Respo{
		Error_code: 0,
		Data: types.Data{
			Uid: int(real.Uid),	
			Name:real.Name,
			Email: real.Email,
		},
	},nil
}
