package handlers

import (
	"context"
	"fmt"
	"github.com/eko/gocache/store"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
	"web/app"
	"web/db/models"
	userpb "web/grpc/proto/user/go"
	appError "web/pkg/error"
	"web/request"
	"web/response"
	"web/service"
)

type UserHandler struct {
	*app.App
}

// NewUserHandler ..
func NewUserHandler(server *app.App) UserHandler {
	return UserHandler{server}
}

func (u UserHandler) Create(c echo.Context) error {
	req := new(request.CreateUser)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		//@todo wrap this appError in the error packge
		return appError.NewRequestError("Validation failed", http.StatusBadRequest, "validation_failed")
	}

	db := u.DB
	db.Create(&models.User{
		UserName: req.UserName,
		Phone:    req.Phone,
		City:     req.City,
		Name:     req.Name,
	})
	return nil
}

func (u UserHandler) List(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.user.list", trace.WithAttributes(attribute.String("list", "yes")))
	defer span.End()

	var userCol []models.User
	value, _ := u.Cache.Get("user_list", u.findUsers(ctx), &store.Options{Expiration: 500 * time.Second})
	userCol = (value).([]models.User)
	return c.JSON(200, response.NewUserCollection(userCol))
}

func (u UserHandler) Index(c echo.Context) error {
	return u.Cache.Flush()
	//var userCol []model.User
	//value, _ := u.server.Cache.Get("user_list", &userCol, u.findUsers(), &store.Options{Expiration: 500 * time.Second})
	//if userCol == nil {
	//	userCol = (value).([]model.User)
	//}
	//return c.JSON(200, response.NewUserCollection(userCol))
}
func (u UserHandler) Details(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	con := u.GetGrpcCon("UserServer") //@todo send error also back
	resp, err := service.CallTestGRPCAPI(ctx, con, id)
	//resp, err := service.CallTestRestAPI(u.HttpClient, "http://127.0.0.1:8000/user/test/1", "1" )
	if err != nil {
		panic(err)
	}
	return c.JSON(200, resp)
}
func (u UserHandler) findUsers(ctx context.Context) func() (interface{}, error) {
	return func() (interface{}, error) {
		var users []models.User
		u.DB.WithContext(ctx).Find(&users)
		return users, nil
	}
} //		return &users, nil

func (u UserHandler) RestAPITest(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.user.rest_api_test")
	defer span.End()
	id := c.Param("id")
	fmt.Println("id", id)
	user := &userpb.UserResponse{
		Name:  "Maju",
		City:  "TDPA",
		Phone: "1234444",
	}
	return c.JSON(200, user)
}
