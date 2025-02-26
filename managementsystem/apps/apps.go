package apps

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/apis/collect/authorize"
	"github.com/phannita016/management/apis/collect/management"
	"github.com/phannita016/management/apis/middleware"
	"github.com/phannita016/management/apis/serve"
	"github.com/phannita016/management/apis/validate"
	"github.com/phannita016/management/driver"
	"github.com/phannita016/management/store"
	"github.com/phannita016/management/store/cache"
	"go.mongodb.org/mongo-driver/mongo"
)

type Apps struct {
	Addr        string
	Ctx         context.Context
	server      *Server
	mongoClient *mongo.Client

	Secret []byte

	Skippers []string

	Cache store.Cache[string]

	Middleware *middleware.Middleware
}

func AppsServer(app *Apps, opts ...Option) func(c context.Context) error {
	app.ApplyApps()

	var e = serve.New(app.Middleware.Authorization(), app.Middleware.AuthorizeWithToken)
	validate.RegisterValidator(e)
	if app.server == nil {
		app.server = NewServer(app.Addr, e)
	}
	defer app.server.Stop(app.Ctx)

	client, err := app.Database()
	if err != nil {
		slog.Error("Failed to connect to database", slog.Any("err", err))
		return func(c context.Context) error {
			return err
		}
	}
	app.mongoClient = client

	app.HandleFunc(e)

	if err := app.server.Run(app.Ctx); err != nil {
		fmt.Println(err)
	}

	return app.server.Stop
}

func (apps *Apps) HandleFunc(e *echo.Echo) {
	authorize.New(e, apps.Secret, apps.Cache)

	g := e.Group("/management")
	management.New(g, apps.mongoClient)
}

func (apps *Apps) Database() (*mongo.Client, error) {
	client, err := driver.NewMongoClient(driver.MongoDriver{
		Hostname: os.Getenv("MONGODB_URI"),
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		PoolSize: 20,
	})
	if err != nil {
		return nil, err
	}

	slog.Info("Connection Database-Server...")
	return client, nil
}

func (apps *Apps) ApplyApps() {
	if apps.Ctx == nil {
		apps.Ctx = context.Background()
	}
	if apps.Skippers == nil {
		apps.Skippers = []string{
			"/authorize/login",
			"/authorize/refresh",
		}
	}
	if apps.Cache == nil {
		apps.Cache = cache.NewMemCache[string](cache.MemTTL(time.Hour * 24))
	}
	if apps.Middleware == nil {
		apps.Middleware = middleware.NewMiddleware(apps.Secret, apps.Cache, apps.Skippers)
	}
}
