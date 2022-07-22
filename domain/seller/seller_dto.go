package seller

import (
	"TFP/utils"
	"strings"
	"time"
)

type Seller struct {
	Sellerid           string    `json:"seller_id"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Usertype           string    `json:"user_type"`
	Industry           string    `json:"industry"`
	Branch             string    `json:"branch"`
	Swiftcode          string    `json:"swiftcode"`
	Mobile             string    `json:"mobile"`
	CompanyName        string    `json:"company_name"`
	Address            Address   `json:"address"`
	Status             int       `json:"status"`
	ContactPerson      string    `json:"contact_person"`
	Description        string    `json:"description"`
	FaxNo              string    `json:"faxno"`
	ServiceDescription []string  `json:"service_description"`
	ServiceTitle       []string  `json:"service_title"`
	WebsiteAddress     string    `json:"website_address"`
	BranchAddress      []string  `json:"branch_address"`
	Logocdn            string    `json:"logocdn"`
	Bannercdn          string    `json:"bannercdn"`
	Position           []string  `json:"position"`
	Rmposition         []int     `json:"rmposition"`
	AccountNo          string    `json:"account_no"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type Address struct {
	StreetName string `json:"street_name"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Pincode    string `json:"pincode"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (seller *Seller) Validate() *utils.RestErr {
	if strings.TrimSpace(seller.FirstName) == "" {
		return utils.NewBadRequestError("first name required")
	}
	if strings.TrimSpace(seller.LastName) == "" {
		return utils.NewBadRequestError("last name required")
	}
	if strings.TrimSpace(seller.Email) == "" {
		return utils.NewBadRequestError("email required")
	}
	if strings.TrimSpace(seller.Password) == "" {
		return utils.NewBadRequestError("password required")
	}
	if strings.TrimSpace(seller.Mobile) == "" {
		return utils.NewBadRequestError("mobile number required")
	}
	if strings.TrimSpace(seller.CompanyName) == "" {
		return utils.NewBadRequestError("company name required")
	}
	if strings.TrimSpace(seller.Address.StreetName) == "" {
		return utils.NewBadRequestError("street name required")
	}
	if strings.TrimSpace(seller.Address.City) == "" {
		return utils.NewBadRequestError("city required")
	}
	if strings.TrimSpace(seller.Address.State) == "" {
		return utils.NewBadRequestError("state required")
	}
	if strings.TrimSpace(seller.Address.Country) == "" {
		return utils.NewBadRequestError("country required")
	}
	if strings.TrimSpace(seller.Address.Pincode) == "" {
		return utils.NewBadRequestError("pincode number required")
	}

	return nil
}
