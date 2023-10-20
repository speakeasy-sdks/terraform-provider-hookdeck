// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"hashicups/internal/sdk/pkg/models/shared"
	"hashicups/internal/sdk/pkg/utils"
	"net/http"
	"time"
)

// GetRulesetsArchivedAt2 - Date the ruleset was archived
type GetRulesetsArchivedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetRulesetsArchivedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRulesetsArchivedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetRulesetsArchivedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetRulesetsArchivedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetRulesetsArchivedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetRulesetsArchivedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetRulesetsArchivedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetRulesetsArchivedAtType string

const (
	GetRulesetsArchivedAtTypeDateTime               GetRulesetsArchivedAtType = "date-time"
	GetRulesetsArchivedAtTypeGetRulesetsArchivedAt2 GetRulesetsArchivedAtType = "getRulesetsArchivedAt_2"
)

type GetRulesetsArchivedAt struct {
	DateTime               *time.Time
	GetRulesetsArchivedAt2 *GetRulesetsArchivedAt2

	Type GetRulesetsArchivedAtType
}

func CreateGetRulesetsArchivedAtDateTime(dateTime time.Time) GetRulesetsArchivedAt {
	typ := GetRulesetsArchivedAtTypeDateTime

	return GetRulesetsArchivedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetRulesetsArchivedAtGetRulesetsArchivedAt2(getRulesetsArchivedAt2 GetRulesetsArchivedAt2) GetRulesetsArchivedAt {
	typ := GetRulesetsArchivedAtTypeGetRulesetsArchivedAt2

	return GetRulesetsArchivedAt{
		GetRulesetsArchivedAt2: &getRulesetsArchivedAt2,
		Type:                   typ,
	}
}

func (u *GetRulesetsArchivedAt) UnmarshalJSON(data []byte) error {

	getRulesetsArchivedAt2 := new(GetRulesetsArchivedAt2)
	if err := utils.UnmarshalJSON(data, &getRulesetsArchivedAt2, "", true, true); err == nil {
		u.GetRulesetsArchivedAt2 = getRulesetsArchivedAt2
		u.Type = GetRulesetsArchivedAtTypeGetRulesetsArchivedAt2
		return nil
	}

	dateTime := new(time.Time)
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = dateTime
		u.Type = GetRulesetsArchivedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsArchivedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetRulesetsArchivedAt2 != nil {
		return utils.MarshalJSON(u.GetRulesetsArchivedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsDir2 string

const (
	GetRulesetsDir2Asc  GetRulesetsDir2 = "asc"
	GetRulesetsDir2Desc GetRulesetsDir2 = "desc"
)

func (e GetRulesetsDir2) ToPointer() *GetRulesetsDir2 {
	return &e
}

func (e *GetRulesetsDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRulesetsDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsDir2: %v", v)
	}
}

// GetRulesetsDir1 - Sort direction
type GetRulesetsDir1 string

const (
	GetRulesetsDir1Asc  GetRulesetsDir1 = "asc"
	GetRulesetsDir1Desc GetRulesetsDir1 = "desc"
)

func (e GetRulesetsDir1) ToPointer() *GetRulesetsDir1 {
	return &e
}

func (e *GetRulesetsDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRulesetsDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsDir1: %v", v)
	}
}

type GetRulesetsDirType string

const (
	GetRulesetsDirTypeGetRulesetsDir1        GetRulesetsDirType = "getRulesetsDir_1"
	GetRulesetsDirTypeArrayOfgetRulesetsDir2 GetRulesetsDirType = "arrayOfgetRulesetsDir_2"
)

type GetRulesetsDir struct {
	GetRulesetsDir1        *GetRulesetsDir1
	ArrayOfgetRulesetsDir2 []GetRulesetsDir2

	Type GetRulesetsDirType
}

func CreateGetRulesetsDirGetRulesetsDir1(getRulesetsDir1 GetRulesetsDir1) GetRulesetsDir {
	typ := GetRulesetsDirTypeGetRulesetsDir1

	return GetRulesetsDir{
		GetRulesetsDir1: &getRulesetsDir1,
		Type:            typ,
	}
}

func CreateGetRulesetsDirArrayOfgetRulesetsDir2(arrayOfgetRulesetsDir2 []GetRulesetsDir2) GetRulesetsDir {
	typ := GetRulesetsDirTypeArrayOfgetRulesetsDir2

	return GetRulesetsDir{
		ArrayOfgetRulesetsDir2: arrayOfgetRulesetsDir2,
		Type:                   typ,
	}
}

func (u *GetRulesetsDir) UnmarshalJSON(data []byte) error {

	getRulesetsDir1 := new(GetRulesetsDir1)
	if err := utils.UnmarshalJSON(data, &getRulesetsDir1, "", true, true); err == nil {
		u.GetRulesetsDir1 = getRulesetsDir1
		u.Type = GetRulesetsDirTypeGetRulesetsDir1
		return nil
	}

	arrayOfgetRulesetsDir2 := []GetRulesetsDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetRulesetsDir2, "", true, true); err == nil {
		u.ArrayOfgetRulesetsDir2 = arrayOfgetRulesetsDir2
		u.Type = GetRulesetsDirTypeArrayOfgetRulesetsDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsDir) MarshalJSON() ([]byte, error) {
	if u.GetRulesetsDir1 != nil {
		return utils.MarshalJSON(u.GetRulesetsDir1, "", true)
	}

	if u.ArrayOfgetRulesetsDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetRulesetsDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsIDType string

const (
	GetRulesetsIDTypeStr        GetRulesetsIDType = "str"
	GetRulesetsIDTypeArrayOfstr GetRulesetsIDType = "arrayOfstr"
)

type GetRulesetsID struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsIDType
}

func CreateGetRulesetsIDStr(str string) GetRulesetsID {
	typ := GetRulesetsIDTypeStr

	return GetRulesetsID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsIDArrayOfstr(arrayOfstr []string) GetRulesetsID {
	typ := GetRulesetsIDTypeArrayOfstr

	return GetRulesetsID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsName2ContainsType string

const (
	GetRulesetsName2ContainsTypeStr        GetRulesetsName2ContainsType = "str"
	GetRulesetsName2ContainsTypeArrayOfstr GetRulesetsName2ContainsType = "arrayOfstr"
)

type GetRulesetsName2Contains struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsName2ContainsType
}

func CreateGetRulesetsName2ContainsStr(str string) GetRulesetsName2Contains {
	typ := GetRulesetsName2ContainsTypeStr

	return GetRulesetsName2Contains{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsName2ContainsArrayOfstr(arrayOfstr []string) GetRulesetsName2Contains {
	typ := GetRulesetsName2ContainsTypeArrayOfstr

	return GetRulesetsName2Contains{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsName2Contains) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsName2ContainsTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsName2ContainsTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsName2Contains) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsName2GtType string

const (
	GetRulesetsName2GtTypeStr        GetRulesetsName2GtType = "str"
	GetRulesetsName2GtTypeArrayOfstr GetRulesetsName2GtType = "arrayOfstr"
)

type GetRulesetsName2Gt struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsName2GtType
}

func CreateGetRulesetsName2GtStr(str string) GetRulesetsName2Gt {
	typ := GetRulesetsName2GtTypeStr

	return GetRulesetsName2Gt{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsName2GtArrayOfstr(arrayOfstr []string) GetRulesetsName2Gt {
	typ := GetRulesetsName2GtTypeArrayOfstr

	return GetRulesetsName2Gt{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsName2Gt) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsName2GtTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsName2GtTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsName2Gt) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsName2GteType string

const (
	GetRulesetsName2GteTypeStr        GetRulesetsName2GteType = "str"
	GetRulesetsName2GteTypeArrayOfstr GetRulesetsName2GteType = "arrayOfstr"
)

type GetRulesetsName2Gte struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsName2GteType
}

func CreateGetRulesetsName2GteStr(str string) GetRulesetsName2Gte {
	typ := GetRulesetsName2GteTypeStr

	return GetRulesetsName2Gte{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsName2GteArrayOfstr(arrayOfstr []string) GetRulesetsName2Gte {
	typ := GetRulesetsName2GteTypeArrayOfstr

	return GetRulesetsName2Gte{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsName2Gte) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsName2GteTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsName2GteTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsName2Gte) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsName2LeType string

const (
	GetRulesetsName2LeTypeStr        GetRulesetsName2LeType = "str"
	GetRulesetsName2LeTypeArrayOfstr GetRulesetsName2LeType = "arrayOfstr"
)

type GetRulesetsName2Le struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsName2LeType
}

func CreateGetRulesetsName2LeStr(str string) GetRulesetsName2Le {
	typ := GetRulesetsName2LeTypeStr

	return GetRulesetsName2Le{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsName2LeArrayOfstr(arrayOfstr []string) GetRulesetsName2Le {
	typ := GetRulesetsName2LeTypeArrayOfstr

	return GetRulesetsName2Le{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsName2Le) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsName2LeTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsName2LeTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsName2Le) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsName2LteType string

const (
	GetRulesetsName2LteTypeStr        GetRulesetsName2LteType = "str"
	GetRulesetsName2LteTypeArrayOfstr GetRulesetsName2LteType = "arrayOfstr"
)

type GetRulesetsName2Lte struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsName2LteType
}

func CreateGetRulesetsName2LteStr(str string) GetRulesetsName2Lte {
	typ := GetRulesetsName2LteTypeStr

	return GetRulesetsName2Lte{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsName2LteArrayOfstr(arrayOfstr []string) GetRulesetsName2Lte {
	typ := GetRulesetsName2LteTypeArrayOfstr

	return GetRulesetsName2Lte{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsName2Lte) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsName2LteTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsName2LteTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsName2Lte) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetRulesetsName2 - The ruleset name
type GetRulesetsName2 struct {
	Any      *bool                     `queryParam:"name=any"`
	Contains *GetRulesetsName2Contains `queryParam:"name=contains"`
	Gt       *GetRulesetsName2Gt       `queryParam:"name=gt"`
	Gte      *GetRulesetsName2Gte      `queryParam:"name=gte"`
	Le       *GetRulesetsName2Le       `queryParam:"name=le"`
	Lte      *GetRulesetsName2Lte      `queryParam:"name=lte"`
}

func (o *GetRulesetsName2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetRulesetsName2) GetContains() *GetRulesetsName2Contains {
	if o == nil {
		return nil
	}
	return o.Contains
}

func (o *GetRulesetsName2) GetGt() *GetRulesetsName2Gt {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetRulesetsName2) GetGte() *GetRulesetsName2Gte {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetRulesetsName2) GetLe() *GetRulesetsName2Le {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetRulesetsName2) GetLte() *GetRulesetsName2Lte {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetRulesetsName1Type string

const (
	GetRulesetsName1TypeStr        GetRulesetsName1Type = "str"
	GetRulesetsName1TypeArrayOfstr GetRulesetsName1Type = "arrayOfstr"
)

type GetRulesetsName1 struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsName1Type
}

func CreateGetRulesetsName1Str(str string) GetRulesetsName1 {
	typ := GetRulesetsName1TypeStr

	return GetRulesetsName1{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsName1ArrayOfstr(arrayOfstr []string) GetRulesetsName1 {
	typ := GetRulesetsName1TypeArrayOfstr

	return GetRulesetsName1{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsName1) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsName1TypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsName1TypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsName1) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsNameType string

const (
	GetRulesetsNameTypeGetRulesetsName1 GetRulesetsNameType = "getRulesetsName_1"
	GetRulesetsNameTypeGetRulesetsName2 GetRulesetsNameType = "getRulesetsName_2"
)

type GetRulesetsName struct {
	GetRulesetsName1 *GetRulesetsName1
	GetRulesetsName2 *GetRulesetsName2

	Type GetRulesetsNameType
}

func CreateGetRulesetsNameGetRulesetsName1(getRulesetsName1 GetRulesetsName1) GetRulesetsName {
	typ := GetRulesetsNameTypeGetRulesetsName1

	return GetRulesetsName{
		GetRulesetsName1: &getRulesetsName1,
		Type:             typ,
	}
}

func CreateGetRulesetsNameGetRulesetsName2(getRulesetsName2 GetRulesetsName2) GetRulesetsName {
	typ := GetRulesetsNameTypeGetRulesetsName2

	return GetRulesetsName{
		GetRulesetsName2: &getRulesetsName2,
		Type:             typ,
	}
}

func (u *GetRulesetsName) UnmarshalJSON(data []byte) error {

	getRulesetsName2 := new(GetRulesetsName2)
	if err := utils.UnmarshalJSON(data, &getRulesetsName2, "", true, true); err == nil {
		u.GetRulesetsName2 = getRulesetsName2
		u.Type = GetRulesetsNameTypeGetRulesetsName2
		return nil
	}

	getRulesetsName1 := new(GetRulesetsName1)
	if err := utils.UnmarshalJSON(data, &getRulesetsName1, "", true, true); err == nil {
		u.GetRulesetsName1 = getRulesetsName1
		u.Type = GetRulesetsNameTypeGetRulesetsName1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsName) MarshalJSON() ([]byte, error) {
	if u.GetRulesetsName1 != nil {
		return utils.MarshalJSON(u.GetRulesetsName1, "", true)
	}

	if u.GetRulesetsName2 != nil {
		return utils.MarshalJSON(u.GetRulesetsName2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsOrderBy2 string

const (
	GetRulesetsOrderBy2CreatedAt GetRulesetsOrderBy2 = "created_at"
)

func (e GetRulesetsOrderBy2) ToPointer() *GetRulesetsOrderBy2 {
	return &e
}

func (e *GetRulesetsOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetRulesetsOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsOrderBy2: %v", v)
	}
}

// GetRulesetsOrderBy1 - Sort key(s)
type GetRulesetsOrderBy1 string

const (
	GetRulesetsOrderBy1CreatedAt GetRulesetsOrderBy1 = "created_at"
)

func (e GetRulesetsOrderBy1) ToPointer() *GetRulesetsOrderBy1 {
	return &e
}

func (e *GetRulesetsOrderBy1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetRulesetsOrderBy1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsOrderBy1: %v", v)
	}
}

type GetRulesetsOrderByType string

const (
	GetRulesetsOrderByTypeGetRulesetsOrderBy1        GetRulesetsOrderByType = "getRulesetsOrderBy_1"
	GetRulesetsOrderByTypeArrayOfgetRulesetsOrderBy2 GetRulesetsOrderByType = "arrayOfgetRulesetsOrderBy_2"
)

type GetRulesetsOrderBy struct {
	GetRulesetsOrderBy1        *GetRulesetsOrderBy1
	ArrayOfgetRulesetsOrderBy2 []GetRulesetsOrderBy2

	Type GetRulesetsOrderByType
}

func CreateGetRulesetsOrderByGetRulesetsOrderBy1(getRulesetsOrderBy1 GetRulesetsOrderBy1) GetRulesetsOrderBy {
	typ := GetRulesetsOrderByTypeGetRulesetsOrderBy1

	return GetRulesetsOrderBy{
		GetRulesetsOrderBy1: &getRulesetsOrderBy1,
		Type:                typ,
	}
}

func CreateGetRulesetsOrderByArrayOfgetRulesetsOrderBy2(arrayOfgetRulesetsOrderBy2 []GetRulesetsOrderBy2) GetRulesetsOrderBy {
	typ := GetRulesetsOrderByTypeArrayOfgetRulesetsOrderBy2

	return GetRulesetsOrderBy{
		ArrayOfgetRulesetsOrderBy2: arrayOfgetRulesetsOrderBy2,
		Type:                       typ,
	}
}

func (u *GetRulesetsOrderBy) UnmarshalJSON(data []byte) error {

	getRulesetsOrderBy1 := new(GetRulesetsOrderBy1)
	if err := utils.UnmarshalJSON(data, &getRulesetsOrderBy1, "", true, true); err == nil {
		u.GetRulesetsOrderBy1 = getRulesetsOrderBy1
		u.Type = GetRulesetsOrderByTypeGetRulesetsOrderBy1
		return nil
	}

	arrayOfgetRulesetsOrderBy2 := []GetRulesetsOrderBy2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetRulesetsOrderBy2, "", true, true); err == nil {
		u.ArrayOfgetRulesetsOrderBy2 = arrayOfgetRulesetsOrderBy2
		u.Type = GetRulesetsOrderByTypeArrayOfgetRulesetsOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetRulesetsOrderBy1 != nil {
		return utils.MarshalJSON(u.GetRulesetsOrderBy1, "", true)
	}

	if u.ArrayOfgetRulesetsOrderBy2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetRulesetsOrderBy2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsRequest struct {
	Archived   *bool                  `queryParam:"style=form,explode=true,name=archived"`
	ArchivedAt *GetRulesetsArchivedAt `queryParam:"style=form,explode=true,name=archived_at"`
	Dir        *GetRulesetsDir        `queryParam:"style=form,explode=true,name=dir"`
	ID         *GetRulesetsID         `queryParam:"style=form,explode=true,name=id"`
	Limit      *int64                 `queryParam:"style=form,explode=true,name=limit"`
	Name       *GetRulesetsName       `queryParam:"style=form,explode=true,name=name"`
	Next       *string                `queryParam:"style=form,explode=true,name=next"`
	OrderBy    *GetRulesetsOrderBy    `queryParam:"style=form,explode=true,name=order_by"`
	Prev       *string                `queryParam:"style=form,explode=true,name=prev"`
}

func (o *GetRulesetsRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetRulesetsRequest) GetArchivedAt() *GetRulesetsArchivedAt {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *GetRulesetsRequest) GetDir() *GetRulesetsDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetRulesetsRequest) GetID() *GetRulesetsID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetRulesetsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetRulesetsRequest) GetName() *GetRulesetsName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetRulesetsRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetRulesetsRequest) GetOrderBy() *GetRulesetsOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetRulesetsRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type GetRulesetsResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	// HTTP response content type for this operation
	ContentType string
	// List of rulesets
	RulesetPaginatedResult *shared.RulesetPaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRulesetsResponse) GetAPIErrorResponse() *shared.APIErrorResponse {
	if o == nil {
		return nil
	}
	return o.APIErrorResponse
}

func (o *GetRulesetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRulesetsResponse) GetRulesetPaginatedResult() *shared.RulesetPaginatedResult {
	if o == nil {
		return nil
	}
	return o.RulesetPaginatedResult
}

func (o *GetRulesetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRulesetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
