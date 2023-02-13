package controller

import (
	"net/http"
	"strconv"

	"github.com/SkyMack/staledesk/config"
	"github.com/SkyMack/staledesk/internal/models"
	"github.com/gin-gonic/gin"
)

const (
	ParamNameContactID = "id"
)

type Contacts struct {
	CurrentContacts map[int]models.Contact
}

func NewContactsController() *Contacts {
	return &Contacts{
		CurrentContacts: config.Config.Contacts,
	}
}

func (contControl Contacts) GetAll(ctx *gin.Context) {
	respCode := http.StatusOK
	ctx.JSON(respCode, contControl.CurrentContacts)
}

func (contControl Contacts) GetByID(ctx *gin.Context) {
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
	}

	contact, exists := contControl.CurrentContacts[intID]
	if !exists {
		ctx.JSON(http.StatusNotFound, nil)
	} else {
		respCode := http.StatusOK
		ctx.JSON(respCode, contact)
	}
}

func (contControl Contacts) Search(ctx *gin.Context) {

}

func (contControl Contacts) Add(ctx *gin.Context) {

}

func (contControl Contacts) Update(ctx *gin.Context) {

}

func (contControl Contacts) Delete(ctx *gin.Context) {

}
