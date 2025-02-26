package management

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type handleManagement struct {
	managementStore store.Management
}

func New(e *echo.Group, database *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	h := handleManagement{
		managementStore: store.NewDatabaseManagement(ctx, database),
	}

	e.GET("", h.List)
	e.POST("", h.Create)
	e.PUT("/:id", h.Update)
	e.DELETE("/:id", h.Delete)
}
