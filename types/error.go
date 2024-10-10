package types

import "fmt"

type MstrRestError struct {
	Code                   string             `json:"code"`
	IntelligenceServerCode int                `json:"iServerCode"`
	Message                string             `json:"message"`
	TicketID               string             `json:"ticketId"`
	SubErrors              []MstrRestSubError `json:"subErrors"`
}

type MstrRestSubError struct {
	IntelligenceServerCode int                               `json:"iServerCode"`
	Message                string                            `json:"message"`
	AdditionalProperties   MstrRestErrorAdditionalProperties `json:"additionalProperties"`
}

type MstrRestErrorAdditionalProperties struct {
	ObjectID   string `json:"objectId"`
	ObjectName string `json:"objectName"`
}

func (e *MstrRestError) Error() string {
	return fmt.Sprintf("MicroStrategy REST API Error: %s", e.Message)
}
