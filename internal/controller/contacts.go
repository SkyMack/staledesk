package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/SkyMack/staledesk/config"
	"github.com/SkyMack/staledesk/internal/models"
	"github.com/gin-gonic/gin"
)

const (
	ParamNameContactID = "id"
)

var (
	ErrFieldAtLeastOneSet     = fmt.Errorf("at least one of these fields must be set")
	ErrFieldRequired          = fmt.Errorf("field is required")
	ErrFieldValueMustBeUnique = fmt.Errorf("field must have a unique value")
)

type Contacts struct {
	CurrentContacts map[int]models.Contact
}

func NewContactsController() *Contacts {
	return &Contacts{
		CurrentContacts: config.Config.Contacts,
	}
}

func (contControl *Contacts) GetAll(ctx *gin.Context) {
	respCode := http.StatusOK
	ctx.JSON(respCode, contControl.CurrentContacts)
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

	// Ensure the required fields are set
	if newContact.Name == "" {
		respMessage := ErrorResp{
			Description: ErrFieldRequired.Error(),
			Errors: []ErrorDetails{
				{
					Field:   "name",
					Message: ErrFieldRequired.Error(),
					Code:    "missing_field",
				},
			},
		}
		ctx.JSON(http.StatusBadRequest, respMessage)
		return
	}
	if newContact.Email == "" &&
		newContact.Phone == "" &&
		newContact.Mobile == "" &&
		newContact.TwitterID == "" &&
		newContact.ExternalID == "" {

		respMessage := ErrorResp{
			Description: ErrFieldAtLeastOneSet.Error(),
			Errors: []ErrorDetails{
				{
					Field:   "email",
					Message: ErrFieldAtLeastOneSet.Error(),
					Code:    "missing_field",
				},
				{
					Field:   "phone",
					Message: ErrFieldAtLeastOneSet.Error(),
					Code:    "missing_field",
				},
				{
					Field:   "mobile",
					Message: ErrFieldAtLeastOneSet.Error(),
					Code:    "missing_field",
				},
				{
					Field:   "twitter_id",
					Message: ErrFieldAtLeastOneSet.Error(),
					Code:    "missing_field",
				},
				{
					Field:   "unique_external_id",
					Message: ErrFieldAtLeastOneSet.Error(),
					Code:    "missing_field",
				},
			},
		}
		ctx.JSON(http.StatusBadRequest, respMessage)
		return
	}

	// Ensure fields that must be unique will remain so
	respUniqueFieldErr := ErrorResp{
		Description: ErrFieldValueMustBeUnique.Error(),
		Errors: []ErrorDetails{
			{
				Message: ErrFieldValueMustBeUnique.Error(),
				Code:    "existing_field_value",
			},
		},
	}
	for _, contact := range contControl.CurrentContacts {
		if contact.Email == newContact.Email {
			respUniqueFieldErr.Errors[0].Field = "email"
			ctx.JSON(http.StatusBadRequest, respUniqueFieldErr)
			return
		}
		if contact.TwitterID == newContact.TwitterID {
			respUniqueFieldErr.Errors[0].Field = "twitter_id"
			ctx.JSON(http.StatusBadRequest, respUniqueFieldErr)
			return
		}
		if contact.ExternalID == newContact.ExternalID {
			respUniqueFieldErr.Errors[0].Field = "unique_external_id"
			ctx.JSON(http.StatusBadRequest, respUniqueFieldErr)
			return
		}
	}

	// Generate a new, unique numeric ID for the contact and add it to the set of contacts
	var newContactID int
	for {
		newContactID = 100000000000 + rand.Intn(900000000000)
		if _, exists := contControl.CurrentContacts[newContactID]; !exists {
			contControl.CurrentContacts[newContactID] = newContact
			// TODO: Add created_at/updated_at values
			break
		}
	}
	ctx.JSON(http.StatusCreated, contControl.CurrentContacts[newContactID])
}

func (contControl *Contacts) Update(ctx *gin.Context) {

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
