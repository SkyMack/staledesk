package controller

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResp struct {
	Description string         `json:"description"`
	Errors      []ErrorDetails `json:"errors"`
}

type ErrorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// getIntID fetches the ID from the URL path and converts it from a string to an integer
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

// randHexString generates a random hex value of length `n`
func randHexString(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
