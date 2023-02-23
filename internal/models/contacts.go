package models

import "time"

// Contact contains the unmarshalled data for a "FreshDesk" contact
type Contact struct {
	Active         bool                `json:"active,omitempty"`
	Address        string              `json:"address,omitempty"`
	Avatar         ContactAvatar       `json:"avatar,omitempty"`
	CompanyID      int                 `json:"company_id,omitempty"`
	CreatedAt      time.Time           `json:"created_at,omitempty"`
	CustomFields   ContactCustomFields `json:"custom_fields,omitempty"`
	Deleted        bool                `json:"deleted,omitempty"`
	Description    string              `json:"description,omitempty"`
	Email          string              `json:"email,omitempty"`
	ExternalID     string              `json:"external_id,omitempty"`
	FacebookID     string              `json:"facebookID,omitempty"`
	ID             int                 `json:"id,omitempty,omitempty"`
	JobTitle       string              `json:"job_title,omitempty"`
	Language       string              `json:"language,omitempty"`
	Mobile         string              `json:"mobile,omitempty"`
	Name           string              `json:"name,omitempty"`
	OtherCompanies []struct {
		CompanyID      int  `json:"company_id"`
		ViewAllTickets bool `json:"view_all_tickets"`
	} `json:"other_companies,omitempty"`
	OtherEmails    []string      `json:"other_emails,omitempty"`
	PeopleID       string        `json:"unique_external_id,omitempty"`
	Phone          string        `json:"phone,omitempty"`
	Tags           []interface{} `json:"tags,omitempty"`
	TimeZone       string        `json:"time_zone,omitempty"`
	TwitterID      string        `json:"twitter_id,omitempty"`
	UpdatedAt      time.Time     `json:"updated_at,omitempty"`
	ViewAllTickets bool          `json:"view_all_tickets,omitempty"`
}

// ContactAvatar contains the subfields for the avatar field of the Contact type
type ContactAvatar struct {
	AvatarUrl   string     `json:"avatar_url,omitempty"`
	ContentType string     `json:"content_type,omitempty"`
	Id          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Size        int        `json:"size,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
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
