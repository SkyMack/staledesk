package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SkyMack/staledesk/config"
	"github.com/SkyMack/staledesk/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	ParamNameContactID = "id"
)

type Contacts struct {
	CurrentContacts map[int]models.Contact
}

type FilterContactsResp struct {
	Total   int              `json:"total" mapstructure:"totle"`
	Results []models.Contact `json:"results" mapstructure:"results"`
}

func NewContactsController() *Contacts {
	return &Contacts{
		CurrentContacts: config.Config.Contacts,
	}
}

func (contControl *Contacts) GetAll(ctx *gin.Context) {
	mustMatchEmail := false
	mustMatchMobile := false
	mustMatchPhone := false

	email := ctx.Query("email")
	if len(email) > 0 {
		mustMatchEmail = true
	}
	mobile := ctx.Query("mobile")
	if len(mobile) > 0 {
		mustMatchMobile = true
	}
	phone := ctx.Query("phone")
	if len(phone) > 0 {
		mustMatchPhone = true
	}

	respContacts := []models.Contact{}
	for _, cont := range contControl.CurrentContacts {
		if mustMatchEmail && cont.Email != email {
			continue
		}
		if mustMatchMobile && cont.Mobile != mobile {
			continue
		}
		if mustMatchPhone && cont.Phone != phone {
			continue
		}
		respContacts = append(respContacts, cont)
	}
	ctx.JSON(http.StatusOK, respContacts)
}

func (contControl *Contacts) GetByID(ctx *gin.Context) {
	intID, err := getIntID(ctx)
	if err != nil {
		return
	}

	contact, exists := contControl.CurrentContacts[intID]
	if !exists {
		ctx.JSON(http.StatusNotFound, nil)
	} else {
		ctx.JSON(http.StatusOK, contact)
	}
}

func (contControl *Contacts) Search(ctx *gin.Context) {

}

func (contControl *Contacts) Filter(ctx *gin.Context) {
	queryStr := ctx.Query("query")
	if len(queryStr) == 0 {
		contControl.GetAll(ctx)
	} else {
		var matchingContacts []models.Contact
		validValues := getValidValuesFromFilterQueryString(queryStr)
		for _, cont := range contControl.CurrentContacts {
			log.WithFields(log.Fields{
				"contact.people_id": cont.PeopleID,
			}).Debug("checking for matching people id")
			if validValues[cont.PeopleID] {
				log.Debug("match found")
				matchingContacts = append(matchingContacts, cont)
			}
		}
		resp := FilterContactsResp{
			Total:   len(matchingContacts),
			Results: matchingContacts,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

// TODO: Make this more robust when required
// Currently we expect to only filter by a specific field, with one or more statements separated by OR
func getValidValuesFromFilterQueryString(query string) map[string]bool {
	validValues := make(map[string]bool)
	log.WithFields(log.Fields{
		"query": query,
	}).Debug("starting query string processing")
	// Strip the surrounding double quotes form the query value\
	query = strings.Trim(query, "\"")
	// Divide up a *fieldName:fieldValue OR fieldName:fieldValue* query into separate statements
	statements := strings.Split(query, " OR ")
	for _, statement := range statements {
		// Spit a *fieldName:fieldValue* statement into separate pieces
		statementPieces := strings.Split(statement, ":")
		// Add the fieldValue piece to the list of valid values, removing the single quotes if the value is in quotes
		if len(statementPieces) > 1 {
			trimmedValue := strings.Trim(statementPieces[1], "'")
			if len(trimmedValue) > 0 {
				validValues[trimmedValue] = true
			}
		}
	}
	log.WithFields(log.Fields{
		"values.count": len(validValues),
		"values":       validValues,
	}).Debug("query string processed")
	return validValues
}

func (contControl *Contacts) Add(ctx *gin.Context) {
	var newContact models.Contact
	if err := ctx.BindJSON(&newContact); err != nil {
		respMessage := ErrorResp{
			Description: "unable to process new contact",
			Errors: []ErrorDetails{
				{
					Field:   "",
					Message: "bind failed",
					Code:    "new_contact_failure",
				},
			},
		}
		ctx.JSON(http.StatusInternalServerError, respMessage)
		return
	}

	invalidFields, isValid, err := newContact.IsValid(contControl.CurrentContacts)
	if !isValid {
		var respErrorDetails []ErrorDetails
		for _, field := range invalidFields {
			fieldError := ErrorDetails{
				Field:   field,
				Message: err.Error(),
				Code:    "invalid_field_value",
			}
			respErrorDetails = append(respErrorDetails, fieldError)
		}
		respMessage := ErrorResp{
			Description: err.Error(),
			Errors:      respErrorDetails,
		}
		ctx.JSON(http.StatusBadRequest, respMessage)
		return
	}

	// Generate a new, unique numeric ID for the contact and add it to the set of contacts
	var newContactID int
	for {
		newContactID = 100000000000 + rand.Intn(900000000000)
		if _, exists := contControl.CurrentContacts[newContactID]; !exists {
			// Format the current UTC time in the Frontdesk compatible string of "YYYY-MM-DDTHH:MM:SSZ"
			nowStr := time.Now().UTC().Format("2006-02-01T15:04:05Z")
			newContact.ID = newContactID
			newContact.CreatedAt = nowStr
			newContact.UpdatedAt = nowStr
			contControl.CurrentContacts[newContactID] = newContact
			break
		}
	}
	ctx.JSON(http.StatusCreated, contControl.CurrentContacts[newContactID])
}

func (contControl *Contacts) Update(ctx *gin.Context) {
	intID, err := getIntID(ctx)
	if err != nil {
		return
	}

	var updatedContact models.Contact
	if err := ctx.BindJSON(&updatedContact); err != nil {
		respMessage := ErrorResp{
			Description: "unable to process new contact",
			Errors: []ErrorDetails{
				{
					Field:   "",
					Message: fmt.Sprintf("bind failed: %s", err.Error()),
					Code:    "new_contact_failure",
				},
			},
		}
		ctx.JSON(http.StatusInternalServerError, respMessage)
		return
	}

	contactUpdated := false
	finalContact := contControl.CurrentContacts[intID]

	if updatedContact.Address != "" {
		finalContact.Address = updatedContact.Address
		contactUpdated = true
	}
	if updatedContact.Avatar != (models.ContactAvatar{}) {
		finalContact.Avatar = updatedContact.Avatar
		contactUpdated = true
	}
	if updatedContact.CompanyID != 0 {
		finalContact.CompanyID = updatedContact.CompanyID
		contactUpdated = true
	}
	if updatedContact.CustomFields != (models.ContactCustomFields{}) {
		finalContact.CustomFields = updatedContact.CustomFields
		contactUpdated = true
	}
	if updatedContact.Description != "" {
		finalContact.Description = updatedContact.Description
		contactUpdated = true
	}
	if updatedContact.Email != "" {
		finalContact.Email = updatedContact.Email
		contactUpdated = true
	}
	if updatedContact.JobTitle != "" {
		finalContact.JobTitle = updatedContact.JobTitle
		contactUpdated = true
	}
	if updatedContact.Language != "" {
		finalContact.Language = updatedContact.Language
		contactUpdated = true
	}
	if updatedContact.Mobile != "" {
		finalContact.Mobile = updatedContact.Mobile
		contactUpdated = true
	}
	if updatedContact.Name != "" {
		finalContact.Name = updatedContact.Name
		contactUpdated = true
	}
	if len(updatedContact.OtherCompanies) > 0 {
		finalContact.OtherCompanies = updatedContact.OtherCompanies
		contactUpdated = true
	}
	if len(updatedContact.OtherEmails) > 0 {
		finalContact.OtherEmails = updatedContact.OtherEmails
		contactUpdated = true
	}
	if updatedContact.PeopleID != "" {
		finalContact.PeopleID = updatedContact.PeopleID
		contactUpdated = true
	}
	if updatedContact.Phone != "" {
		finalContact.Phone = updatedContact.Phone
		contactUpdated = true
	}
	if len(updatedContact.Tags) > 0 {
		finalContact.Tags = updatedContact.Tags
		contactUpdated = true
	}
	if updatedContact.TimeZone != "" {
		finalContact.TimeZone = updatedContact.TimeZone
		contactUpdated = true
	}
	if updatedContact.TwitterID != "" {
		finalContact.TwitterID = updatedContact.TwitterID
		contactUpdated = true
	}
	if updatedContact.ViewAllTickets != nil {
		finalContact.ViewAllTickets = updatedContact.ViewAllTickets
		contactUpdated = true
	}

	invalidFields, isValid, err := finalContact.IsValid(contControl.CurrentContacts)
	if !isValid {
		var respErrorDetails []ErrorDetails
		for _, field := range invalidFields {
			fieldError := ErrorDetails{
				Field:   field,
				Message: err.Error(),
				Code:    "invalid_field_value",
			}
			respErrorDetails = append(respErrorDetails, fieldError)
		}
		respMessage := ErrorResp{
			Description: err.Error(),
			Errors:      respErrorDetails,
		}
		ctx.JSON(http.StatusBadRequest, respMessage)
		return
	}

	if contactUpdated {
		// Format the current UTC time in the Frontdesk compatible string of "YYYY-MM-DDTHH:MM:SSZ"
		nowStr := time.Now().UTC().Format("2006-02-01T15:04:05Z")
		finalContact.UpdatedAt = nowStr
	}
	contControl.CurrentContacts[intID] = finalContact
	ctx.JSON(http.StatusOK, contControl.CurrentContacts[intID])
}

func (contControl *Contacts) Delete(ctx *gin.Context) {
	intID, err := getIntID(ctx)
	if err != nil {
		return
	}
	delete(contControl.CurrentContacts, intID)
	ctx.JSON(http.StatusNoContent, nil)
}

func getIntID(ctx *gin.Context) (int, error) {
	ID := ctx.Param(ParamNameContactID)
	intID, err := strconv.Atoi(ID)
	if err != nil {
		respMessage := ErrorResp{
			Description: "invalid contact id specified",
			Errors: []ErrorDetails{
				{
					Field:   "id",
					Message: "id is not an integer",
					Code:    "invalid_id",
				},
			},
		}
		ctx.JSON(http.StatusBadRequest, respMessage)
		return 0, err
	}
	return intID, nil
}
