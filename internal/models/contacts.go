package models

import "time"

// Contact contains the unmarshalled data for a "FreshDesk" contact
type Contact struct {
	Active         bool                `json:"active"`
	Address        string              `json:"address"`
	Avatar         ContactAvatar       `json:"avatar,omitempty"`
	CompanyID      int                 `json:"company_id,omitempty"`
	CreatedAt      time.Time           `json:"created_at"`
	CustomFields   ContactCustomFields `json:"custom_fields"`
	Deleted        bool                `json:"deleted,omitempty"`
	Description    string              `json:"description,omitempty"`
	Email          string              `json:"email"`
	ExternalID     string              `json:"external_id,omitempty"`
	FacebookID     string              `json:"facebookID,omitempty"`
	ID             int                 `json:"id"`
	JobTitle       string              `json:"job_title,omitempty"`
	Language       string              `json:"language"`
	Mobile         string              `json:"mobile,omitempty"`
	Name           string              `json:"name"`
	OtherCompanies []struct {
		CompanyID      int  `json:"company_id"`
		ViewAllTickets bool `json:"view_all_tickets"`
	} `json:"other_companies,omitempty"`
	OtherEmails    []string      `json:"other_emails"`
	PeopleID       string        `json:"unique_external_id"`
	Phone          string        `json:"phone"`
	Tags           []interface{} `json:"tags,omitempty"`
	TimeZone       string        `json:"time_zone"`
	TwitterID      string        `json:"twitter_id,omitempty"`
	UpdatedAt      time.Time     `json:"updated_at"`
	ViewAllTickets bool          `json:"view_all_tickets,omitempty"`
}

// ContactAvatar contains the subfields for the avatar field of the Contact type
type ContactAvatar struct {
	AvatarUrl   string    `json:"avatar_url"`
	ContentType string    `json:"content_type"`
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Size        int       `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ContactCustomFields contains the values for the custom_fields portion of a Contact
type ContactCustomFields struct {
	Benefit           string `json:"benefit,omitempty"`
	DateOfBirth       string `json:"date_of_birth,omitempty"`
	EligibilityStatus string `json:"eligibility_status,omitempty"`
	LoginEmail        string `json:"login_email,omitempty"`
	LoginPassword     string `json:"login_password,omitempty"`
	Phone1            string `json:"phone_1,omitempty"`
	Phone2            string `json:"phone_2,omitempty"`
}
