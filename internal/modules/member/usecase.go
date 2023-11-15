package member

import (
	"errors"
	"time"
)

type (
	IUseCase interface {
		CreateMember(name, user, pass, email string) (*MemberProfile, error)
	}

	usecase struct {
		memberRepo IMemberRepository
	}
)

func NewUseCase(memberRepo IMemberRepository) IUseCase {
	return usecase{
		memberRepo: memberRepo,
	}
}

// CreateMember implements IUseCase.
func (u usecase) CreateMember(name string, user string, pass string, email string) (*MemberProfile, error) {
	// validate duplicate
	if validEmail, err := u.memberRepo.FindOneByKey("email", email); err != nil {
		return nil, err
	} else if !validEmail.Id.IsZero() {
		return nil, errors.New("email is duplicated")
	}
	if validUser, err := u.memberRepo.FindOneByKey("username", user); err != nil {
		return nil, err
	} else if !validUser.Id.IsZero() {
		return nil, errors.New("username is duplicated")
	}

	result, err := u.memberRepo.InsertOne(name, user, pass, email)
	if err != nil {
		return nil, err
	}
	return &MemberProfile{
		Id:        result.Hex(),
		FullName:  name,
		Username:  user,
		Email:     email,
		CreatedAt: time.Time{},
	}, err
}
