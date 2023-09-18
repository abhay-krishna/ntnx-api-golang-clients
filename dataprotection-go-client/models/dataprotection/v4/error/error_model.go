/*
 * Generated file models/dataprotection/v4/error/error_model.go.
 *
 * Product version: 4.0.1-alpha-4
 *
 * Part of the Nutanix Dataprotection Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Auto-injected error response object by the dev platform
*/
package error

import (
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/config"
)

/*
This schema is generated from AppMessage.java
*/
type AppMessage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ArgumentsMap map[string]string `json:"argumentsMap,omitempty"`

	Code *string `json:"code,omitempty"`

	ErrorGroup *string `json:"errorGroup,omitempty"`
	/*
	  The locale for the message description.
	*/
	Locale *string `json:"locale,omitempty"`

	Message *string `json:"message,omitempty"`

	Severity *import1.MessageSeverity `json:"severity,omitempty"`
}

func NewAppMessage() *AppMessage {
	p := new(AppMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.error.AppMessage"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.error.AppMessage"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Locale = new(string)
	*p.Locale = "en_US"

	return p
}

/*
This schema is auto-generated by the Open API Dev Platform as REST response for 4xx and 5xx error responses.
*/
type ErrorResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ErrorItemDiscriminator_ *string `json:"$errorItemDiscriminator,omitempty"`

	Error *OneOfErrorResponseError `json:"error,omitempty"`
}

func NewErrorResponse() *ErrorResponse {
	p := new(ErrorResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.error.ErrorResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.error.ErrorResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ErrorResponse) GetError() interface{} {
	if nil == p.Error {
		return nil
	}
	return p.Error.GetValue()
}

func (p *ErrorResponse) SetError(v interface{}) error {
	if nil == p.Error {
		p.Error = NewOneOfErrorResponseError()
	}
	e := p.Error.SetValue(v)
	if nil == e {
		if nil == p.ErrorItemDiscriminator_ {
			p.ErrorItemDiscriminator_ = new(string)
		}
		*p.ErrorItemDiscriminator_ = *p.Error.Discriminator
	}
	return e
}

/*
This schema is generated from SchemaValidationError.java
*/
type SchemaValidationError struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Error *string `json:"error,omitempty"`

	Path *string `json:"path,omitempty"`

	StatusCode *int `json:"statusCode,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`

	ValidationErrorMessages []SchemaValidationErrorMessage `json:"validationErrorMessages,omitempty"`
}

func NewSchemaValidationError() *SchemaValidationError {
	p := new(SchemaValidationError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.error.SchemaValidationError"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.error.SchemaValidationError"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This schema is generated from SchemaValidationErrorMessage.java
*/
type SchemaValidationErrorMessage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AttributePath *string `json:"attributePath,omitempty"`

	Location *string `json:"location,omitempty"`

	Message *string `json:"message,omitempty"`
}

func NewSchemaValidationErrorMessage() *SchemaValidationErrorMessage {
	p := new(SchemaValidationErrorMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.error.SchemaValidationErrorMessage"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.error.SchemaValidationErrorMessage"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfErrorResponseError struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType201  []AppMessage           `json:"-"`
	oneOfType202  *SchemaValidationError `json:"-"`
}

func NewOneOfErrorResponseError() *OneOfErrorResponseError {
	p := new(OneOfErrorResponseError)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfErrorResponseError) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfErrorResponseError is nil"))
	}
	switch v.(type) {
	case []AppMessage:
		p.oneOfType201 = v.([]AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.error.AppMessage>"
	case SchemaValidationError:
		if nil == p.oneOfType202 {
			p.oneOfType202 = new(SchemaValidationError)
		}
		*p.oneOfType202 = v.(SchemaValidationError)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType202.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType202.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfErrorResponseError) GetValue() interface{} {
	if "List<dataprotection.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType201
	}
	if p.oneOfType202 != nil && *p.oneOfType202.ObjectType_ == *p.Discriminator {
		return *p.oneOfType202
	}
	return nil
}

func (p *OneOfErrorResponseError) UnmarshalJSON(b []byte) error {
	vOneOfType201 := new([]AppMessage)
	if err := json.Unmarshal(b, vOneOfType201); err == nil {
		if len(*vOneOfType201) == 0 || "dataprotection.v4.error.AppMessage" == *((*vOneOfType201)[0].ObjectType_) {
			p.oneOfType201 = *vOneOfType201
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.error.AppMessage>"
			return nil

		}
	}
	vOneOfType202 := new(SchemaValidationError)
	if err := json.Unmarshal(b, vOneOfType202); err == nil {
		if "dataprotection.v4.error.SchemaValidationError" == *vOneOfType202.ObjectType_ {
			if nil == p.oneOfType202 {
				p.oneOfType202 = new(SchemaValidationError)
			}
			*p.oneOfType202 = *vOneOfType202
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType202.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType202.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfErrorResponseError"))
}

func (p *OneOfErrorResponseError) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType201)
	}
	if p.oneOfType202 != nil && *p.oneOfType202.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType202)
	}
	return nil, errors.New("No value to marshal for OneOfErrorResponseError")
}
