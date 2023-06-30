// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/suyuan32/simple-admin-member-rpc/ent/member"
	"github.com/suyuan32/simple-admin-member-rpc/ent/memberrank"
	"github.com/suyuan32/simple-admin-member-rpc/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-member-rpc/ent/token"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type MemberPager struct {
	Order  member.OrderOption
	Filter func(*MemberQuery) (*MemberQuery, error)
}

// MemberPaginateOption enables pagination customization.
type MemberPaginateOption func(*MemberPager)

// DefaultMemberOrder is the default ordering of Member.
var DefaultMemberOrder = Desc(member.FieldID)

func newMemberPager(opts []MemberPaginateOption) (*MemberPager, error) {
	pager := &MemberPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultMemberOrder
	}
	return pager, nil
}

func (p *MemberPager) ApplyFilter(query *MemberQuery) (*MemberQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// MemberPageList is Member PageList result.
type MemberPageList struct {
	List        []*Member    `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (m *MemberQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...MemberPaginateOption,
) (*MemberPageList, error) {

	pager, err := newMemberPager(opts)
	if err != nil {
		return nil, err
	}

	if m, err = pager.ApplyFilter(m); err != nil {
		return nil, err
	}

	ret := &MemberPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := m.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		m = m.Order(pager.Order)
	} else {
		m = m.Order(DefaultMemberOrder)
	}

	m = m.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type MemberRankPager struct {
	Order  memberrank.OrderOption
	Filter func(*MemberRankQuery) (*MemberRankQuery, error)
}

// MemberRankPaginateOption enables pagination customization.
type MemberRankPaginateOption func(*MemberRankPager)

// DefaultMemberRankOrder is the default ordering of MemberRank.
var DefaultMemberRankOrder = Desc(memberrank.FieldID)

func newMemberRankPager(opts []MemberRankPaginateOption) (*MemberRankPager, error) {
	pager := &MemberRankPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultMemberRankOrder
	}
	return pager, nil
}

func (p *MemberRankPager) ApplyFilter(query *MemberRankQuery) (*MemberRankQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// MemberRankPageList is MemberRank PageList result.
type MemberRankPageList struct {
	List        []*MemberRank `json:"list"`
	PageDetails *PageDetails  `json:"pageDetails"`
}

func (mr *MemberRankQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...MemberRankPaginateOption,
) (*MemberRankPageList, error) {

	pager, err := newMemberRankPager(opts)
	if err != nil {
		return nil, err
	}

	if mr, err = pager.ApplyFilter(mr); err != nil {
		return nil, err
	}

	ret := &MemberRankPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := mr.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		mr = mr.Order(pager.Order)
	} else {
		mr = mr.Order(DefaultMemberRankOrder)
	}

	mr = mr.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := mr.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type OauthProviderPager struct {
	Order  oauthprovider.OrderOption
	Filter func(*OauthProviderQuery) (*OauthProviderQuery, error)
}

// OauthProviderPaginateOption enables pagination customization.
type OauthProviderPaginateOption func(*OauthProviderPager)

// DefaultOauthProviderOrder is the default ordering of OauthProvider.
var DefaultOauthProviderOrder = Desc(oauthprovider.FieldID)

func newOauthProviderPager(opts []OauthProviderPaginateOption) (*OauthProviderPager, error) {
	pager := &OauthProviderPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultOauthProviderOrder
	}
	return pager, nil
}

func (p *OauthProviderPager) ApplyFilter(query *OauthProviderQuery) (*OauthProviderQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// OauthProviderPageList is OauthProvider PageList result.
type OauthProviderPageList struct {
	List        []*OauthProvider `json:"list"`
	PageDetails *PageDetails     `json:"pageDetails"`
}

func (op *OauthProviderQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...OauthProviderPaginateOption,
) (*OauthProviderPageList, error) {

	pager, err := newOauthProviderPager(opts)
	if err != nil {
		return nil, err
	}

	if op, err = pager.ApplyFilter(op); err != nil {
		return nil, err
	}

	ret := &OauthProviderPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := op.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		op = op.Order(pager.Order)
	} else {
		op = op.Order(DefaultOauthProviderOrder)
	}

	op = op.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := op.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type TokenPager struct {
	Order  token.OrderOption
	Filter func(*TokenQuery) (*TokenQuery, error)
}

// TokenPaginateOption enables pagination customization.
type TokenPaginateOption func(*TokenPager)

// DefaultTokenOrder is the default ordering of Token.
var DefaultTokenOrder = Desc(token.FieldID)

func newTokenPager(opts []TokenPaginateOption) (*TokenPager, error) {
	pager := &TokenPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultTokenOrder
	}
	return pager, nil
}

func (p *TokenPager) ApplyFilter(query *TokenQuery) (*TokenQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// TokenPageList is Token PageList result.
type TokenPageList struct {
	List        []*Token     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (t *TokenQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...TokenPaginateOption,
) (*TokenPageList, error) {

	pager, err := newTokenPager(opts)
	if err != nil {
		return nil, err
	}

	if t, err = pager.ApplyFilter(t); err != nil {
		return nil, err
	}

	ret := &TokenPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := t.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		t = t.Order(pager.Order)
	} else {
		t = t.Order(DefaultTokenOrder)
	}

	t = t.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
