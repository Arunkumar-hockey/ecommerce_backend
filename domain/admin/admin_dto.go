package admin

import (
	"TFP/utils"
	"strings"
	"time"
)

type Admin struct {
	AdminID        string    `json:"admin_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Mobile         string    `json:"mobile"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Pattern        string    `json:"pattern"`
	UserType       string    `json:"user_type"`
	NewPattern     string    `json:"new_pattern"`
	ConfirmPattern string    `json:"confirm_pattern"`
	Privileges     []string  `json:"privileges"`
	Status         int       `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Admin) Validate() *utils.RestErr {
	if strings.TrimSpace(a.FirstName) == "" {
		return utils.NewBadRequestError("first name required")
	}
	if strings.TrimSpace(a.LastName) == "" {
		return utils.NewBadRequestError("last name required")
	}
	if strings.TrimSpace(a.Mobile) == "" {
		return utils.NewBadRequestError("mobile number required")
	}
	if strings.TrimSpace(a.Email) == "" {
		return utils.NewBadRequestError("email required")
	}
	if strings.TrimSpace(a.Password) == "" {
		return utils.NewBadRequestError("password required")
	}
	if strings.TrimSpace(a.UserType) == "" {
		a.UserType = "admin"
	}
	return nil
}
