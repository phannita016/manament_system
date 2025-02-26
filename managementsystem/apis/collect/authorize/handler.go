package authorize

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/apis/validate"
	"github.com/phannita016/management/dtos"
	"github.com/phannita016/management/types"
)

func (h *handleAuthorize) Login(c echo.Context) error {
	body, err := validate.BodyParser[dtos.LoginRequest](c)
	if err != nil {
		return err
	}

	check := h.CheckPermission(body.Username, body.Password)
	if !check {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(errors.New("permission denied"))
	}

	id := c.Response().Header().Get(echo.HeaderXRequestID)
	genToken, err := types.GenerateToken(h.secret, body.Username, body.Password, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(err)
	}

	accessToken, err := genToken.AccessToken()
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(err)
	}

	refreshToken, err := genToken.RefreshToken()
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(err)
	}

	res := dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// set ID in memory for token usage verification
	h.cacheStore.Set(id, "set ID successfully")
	return c.JSON(http.StatusOK, res)
}

func (h handleAuthorize) RefreshToken(c echo.Context) error {
	body, err := validate.BodyParser[dtos.RefreshTokenRequest](c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).WithInternal(err)
	}

	claims, err := types.ParseWithClaims(h.secret, body.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err).WithInternal(err)
	}

	if !claims.IsRefresh() {
		return echo.NewHTTPError(http.StatusNotAcceptable).WithInternal(jwt.ErrInvalidKey)
	}

	if _, ok := h.cacheStore.Get(claims.ClaimID); !ok {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(errors.New("token not found"))
	}

	genToken, err := types.GenerateToken(h.secret, claims.Username, claims.Password, claims.ClaimID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(err)
	}

	accessToken, err := genToken.AccessToken()
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(err)
	}

	res := dtos.RefreshTokenResponse{
		AccessToken: accessToken,
	}

	return c.JSON(http.StatusOK, res)
}

func (h handleAuthorize) Logout(c echo.Context) error {
	token, ok := types.Restricted(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(errors.New("parse token error"))
	}

	if _, ok := h.cacheStore.Get(token.ClaimID); !ok {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(errors.New("token not found"))
	}

	// remove ID from memory to stop token verification
	if ok := h.cacheStore.Delete(token.ClaimID); !ok {
		return echo.NewHTTPError(http.StatusInternalServerError).WithInternal(errors.New("error delete ID in cache"))
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "OK", "message": "logout succesfully"})
}

func (h handleAuthorize) CheckPermission(username, password string) bool {
	return username == "admin" && password == "P@ssw0rd"
}
