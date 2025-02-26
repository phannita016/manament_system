package management

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/apis/validate"
	"github.com/phannita016/management/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *handleManagement) Create(c echo.Context) error {
	ctx := c.Request().Context()

	body, err := validate.BodyParser[dtos.ManagementRequest](c)
	if err != nil {
		return err
	}

	if err := h.managementStore.Create(ctx, *body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create new management").WithInternal(err)
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "success"})
}

func (h *handleManagement) Update(c echo.Context) error {
	ctx := c.Request().Context()

	body, err := validate.BodyParser[dtos.ManagementRequest](c)
	if err != nil {
		return err
	}

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest).WithInternal(errors.New("id is missing"))
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id").WithInternal(err)
	}

	_, err = h.managementStore.FindByID(ctx, objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "management not found").WithInternal(err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "error to get management by id").WithInternal(err)
	}

	model := dtos.Management{
		ID:         objectID,
		Name:       body.Name,
		Nickname:   body.Nickname,
		Gender:     body.Gender,
		Age:        body.Age,
		Role:       body.Role,
		UpdateDate: time.Now(),
	}

	if err := h.managementStore.Update(ctx, model); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update").WithInternal(err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "success", "data": model})
}

func (h *handleManagement) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest).WithInternal(errors.New("id is missing"))
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id").WithInternal(err)
	}

	filter := bson.M{"_id": objectID}

	_, err = h.managementStore.FindByID(ctx, objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "management not found").WithInternal(err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "error to get management by id").WithInternal(err)
	}
	if err := h.managementStore.Delete(ctx, filter); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete").WithInternal(err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "success"})
}

func (h *handleManagement) List(c echo.Context) error {
	ctx := c.Request().Context()

	results, err := h.managementStore.GetAll(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get list").WithInternal(err)
	}

	return c.JSON(http.StatusOK, results)
}
