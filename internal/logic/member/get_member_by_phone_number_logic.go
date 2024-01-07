package member

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-member-rpc/ent/member"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByPhoneNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByPhoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByPhoneNumberLogic {
	return &GetMemberByPhoneNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberByPhoneNumberLogic) GetMemberByPhoneNumber(in *mms.PhoneNumberReq) (*mms.MemberInfo, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.DB.Member.Query().Where(member.MobileEQ(in.Mobile)).WithRanks().First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.MemberInfo{
		Id:        pointy.GetPointer(result.ID.String()),
		CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:    pointy.GetPointer(uint32(result.Status)),
		Username:  &result.Username,
		Password:  &result.Password,
		Nickname:  &result.Nickname,
		RankId:    &result.RankID,
		RankCode:  &result.Edges.Ranks.Code,
		Mobile:    &result.Mobile,
		Email:     &result.Email,
		Avatar:    &result.Avatar,
	}, nil
}
