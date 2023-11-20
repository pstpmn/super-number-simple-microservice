package member

import (
	"errors"
	"super-number-simple-microservice/pkg"
	"time"
)

type (
	IUseCase interface {
		CreateMember(name, user, pass, email string) (*Profile, error)
		Authentication(user, pass string) (*CredentialCombindProfile, error)
	}
	usecase struct {
		memberRepo IMemberRepository
		jwt        pkg.IJwt
		hash       pkg.IHash
	}
)

func (ob usecase) Authentication(user string, pass string) (*CredentialCombindProfile, error) {
	member, err := ob.memberRepo.FindOneByKey("username", user)
	if err != nil {
		return nil, err
	} else if member.Id.IsZero() {
		return nil, errors.New("invalid username or password")
	} else if !ob.hash.CompareBcrypt([]byte(pass), []byte(member.Password)) {
		return nil, errors.New("invalid username or password")
	}
	profile := Profile{
		Id:        member.Id.Hex(),
		FullName:  member.FullName,
		Username:  member.Username,
		Email:     member.Email,
		CreatedAt: time.Time{},
	}
	return &CredentialCombindProfile{&profile, &Credential{
		AccessToken: ob.jwt.SignToken(profile),
		CreatedAt:   time.Now(),
	}}, nil
}

func NewUseCase(memberRepo IMemberRepository, jwt pkg.IJwt, hash pkg.IHash) IUseCase {
	return usecase{
		memberRepo: memberRepo,
		jwt:        jwt,
		hash:       hash,
	}
}

// CreateMember implements IUseCase.
func (u usecase) CreateMember(name string, user string, pass string, email string) (*Profile, error) {
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

	// encript password
	encryptPass, err := u.hash.Bcrypt([]byte(pass))
	if err != nil {
		return nil, err
	}

	result, err := u.memberRepo.InsertOne(name, user, encryptPass, email)
	if err != nil {
		return nil, err
	}
	return &Profile{
		Id:        result.Hex(),
		FullName:  name,
		Username:  user,
		Email:     email,
		CreatedAt: time.Time{},
	}, err
}
