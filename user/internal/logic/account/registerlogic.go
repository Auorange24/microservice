package account

import (
	"context"
	"errors"
	"time"

	"user/internal/model"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	// 根据Mysql数据库连接创建操作user表的model
	userModel := model.NewUserModel(l.svcCtx.Conn)
	user, err := userModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("未找到该用户", err)
		return nil, err
	}
	if user != nil {
		return nil, errors.New("此用户名已经被注册")
	}
	_, err = userModel.Insert(l.ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
		RegisterTime: time.Now(),
		LastLoginTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return
}
