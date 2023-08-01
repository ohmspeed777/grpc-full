package user

import (
	"app/internal/core/domain"

	"github.com/jinzhu/copier"
)

type transformer struct{}

func (t *transformer) toSingUpEntity(r SignUpReq) domain.SignUpReq {
	e := domain.SignUpReq{}
	copier.Copy(&e, &r)
	return e
}

func (t *transformer) toSingInEntity(r SignInReq) domain.SignInReq {
	e := domain.SignInReq{}
	copier.Copy(&e, &r)
	return e
}

func (t *transformer) toSingInJSON(r domain.SignInRes) SignInRes {
	e := SignInRes{}
	copier.Copy(&e, &r)
	return e
}
