package snappost

import (
	"context"
	"database/sql"

	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"
	"gews_more/service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsnapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsnapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsnapLogic {
	return &PostsnapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsnapLogic) Postsnap(req *types.Snappostreque) (resp *types.Snappostrespo, err error) {
	// todo: add your logic here and delete this line
	//在没有redis的时候的措施
	//速度和性能应该比redis差
	num:=0
	for{
		_,err:= l.svcCtx.SnapModel.FindOne(l.ctx,int64(num))
		if err==model.ErrNotFound{
			break
		}
		num++
	}
	insertsnap:=model.Snaps{
		Sid: int64(num),
		Speaker: req.Speaker,
		Message: req.Message,
		At: sql.NullString{
			String: req.Date,
		},
		Date: req.Date,
	}
	_,er:=l.svcCtx.SnapModel.Insert(l.ctx,&insertsnap)
	if er!=nil{
		return nil,er
	}else{
		return &types.Snappostrespo{
			Error_code: 0,
		},nil
	}
}
