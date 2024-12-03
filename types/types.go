package types

import (
	"time"
)

type MstrAuthenticationMode int

const (
	Standard  MstrAuthenticationMode = 1
	Anonymous MstrAuthenticationMode = 8
)

type MstrAuthentication interface {
	AuthenticationMode() MstrAuthenticationMode
	AuthenticationRequest(applicationType *MstrApplicationType) *AuthenticationRequest
}

type AnonymousAuthentication struct {
}

func (a *AnonymousAuthentication) AuthenticationMode() MstrAuthenticationMode {
	return Anonymous
}

func (a *AnonymousAuthentication) AuthenticationRequest(applicationType *MstrApplicationType) *AuthenticationRequest {
	req := AuthenticationRequest{
		AuthenticationMode: Anonymous,
	}
	if applicationType != nil {
		req.ApplicationType = applicationType
	}
	return &req
}

type StandardAuthentication struct {
	Username string
	Password string
}

func (a *StandardAuthentication) AuthenticationMode() MstrAuthenticationMode {
	return Standard
}

func (a *StandardAuthentication) AuthenticationRequest(applicationType *MstrApplicationType) *AuthenticationRequest {
	req := AuthenticationRequest{
		Username:           &a.Username,
		Password:           &a.Password,
		AuthenticationMode: Standard,
	}
	if applicationType != nil {
		req.ApplicationType = applicationType
	}
	return &req
}

type MstrRestTimestamp time.Time

func (t MstrRestTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format("\"2006-01-02T15:04:05.000-0700\"")), nil
}

func (t *MstrRestTimestamp) UnmarshalJSON(data []byte) error {
	tm, err := time.Parse("\"2006-01-02T15:04:05.000-0700\"", string(data))
	if err != nil {
		return err
	}
	*t = MstrRestTimestamp(tm)
	return nil
}

type MstrRestMultiStatusResponse struct {
	Results []MstrRestStatusResponse `json:"results"`
}

type MstrRestStatusResponse struct {
	ID     string      `json:"id"`
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

type MstrObject struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Description  *string            `json:"description,omitempty"`
	DateCreated  *MstrRestTimestamp `json:"dateCreated,omitempty"`
	DateModified *MstrRestTimestamp `json:"dateModified,omitempty"`
}

type MstrOwner struct {
	MstrObject
	Expired bool `json:"expired"`
}

type ClientConfig struct {
	ApplicationType MstrApplicationType
}

type APIRequestInput struct {
	Method       string
	APIPath      string
	QueryParams  *map[string]string
	Body         interface{}
	ResponseJSON interface{}
}
