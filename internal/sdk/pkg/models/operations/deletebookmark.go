// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteBookmarkRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteBookmarkResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// An object with deleted bookmark's id
	DeletedBookmarkResponse *shared.DeletedBookmarkResponse
	StatusCode              int
	RawResponse             *http.Response
}
