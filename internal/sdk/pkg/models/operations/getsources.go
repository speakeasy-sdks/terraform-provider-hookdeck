// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hashicups/internal/sdk/pkg/models/shared"
	"net/http"
	"time"
)

// GetSourcesArchivedAt2 - Date the source was archived
type GetSourcesArchivedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

type GetSourcesArchivedAtType string

const (
	GetSourcesArchivedAtTypeDateTime              GetSourcesArchivedAtType = "date-time"
	GetSourcesArchivedAtTypeGetSourcesArchivedAt2 GetSourcesArchivedAtType = "getSourcesArchivedAt_2"
)

type GetSourcesArchivedAt struct {
	DateTime              *time.Time
	GetSourcesArchivedAt2 *GetSourcesArchivedAt2

	Type GetSourcesArchivedAtType
}

func CreateGetSourcesArchivedAtDateTime(dateTime time.Time) GetSourcesArchivedAt {
	typ := GetSourcesArchivedAtTypeDateTime

	return GetSourcesArchivedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetSourcesArchivedAtGetSourcesArchivedAt2(getSourcesArchivedAt2 GetSourcesArchivedAt2) GetSourcesArchivedAt {
	typ := GetSourcesArchivedAtTypeGetSourcesArchivedAt2

	return GetSourcesArchivedAt{
		GetSourcesArchivedAt2: &getSourcesArchivedAt2,
		Type:                  typ,
	}
}

func (u *GetSourcesArchivedAt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = GetSourcesArchivedAtTypeDateTime
		return nil
	}

	getSourcesArchivedAt2 := new(GetSourcesArchivedAt2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getSourcesArchivedAt2); err == nil {
		u.GetSourcesArchivedAt2 = getSourcesArchivedAt2
		u.Type = GetSourcesArchivedAtTypeGetSourcesArchivedAt2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesArchivedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.GetSourcesArchivedAt2 != nil {
		return json.Marshal(u.GetSourcesArchivedAt2)
	}

	return nil, nil
}

type GetSourcesDir2 string

const (
	GetSourcesDir2Asc  GetSourcesDir2 = "asc"
	GetSourcesDir2Desc GetSourcesDir2 = "desc"
)

func (e GetSourcesDir2) ToPointer() *GetSourcesDir2 {
	return &e
}

func (e *GetSourcesDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetSourcesDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSourcesDir2: %v", v)
	}
}

// GetSourcesDir1 - Sort direction
type GetSourcesDir1 string

const (
	GetSourcesDir1Asc  GetSourcesDir1 = "asc"
	GetSourcesDir1Desc GetSourcesDir1 = "desc"
)

func (e GetSourcesDir1) ToPointer() *GetSourcesDir1 {
	return &e
}

func (e *GetSourcesDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetSourcesDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSourcesDir1: %v", v)
	}
}

type GetSourcesDirType string

const (
	GetSourcesDirTypeGetSourcesDir1        GetSourcesDirType = "getSourcesDir_1"
	GetSourcesDirTypeArrayOfgetSourcesDir2 GetSourcesDirType = "arrayOfgetSourcesDir_2"
)

type GetSourcesDir struct {
	GetSourcesDir1        *GetSourcesDir1
	ArrayOfgetSourcesDir2 []GetSourcesDir2

	Type GetSourcesDirType
}

func CreateGetSourcesDirGetSourcesDir1(getSourcesDir1 GetSourcesDir1) GetSourcesDir {
	typ := GetSourcesDirTypeGetSourcesDir1

	return GetSourcesDir{
		GetSourcesDir1: &getSourcesDir1,
		Type:           typ,
	}
}

func CreateGetSourcesDirArrayOfgetSourcesDir2(arrayOfgetSourcesDir2 []GetSourcesDir2) GetSourcesDir {
	typ := GetSourcesDirTypeArrayOfgetSourcesDir2

	return GetSourcesDir{
		ArrayOfgetSourcesDir2: arrayOfgetSourcesDir2,
		Type:                  typ,
	}
}

func (u *GetSourcesDir) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	getSourcesDir1 := new(GetSourcesDir1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getSourcesDir1); err == nil {
		u.GetSourcesDir1 = getSourcesDir1
		u.Type = GetSourcesDirTypeGetSourcesDir1
		return nil
	}

	arrayOfgetSourcesDir2 := []GetSourcesDir2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetSourcesDir2); err == nil {
		u.ArrayOfgetSourcesDir2 = arrayOfgetSourcesDir2
		u.Type = GetSourcesDirTypeArrayOfgetSourcesDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesDir) MarshalJSON() ([]byte, error) {
	if u.GetSourcesDir1 != nil {
		return json.Marshal(u.GetSourcesDir1)
	}

	if u.ArrayOfgetSourcesDir2 != nil {
		return json.Marshal(u.ArrayOfgetSourcesDir2)
	}

	return nil, nil
}

type GetSourcesIDType string

const (
	GetSourcesIDTypeStr        GetSourcesIDType = "str"
	GetSourcesIDTypeArrayOfstr GetSourcesIDType = "arrayOfstr"
)

type GetSourcesID struct {
	Str        *string
	ArrayOfstr []string

	Type GetSourcesIDType
}

func CreateGetSourcesIDStr(str string) GetSourcesID {
	typ := GetSourcesIDTypeStr

	return GetSourcesID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetSourcesIDArrayOfstr(arrayOfstr []string) GetSourcesID {
	typ := GetSourcesIDTypeArrayOfstr

	return GetSourcesID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetSourcesID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetSourcesIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetSourcesIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

// GetSourcesIntegrationID2 - Filter by integration IDs
type GetSourcesIntegrationID2 struct {
	Any *bool `queryParam:"name=any"`
}

type GetSourcesIntegrationIDType string

const (
	GetSourcesIntegrationIDTypeStr                      GetSourcesIntegrationIDType = "str"
	GetSourcesIntegrationIDTypeGetSourcesIntegrationID2 GetSourcesIntegrationIDType = "getSourcesIntegrationID_2"
)

type GetSourcesIntegrationID struct {
	Str                      *string
	GetSourcesIntegrationID2 *GetSourcesIntegrationID2

	Type GetSourcesIntegrationIDType
}

func CreateGetSourcesIntegrationIDStr(str string) GetSourcesIntegrationID {
	typ := GetSourcesIntegrationIDTypeStr

	return GetSourcesIntegrationID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetSourcesIntegrationIDGetSourcesIntegrationID2(getSourcesIntegrationID2 GetSourcesIntegrationID2) GetSourcesIntegrationID {
	typ := GetSourcesIntegrationIDTypeGetSourcesIntegrationID2

	return GetSourcesIntegrationID{
		GetSourcesIntegrationID2: &getSourcesIntegrationID2,
		Type:                     typ,
	}
}

func (u *GetSourcesIntegrationID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetSourcesIntegrationIDTypeStr
		return nil
	}

	getSourcesIntegrationID2 := new(GetSourcesIntegrationID2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getSourcesIntegrationID2); err == nil {
		u.GetSourcesIntegrationID2 = getSourcesIntegrationID2
		u.Type = GetSourcesIntegrationIDTypeGetSourcesIntegrationID2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesIntegrationID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.GetSourcesIntegrationID2 != nil {
		return json.Marshal(u.GetSourcesIntegrationID2)
	}

	return nil, nil
}

// GetSourcesName2 - The source name
type GetSourcesName2 struct {
	Any      *bool   `queryParam:"name=any"`
	Contains *string `queryParam:"name=contains"`
	Gt       *string `queryParam:"name=gt"`
	Gte      *string `queryParam:"name=gte"`
	Le       *string `queryParam:"name=le"`
	Lte      *string `queryParam:"name=lte"`
}

type GetSourcesNameType string

const (
	GetSourcesNameTypeStr             GetSourcesNameType = "str"
	GetSourcesNameTypeGetSourcesName2 GetSourcesNameType = "getSourcesName_2"
)

type GetSourcesName struct {
	Str             *string
	GetSourcesName2 *GetSourcesName2

	Type GetSourcesNameType
}

func CreateGetSourcesNameStr(str string) GetSourcesName {
	typ := GetSourcesNameTypeStr

	return GetSourcesName{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetSourcesNameGetSourcesName2(getSourcesName2 GetSourcesName2) GetSourcesName {
	typ := GetSourcesNameTypeGetSourcesName2

	return GetSourcesName{
		GetSourcesName2: &getSourcesName2,
		Type:            typ,
	}
}

func (u *GetSourcesName) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetSourcesNameTypeStr
		return nil
	}

	getSourcesName2 := new(GetSourcesName2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getSourcesName2); err == nil {
		u.GetSourcesName2 = getSourcesName2
		u.Type = GetSourcesNameTypeGetSourcesName2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesName) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.GetSourcesName2 != nil {
		return json.Marshal(u.GetSourcesName2)
	}

	return nil, nil
}

type GetSourcesOrderBy2 string

const (
	GetSourcesOrderBy2CreatedAt GetSourcesOrderBy2 = "created_at"
)

func (e GetSourcesOrderBy2) ToPointer() *GetSourcesOrderBy2 {
	return &e
}

func (e *GetSourcesOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetSourcesOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSourcesOrderBy2: %v", v)
	}
}

// GetSourcesOrderBy1 - Sort key(s)
type GetSourcesOrderBy1 string

const (
	GetSourcesOrderBy1CreatedAt GetSourcesOrderBy1 = "created_at"
)

func (e GetSourcesOrderBy1) ToPointer() *GetSourcesOrderBy1 {
	return &e
}

func (e *GetSourcesOrderBy1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetSourcesOrderBy1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSourcesOrderBy1: %v", v)
	}
}

type GetSourcesOrderByType string

const (
	GetSourcesOrderByTypeGetSourcesOrderBy1        GetSourcesOrderByType = "getSourcesOrderBy_1"
	GetSourcesOrderByTypeArrayOfgetSourcesOrderBy2 GetSourcesOrderByType = "arrayOfgetSourcesOrderBy_2"
)

type GetSourcesOrderBy struct {
	GetSourcesOrderBy1        *GetSourcesOrderBy1
	ArrayOfgetSourcesOrderBy2 []GetSourcesOrderBy2

	Type GetSourcesOrderByType
}

func CreateGetSourcesOrderByGetSourcesOrderBy1(getSourcesOrderBy1 GetSourcesOrderBy1) GetSourcesOrderBy {
	typ := GetSourcesOrderByTypeGetSourcesOrderBy1

	return GetSourcesOrderBy{
		GetSourcesOrderBy1: &getSourcesOrderBy1,
		Type:               typ,
	}
}

func CreateGetSourcesOrderByArrayOfgetSourcesOrderBy2(arrayOfgetSourcesOrderBy2 []GetSourcesOrderBy2) GetSourcesOrderBy {
	typ := GetSourcesOrderByTypeArrayOfgetSourcesOrderBy2

	return GetSourcesOrderBy{
		ArrayOfgetSourcesOrderBy2: arrayOfgetSourcesOrderBy2,
		Type:                      typ,
	}
}

func (u *GetSourcesOrderBy) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	getSourcesOrderBy1 := new(GetSourcesOrderBy1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getSourcesOrderBy1); err == nil {
		u.GetSourcesOrderBy1 = getSourcesOrderBy1
		u.Type = GetSourcesOrderByTypeGetSourcesOrderBy1
		return nil
	}

	arrayOfgetSourcesOrderBy2 := []GetSourcesOrderBy2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetSourcesOrderBy2); err == nil {
		u.ArrayOfgetSourcesOrderBy2 = arrayOfgetSourcesOrderBy2
		u.Type = GetSourcesOrderByTypeArrayOfgetSourcesOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetSourcesOrderBy1 != nil {
		return json.Marshal(u.GetSourcesOrderBy1)
	}

	if u.ArrayOfgetSourcesOrderBy2 != nil {
		return json.Marshal(u.ArrayOfgetSourcesOrderBy2)
	}

	return nil, nil
}

type GetSourcesRequest struct {
	Archived      *bool                    `queryParam:"style=form,explode=true,name=archived"`
	ArchivedAt    *GetSourcesArchivedAt    `queryParam:"style=form,explode=true,name=archived_at"`
	Dir           *GetSourcesDir           `queryParam:"style=form,explode=true,name=dir"`
	ID            *GetSourcesID            `queryParam:"style=form,explode=true,name=id"`
	IntegrationID *GetSourcesIntegrationID `queryParam:"style=form,explode=true,name=integration_id"`
	Limit         *int64                   `queryParam:"style=form,explode=true,name=limit"`
	Name          *GetSourcesName          `queryParam:"style=form,explode=true,name=name"`
	Next          *string                  `queryParam:"style=form,explode=true,name=next"`
	OrderBy       *GetSourcesOrderBy       `queryParam:"style=form,explode=true,name=order_by"`
	Prev          *string                  `queryParam:"style=form,explode=true,name=prev"`
}

type GetSourcesResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// List of sources
	SourcePaginatedResult *shared.SourcePaginatedResult
	StatusCode            int
	RawResponse           *http.Response
}
