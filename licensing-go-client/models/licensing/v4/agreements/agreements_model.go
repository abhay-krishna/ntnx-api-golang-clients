/*
 * Generated file models/licensing/v4/agreements/agreements_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Licensing APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module licensing.v4.agreements of Nutanix Licensing APIs
*/
package agreements

import (
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/error"
	"time"
)

/*
Model containing the EULA acceptance details.
*/
type Acceptance struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Date-time at which EULA was accepted.
	*/
	AcceptanceTime *time.Time `json:"acceptanceTime,omitempty"`

	AcceptedBy *EndUser `json:"acceptedBy"`
}

func (p *Acceptance) MarshalJSON() ([]byte, error) {
	type AcceptanceProxy Acceptance
	return json.Marshal(struct {
		*AcceptanceProxy
		AcceptedBy *EndUser `json:"acceptedBy,omitempty"`
	}{
		AcceptanceProxy: (*AcceptanceProxy)(p),
		AcceptedBy:      p.AcceptedBy,
	})
}

func NewAcceptance() *Acceptance {
	p := new(Acceptance)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.agreements.Acceptance"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /licensing/v4.0/agreements/eula/$actions/add-user Post operation
*/
type AddUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddUserApiResponse() *AddUserApiResponse {
	p := new(AddUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.agreements.AddUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddUserApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddUserApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddUserApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
Model containing the EULA User Details attributes.
*/
type EndUser struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Company name of the user accepting the EULA.
	*/
	CompanyName *string `json:"companyName"`
	/*
	  Job title of the user accepting the EULA.
	*/
	JobTitle *string `json:"jobTitle"`
	/*
	  Login ID of the user accepting the EULA.
	*/
	LoginId *string `json:"loginId"`
	/*
	  User name of the user accepting the EULA.
	*/
	UserName *string `json:"userName"`
}

func (p *EndUser) MarshalJSON() ([]byte, error) {
	type EndUserProxy EndUser
	return json.Marshal(struct {
		*EndUserProxy
		CompanyName *string `json:"companyName,omitempty"`
		JobTitle    *string `json:"jobTitle,omitempty"`
		LoginId     *string `json:"loginId,omitempty"`
		UserName    *string `json:"userName,omitempty"`
	}{
		EndUserProxy: (*EndUserProxy)(p),
		CompanyName:  p.CompanyName,
		JobTitle:     p.JobTitle,
		LoginId:      p.LoginId,
		UserName:     p.UserName,
	})
}

func NewEndUser() *EndUser {
	p := new(EndUser)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.agreements.EndUser"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model containing the EULA info attributes.
*/
type Eula struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of users accepting the EULA along with acceptance time for each.
	*/
	Acceptances []Acceptance `json:"acceptances,omitempty"`
	/*
	  Textual contents of the end user license agreement.
	*/
	Content *string `json:"content"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether this is the current EULA of the cluster or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  EULA update time since epoch in ISO date time.
	*/
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
	/*
	  Version of the EULA.
	*/
	Version *string `json:"version"`
}

func (p *Eula) MarshalJSON() ([]byte, error) {
	type EulaProxy Eula
	return json.Marshal(struct {
		*EulaProxy
		Content *string `json:"content,omitempty"`
		Version *string `json:"version,omitempty"`
	}{
		EulaProxy: (*EulaProxy)(p),
		Content:   p.Content,
		Version:   p.Version,
	})
}

func NewEula() *Eula {
	p := new(Eula)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.agreements.Eula"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /licensing/v4.0/agreements/eula Get operation
*/
type GetEulaApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetEulaApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetEulaApiResponse() *GetEulaApiResponse {
	p := new(GetEulaApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.agreements.GetEulaApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetEulaApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetEulaApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetEulaApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /licensing/v4.0/agreements/eula Put operation
*/
type UpdateEulaApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateEulaApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateEulaApiResponse() *UpdateEulaApiResponse {
	p := new(UpdateEulaApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.agreements.UpdateEulaApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateEulaApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateEulaApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateEulaApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type OneOfGetEulaApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21019 *Eula                  `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetEulaApiResponseData() *OneOfGetEulaApiResponseData {
	p := new(OneOfGetEulaApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetEulaApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetEulaApiResponseData is nil"))
	}
	switch v.(type) {
	case Eula:
		if nil == p.oneOfType21019 {
			p.oneOfType21019 = new(Eula)
		}
		*p.oneOfType21019 = v.(Eula)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType21019.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType21019.ObjectType_
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetEulaApiResponseData) GetValue() interface{} {
	if p.oneOfType21019 != nil && *p.oneOfType21019.ObjectType_ == *p.Discriminator {
		return *p.oneOfType21019
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetEulaApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21019 := new(Eula)
	if err := json.Unmarshal(b, vOneOfType21019); err == nil {
		if "licensing.v4.agreements.Eula" == *vOneOfType21019.ObjectType_ {
			if nil == p.oneOfType21019 {
				p.oneOfType21019 = new(Eula)
			}
			*p.oneOfType21019 = *vOneOfType21019
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType21019.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType21019.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEulaApiResponseData"))
}

func (p *OneOfGetEulaApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType21019 != nil && *p.oneOfType21019.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType21019)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetEulaApiResponseData")
}

type OneOfUpdateEulaApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType21021 []import1.AppMessage   `json:"-"`
}

func NewOneOfUpdateEulaApiResponseData() *OneOfUpdateEulaApiResponseData {
	p := new(OneOfUpdateEulaApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateEulaApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateEulaApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []import1.AppMessage:
		p.oneOfType21021 = v.([]import1.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateEulaApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType21021
	}
	return nil
}

func (p *OneOfUpdateEulaApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType21021 := new([]import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType21021); err == nil {
		if len(*vOneOfType21021) == 0 || "licensing.v4.error.AppMessage" == *((*vOneOfType21021)[0].ObjectType_) {
			p.oneOfType21021 = *vOneOfType21021
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateEulaApiResponseData"))
}

func (p *OneOfUpdateEulaApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21021)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateEulaApiResponseData")
}

type OneOfAddUserApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21023 []import1.AppMessage   `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
}

func NewOneOfAddUserApiResponseData() *OneOfAddUserApiResponseData {
	p := new(OneOfAddUserApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddUserApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddUserApiResponseData is nil"))
	}
	switch v.(type) {
	case []import1.AppMessage:
		p.oneOfType21023 = v.([]import1.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddUserApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType21023
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddUserApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21023 := new([]import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType21023); err == nil {
		if len(*vOneOfType21023) == 0 || "licensing.v4.error.AppMessage" == *((*vOneOfType21023)[0].ObjectType_) {
			p.oneOfType21023 = *vOneOfType21023
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddUserApiResponseData"))
}

func (p *OneOfAddUserApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21023)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddUserApiResponseData")
}

type FileDetail struct {
	Path        *string `json:"-"`
	ObjectType_ *string `json:"-"`
}

func NewFileDetail() *FileDetail {
	p := new(FileDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "FileDetail"

	return p
}
