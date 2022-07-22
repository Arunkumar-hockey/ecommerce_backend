package buyer

import (
	"TFP/utils"
	"strings"
	"time"
)

type User struct {
	Userid           string           `json:"user_id"`
	Status           int              `json:"status"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	ContactPerson    string           `json:"contact_person"`
	FaxNo            string           `json:"faxno"`
	Logo             string           `json:"logo_cdn"`
	Banner           string           `json:"bannercdn"`
	WebsiteAddress   string           `json:"website_address"`
	UserAddress      UserAddress      `json:"address"`
	SocialProfileURL SocialProfileURL `json:"social_profile_url"`
	Description      string           `json:"description"`
	Email            string           `json:"email"`
	Password         string           `json:"password"`
	Usertype         string           `json:"user_type"`
	Usertypecategory string           `json:"category"`
	TypeUser         int              `json:"type"`
	Mobile           string           `json:"mobile"`
	Otp              string           `json:"otp"`
	OTPExpires       time.Time        `json:"otp_expires"`
	CompanyName      string           `json:"company_name"`

	//Admin
	CompanyTagLine      []string           `json:"tagline"`
	RegisteredIn        string             `json:"registeredin"`
	CompanyTeamSize     string             `json:"teamsize"`
	Certification       string             `json:"certification"`
	OperationalAddress  OperationalAddress `json:"operational_Address"`
	CorporateAddress    CorporateAddress   `json:"corporate_address"`
	Overview            []string           `json:"overview"`
	MetaKeywords        []string           `json:"metakeywords"`
	MetaDescription     []string           `json:"metadescription"`
	TaxNumber           string             `json:"taxnumber"`
	Policy              Policy             `json:"policy"`
	KYCInfo             KYCInfo            `json:"kyc_info"`
	Shareholdername     []string           `json:"kycshareholdername"`
	Shareholderposition []string           `json:"kycshareholderposition"`
	Shareholderidnumber []string           `json:"kycshareholderidno"`
	Idproofstatus       string             `json:"idproofstatus"`
	Selfieproofstatus   string             `json:"selfieproofstatus"`
	Idproofmessage      string             `json:"idproofmessage"`
	Selfieproofmessage  string             `json:"selfieproofmessage"`
	Message             string             `json:"message"`
	Planname            string             `json:"planname"`
	Monthlyplan         string             `json:"monthlyplan"`
	Annualplan          string             `json:"annualplan"`
	Currency            string             `json:"currency"`
	Validity            string             `json:"validity"`
	Action              string             `json:"action"`
	ExpiryDate          string             `json:"expirydate"`
	PlanCreatedAt       string             `json:"plan_created_at"`
	LimitOrder          string             `json:"limitorder"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
}

type UserAddress struct {
	StreetName string `json:"street_name"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PinCode    string `json:"pincode"`
}

type SocialProfileURL struct {
	Gplusurl    string `json:"gplus_url"`
	Linkedinurl string `json:"linkedin_url"`
	Fburl       string `json:"fburl"`
}

type KYCInfo struct {
	Kyccompanyname      string `json:"kyccompanyname"`
	Kycauthorisedperson string `json:"kycauthorise"`
	Kycidnumber         string `json:"kycidnumber"`
	Kycregno            string `json:"kycregno"`
	Kycdocument         string `json:"kycdocument"`
	Kycposition         string `json:"kycposition"`
	Kycprooftype        string `json:"kycprooftype"`
	Kycidprooffront     string `json:"kycidprooffront"`
	Kycidproofback      string `json:"kycidproofback"`
	Kycselfie           string `json:"kycselfie"`
	Kycstatus           string `json:"kycstatus"`
}

type Policy struct {
	ReturnPolicy   []string `json:"returnpolicy"`
	ShippingPolicy []string `json:"shippingpolicy"`
	PrivacyPolicy  []string `json:"privacypolicy"`
}

type CorporateAddress struct {
	StreetName string `json:"street_name"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postcode"`
	PhoneNo    string `json:"phone_no"`
}

type OperationalAddress struct {
	StreetName string `json:"street_name"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postcode"`
	PhoneNo    string `json:"phone_no"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyOTP struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func (user *User) Validate() *utils.RestErr {
	if strings.TrimSpace(user.FirstName) == "" {
		return utils.NewBadRequestError("first name required")
	}
	if strings.TrimSpace(user.LastName) == "" {
		return utils.NewBadRequestError("last name required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return utils.NewBadRequestError("email required")
	}
	if strings.TrimSpace(user.Password) == "" {
		return utils.NewBadRequestError("password required")
	}
	if strings.TrimSpace(user.Mobile) == "" {
		return utils.NewBadRequestError("mobile number required")
	}
	if strings.TrimSpace(user.CompanyName) == "" {
		return utils.NewBadRequestError("company name required")
	}
	//if strings.TrimSpace(user.Country) == "" {
	//	return utils.NewBadRequestError("country name required")
	//}
	return nil
}
