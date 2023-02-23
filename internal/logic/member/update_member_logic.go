package member

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/ent"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/mms"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMemberLogic) UpdateMember(in *mms.MemberInfo) (*mms.BaseResp, error) {
	err := l.svcCtx.DB.Member.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
		SetNotEmptyStatus(uint8(in.Status)).
		SetNotEmptyUsername(in.Username).
		SetNotEmptyPassword(in.Password).
		SetNotEmptyNickname(in.Nickname).
		SetNotEmptyRankID(in.RankId).
		SetNotEmptyMobile(in.Mobile).
		SetNotEmptyEmail(in.Email).
		SetNotEmptyAvatar(in.Avatar).
		Exec(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &mms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
