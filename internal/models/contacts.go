package models

// Contact contains the unmarshalled data for a "FreshDesk" contact
type Contact struct {
	Active         bool                `json:"active,omitempty" mapstructure:"active,omitempty"`
	Address        string              `json:"address,omitempty" mapstructure:"address,omitempty"`
	Avatar         ContactAvatar       `json:"avatar,omitempty" mapstructure:"avatar,omitempty"`
	CompanyID      int                 `json:"company_id,omitempty" mapstructure:"company_id,omitempty"`
	CreatedAt      string              `json:"created_at,omitempty" mapstructure:"created_at,omitempty"`
	CustomFields   ContactCustomFields `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	Deleted        bool                `json:"deleted,omitempty" mapstructure:"deleted,omitempty"`
	Description    string              `json:"description,omitempty" mapstructure:"description,omitempty"`
	Email          string              `json:"email,omitempty" mapstructure:"email,omitempty"`
	ExternalID     string              `json:"external_id,omitempty" mapstructure:"external_id,omitempty"`
	ID             int                 `json:"id,omitempty,omitempty" mapstructure:"id,omitempty"`
	JobTitle       string              `json:"job_title,omitempty" mapstructure:"job_title,omitempty"`
	Language       string              `json:"language,omitempty" mapstructure:"language,omitempty"`
	Mobile         string              `json:"mobile,omitempty" mapstructure:"mobile,omitempty"`
	Name           string              `json:"name,omitempty" mapstructure:"name,omitempty"`
	OtherCompanies []struct {
		CompanyID      int  `json:"company_id" mapstructure:"company_id"`
		ViewAllTickets bool `json:"view_all_tickets" mapstructure:"view_all_tickets"`
	} `json:"other_companies,omitempty"`
	OtherEmails    []string      `json:"other_emails,omitempty" mapstructure:"other_emails,omitempty"`
	PeopleID       string        `json:"unique_external_id,omitempty" mapstructure:"unique_external_id,omitempty"`
	Phone          string        `json:"phone,omitempty" mapstructure:"phone,omitempty"`
	Tags           []interface{} `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	TimeZone       string        `json:"time_zone,omitempty" mapstructure:"time_zone,omitempty"`
	TwitterID      string        `json:"twitter_id,omitempty" mapstructure:"twitter_id,omitempty"`
	UpdatedAt      string        `json:"updated_at,omitempty" mapstructure:"updated_at,omitempty"`
	ViewAllTickets bool          `json:"view_all_tickets,omitempty" mapstructure:"view_all_tickets,omitempty"`
}

// ContactAvatar contains the subfields for the avatar field of the Contact type
type ContactAvatar struct {
	AvatarUrl   string `json:"avatar_url,omitempty" mapstructure:"avatar_url,omitempty"`
	ContentType string `json:"content_type,omitempty" mapstructure:"content_type,omitempty"`
	Id          int    `json:"id,omitempty" mapstructure:"id,omitempty"`
	Name        string `json:"name,omitempty" mapstructure:"name,omitempty"`
	Size        int    `json:"size,omitempty" mapstructure:"size,omitempty"`
	CreatedAt   string `json:"created_at,omitempty" mapstructure:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty" mapstructure:"updated_at,omitempty"`
}

// ContactCustomFields contains the values for the custom_fields portion of a Contact
type ContactCustomFields struct {
	Benefit           string `json:"benefit,omitempty" mapstructure:"benefit,omitempty"`
	DateOfBirth       string `json:"date_of_birth,omitempty" mapstructure:"date_of_birth,omitempty"`
	EligibilityStatus string `json:"eligibility_status,omitempty" mapstructure:"eligibility_status,omitempty"`
	LoginEmail        string `json:"login_email,omitempty" mapstructure:"login_email,omitempty"`
	LoginPassword     string `json:"login_password,omitempty" mapstructure:"login_password,omitempty"`
	Phone1            string `json:"phone_1,omitempty" mapstructure:"phone_1,omitempty"`
	Phone2            string `json:"phone_2,omitempty" mapstructure:"phone_2,omitempty"`
}
