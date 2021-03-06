package auth

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pobyzaarif/pos_lite/business"
	"github.com/pobyzaarif/pos_lite/business/user"
)

type (
	JwtClaims struct {
		UserID int    `json:"user_id"`
		Role   string `json:"role"`

		jwt.StandardClaims
	}

	service struct {
		userService user.Service
	}

	Service interface {
		VerifyLogin(ic business.InternalContext, email string, salt string, plainPassword string) (getUser user.User, validPassword bool)

		GenerateToken(jwtSign string, userID int, userRole string) (signedToken string, err error)
	}
)

func NewService(userService user.Service) Service {
	return &service{
		userService,
	}
}

func (s *service) VerifyLogin(ic business.InternalContext, email string, salt, plainPassword string) (getUser user.User, validPassword bool) {
	getUser, err := s.userService.FindByEmail(ic, email)
	if err != nil {
		return
	}

	if createPasswordHash(salt, plainPassword) != getUser.Password {
		return
	}

	validPassword = true

	return
}

func createPasswordHash(salt, plainPassword string) (passwordHash string) {
	hasher := md5.New()
	hasher.Write([]byte(salt + plainPassword))

	return hex.EncodeToString(hasher.Sum(nil))
}

func newJWTClaims(userID int, role string, issuedAt int64, expiredAt int64) JwtClaims {
	return JwtClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedAt,
			ExpiresAt: expiredAt,
		},
	}
}

func (s *service) GenerateToken(jwtSign string, userID int, userRole string) (signedToken string, err error) {
	timeNow := time.Now()
	claims := newJWTClaims(userID, userRole, timeNow.Unix(), timeNow.Add(time.Hour*24).Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(jwtSign))
	if err != nil {
		return
	}

	return signedToken, nil
}
