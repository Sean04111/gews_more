package getsnaps

import (
	"context"

	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"
	"gews_more/service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetsnapsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetsnapsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetsnapsLogic {
	return &GetsnapsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetsnapsLogic) Getsnaps(req *types.Snapreque) (resp *types.Snaprespo, err error) {
	// todo: add your logic here and delete this line
	var snapgets []types.Snap
	for i:=0;;i++{
		now,err:= l.svcCtx.SnapModel.FindOne(l.ctx,int64(i))
		if err==model.ErrNotFound{
			break
		}
		if err!=nil{
			return nil,err
		}
		if err==nil{
			buffer:=types.Snap{
				Sid: int(now.Sid),
				Speaker: now.Speaker,
				Message: now.Message,
				Date: now.Date,
				At:now.At.String,
			}
			snapgets = append(snapgets, buffer)
		}
	}
	return &types.Snaprespo{
		Error_code: 0,
		Snaps: snapgets,
	},nil
}
