package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	Authorization = "AUTHORIZATION"

	ACCESSTYPE  = "Access"
	REFRESHTYPE = "Refresh"

	ACCESSTOKENEXPIRED  = time.Minute * 15
	REFRESHTOKENEXPIRED = time.Hour * 24
)

type JwtCustomClaims struct {
	secret []byte `json:"-"`

	Username string `json:"username"`
	Password string `json:"password"`

	ClaimID string `json:"claim_id"`
	Types   string `json:"types"`
	jwt.RegisteredClaims
}

func (c *JwtCustomClaims) IsAccess() bool {
	return c.Types == ACCESSTYPE
}

func (c *JwtCustomClaims) IsRefresh() bool {
	return c.Types == REFRESHTYPE
}

func GenerateToken(secret []byte, username, password, id string) (*JwtCustomClaims, error) {
	if secret == nil {
		return nil, jwt.ErrInvalidKey
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		secret:   secret,
		Username: username,
		Password: password,
		ClaimID:  id,
	}

	return claims, nil
}

func (c *JwtCustomClaims) AccessToken() (string, error) {
	c.Types = ACCESSTYPE
	c.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "management_system",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESSTOKENEXPIRED)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	t, err := token.SignedString(c.secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (c *JwtCustomClaims) RefreshToken() (string, error) {
	c.Types = REFRESHTYPE
	c.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "management_system",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(REFRESHTOKENEXPIRED)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	t, err := token.SignedString(c.secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func Restricted(c echo.Context) (*JwtCustomClaims, bool) {
	tokenJwt, ok := c.Get(Authorization).(*jwt.Token)
	if !ok || tokenJwt == nil {
		return nil, false
	}

	cliams, ok := tokenJwt.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, false
	}
	return cliams, ok
}

func ParseWithClaims(secret []byte, tokenStr string) (*JwtCustomClaims, error) {
	tpf := &JwtCustomClaims{}
	tokenJwt, err := jwt.ParseWithClaims(tokenStr, tpf, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !tokenJwt.Valid {
		return nil, err
	}

	cliams, ok := tokenJwt.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	return cliams, nil
}
