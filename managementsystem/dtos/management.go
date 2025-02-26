package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Management struct {
		ID         primitive.ObjectID `bson:"_id" json:"id"`
		Name       string             `bson:"name" json:"name"`
		Nickname   string             `bson:"nickname" json:"nickname"`
		Gender     string             `bson:"gender" json:"gender"`
		Age        int                `bson:"age" json:"age"`
		Role       string             `bson:"role" json:"role"`
		CreateDate time.Time          `bson:"createDate" json:"createDate"`
		UpdateDate time.Time          `bson:"updateDate" json:"updateDate"`
	}

	ManagementRequest struct {
		Name     string `json:"name" validate:"required,min=3,max=30"`
		Nickname string `json:"nickname" validate:"required,min=3,max=30"`
		Gender   string `json:"gender" validate:"required,oneof=male female"`
		Age      int    `json:"age" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}
)
