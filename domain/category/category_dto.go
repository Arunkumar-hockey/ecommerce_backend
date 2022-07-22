package category

import (
	"TFP/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type Category struct {
	ID           primitive.ObjectID `bson:"_id"`
	CategoryName string             `json:"category_name"`
	CategoryID   string             `json:"category_id"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

func (category *Category) Validate() *utils.RestErr {
	if strings.TrimSpace(category.CategoryName) == "" {
		return utils.NewBadRequestError("name required")
	}
	return nil
}
