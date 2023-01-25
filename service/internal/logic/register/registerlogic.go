package register

import (
	"context"

	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"
	"gews_more/service/model"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Registerreque) (resp *types.Registerrespo, err error) {
	// todo: add your logic here and delete this line
	var NewUser model.User
	NewUser.Email = req.Emial
	NewUser.Name = req.Name
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return
	}
	NewUser.Password = string(hashedpassword)

	//计划使用redis储存用户个数然后来获取uid
	//使用1代替（非主键）
	NewUser.Uid = 1

	_, error := l.svcCtx.UserModel.Insert(l.ctx, &NewUser)
	if error != nil {
		return &types.Registerrespo{
			Error_code: 1,
		}, nil
	}
	return &types.Registerrespo{
		Error_code: 0,
		Data: types.Registerdata{
			Uid:   int(NewUser.Uid),
			Email: NewUser.Email,
			Name:  NewUser.Name,
		},
	}, nil
}
