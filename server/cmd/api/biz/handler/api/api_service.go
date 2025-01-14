// Code generated by hertz generator.

package api

import (
	"context"
	"time"

	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/model/server/cmd/api"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/config"
	"github.com/CyanAsterisk/FreeCar/server/shared/consts"
	"github.com/CyanAsterisk/FreeCar/server/shared/errno"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/auth"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/car"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/profile"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/trip"
	"github.com/CyanAsterisk/FreeCar/server/shared/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	// rpc to get accountID
	resp, err := config.GlobalAuthClient.Login(ctx, &auth.LoginRequest{Code: req.Code})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	// create a JWT
	j := middleware.NewJWT()
	claims := middleware.CustomClaims{
		ID: resp.AccountId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.ThirtyDays,
			Issuer:    consts.JWTIssuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		errno.SendResponse(c, errno.GenerateTokenFail, nil)
		return
	}
	// return token
	errno.SendResponse(c, errno.Success, api.LoginResponse{
		Token:     token,
		ExpiredAt: time.Now().Unix() + consts.ThirtyDays,
	})
}

// GetUserInfo .
// @router /auth/info [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalAuthClient.GetUser(ctx, &auth.GetUserRequest{AccontId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, api.UserInfo{
		AccountId:   resp.AccountId,
		Username:    resp.Username,
		AvatarUrl:   resp.AvatarUrl,
		PhoneNumber: resp.PhoneNumber,
	})
}

// UpdateUserInfo .
// @router /auth/info [POST]
func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	_, err = config.GlobalAuthClient.UpdateUser(ctx, &auth.UpdateUserRequest{
		AccountId:   aid.(int64),
		Username:    req.Username,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, api.UpdateUserResponse{})
}

// UploadAvatar .
// @router /auth/avatar [POST]
func UploadAvatar(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalAuthClient.UploadAvatar(ctx, &auth.UploadAvatarRequset{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, api.UploadAvatarResponse{UploadUrl: resp.UploadUrl})
}

// CreateCar .
// @router /car [POST]
func CreateCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalCarClient.CreateCar(ctx, &car.CreateCarRequest{
		AccountId: aid.(int64),
		PlateNum:  req.PlateNum,
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetCar .
// @router /car [GET]
func GetCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalCarClient.GetCar(ctx, &car.GetCarRequest{
		AccountId: aid.(int64),
		Id:        req.Id,
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetCars .
// @router /cars [GET]
func GetCars(ctx context.Context, c *app.RequestContext) {
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.AuthorizeFail, nil)
		return
	}
	resp, err := config.GlobalCarClient.GetCars(ctx, &car.GetCarsRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetProfile .
// @router /profile [GET]
func GetProfile(ctx context.Context, c *app.RequestContext) {
	var err error

	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.GetProfile(ctx, &profile.GetProfileRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// SubmitProfile .
// @router /profile [POST]
func SubmitProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SubmitProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.SubmitProfile(ctx, &profile.SubmitProfileRequest{
		AccountId: aid.(int64),
		Identity: &profile.Identity{
			LicNumber:       req.Identity.LicNumber,
			Name:            req.Identity.Name,
			Gender:          profile.Gender(req.Identity.Gender),
			BirthDateMillis: req.Identity.BirthDateMillis,
		},
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// ClearProfile .
// @router /profile [DELETE]
func ClearProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ClearProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.ClearProfile(ctx, &profile.ClearProfileRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetProfilePhoto .
// @router /profile/photo [GET]
func GetProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.GetProfilePhoto(ctx, &profile.GetProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CreateProfilePhoto .
// @router /profile/photo [POST]
func CreateProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.CreateProfilePhoto(ctx, &profile.CreateProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CompleteProfilePhoto .
// @router /profile/photo/complete [POST]
func CompleteProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.CompleteProfilePhoto(ctx, &profile.CompleteProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// ClearProfilePhoto .
// @router /profile/photo [DELETE]
func ClearProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.ClearProfilePhoto(ctx, &profile.ClearProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CreateTrip .
// @router /trip [POST]
func CreateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalTripClient.CreateTrip(ctx, &trip.CreateTripRequest{
		Start: &trip.Location{
			Latitude:  req.Start.Latitude,
			Longitude: req.Start.Longitude,
		},
		CarId:     req.CarId,
		AvatarUrl: req.AvatarUrl,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetTrip .
// @router /trip/:id [GET]
func GetTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetTripRequest
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}
	req.Id = c.Param("id")

	resp, err := config.GlobalTripClient.GetTrip(ctx, &trip.GetTripRequest{
		Id:        req.Id,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetTrips .
// @router /trips [GET]
func GetTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetTripsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := config.GlobalTripClient.GetTrips(ctx, &trip.GetTripsRequest{
		Status:    trip.TripStatus(req.Status),
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// UpdateTrip .
// @router /trip/:id [PUT]
func UpdateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}
	req.Id = c.Param(consts.ID)

	resp, err := config.GlobalTripClient.UpdateTrip(ctx, &trip.UpdateTripRequest{
		Id: req.Id,
		Current: &trip.Location{
			Latitude:  req.Current.Latitude,
			Longitude: req.Current.Longitude,
		},
		EndTrip:   req.EndTrip,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}
