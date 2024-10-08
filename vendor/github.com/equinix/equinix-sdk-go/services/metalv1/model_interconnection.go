/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"time"
)

// checks if the Interconnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Interconnection{}

// Interconnection struct for Interconnection
type Interconnection struct {
	ContactEmail *string              `json:"contact_email,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Facility     *Facility            `json:"facility,omitempty"`
	Id           *string              `json:"id,omitempty"`
	Metro        *Metro               `json:"metro,omitempty"`
	Mode         *InterconnectionMode `json:"mode,omitempty"`
	Name         *string              `json:"name,omitempty"`
	Organization *Organization        `json:"organization,omitempty"`
	// For Fabric VCs, these represent Virtual Port(s) created for the interconnection. For dedicated interconnections, these represent the Dedicated Port(s).
	Ports      []InterconnectionPort      `json:"ports,omitempty"`
	Project    *Project                   `json:"project,omitempty"`
	Redundancy *InterconnectionRedundancy `json:"redundancy,omitempty"`
	// For Fabric VCs (Metal Billed), this will show details of the A-Side service tokens issued for the interconnection. For Fabric VCs (Fabric Billed), this will show the details of the Z-Side service tokens issued for the interconnection. Dedicated interconnections will not have any service tokens issued. There will be one per interconnection, so for redundant interconnections, there should be two service tokens issued.
	ServiceTokens []FabricServiceToken `json:"service_tokens,omitempty"`
	// For Fabric VCs (Metal Billed), this allows Fabric to connect the Metal network to any connection Fabric facilitates. Fabric uses this token to be able to give more detailed information about the Metal end of the network, when viewing resources from within Fabric.
	AuthorizationCode *string `json:"authorization_code,omitempty"`
	// For interconnections on Dedicated Ports and shared connections, this represents the interconnection's speed in bps. For Fabric VCs, this field refers to the maximum speed of the interconnection in bps. This value will default to 10Gbps for Fabric VCs (Fabric Billed).
	Speed  *int64   `json:"speed,omitempty"`
	Status *string  `json:"status,omitempty"`
	Tags   []string `json:"tags,omitempty"`
	// This token is used for shared interconnections to be used as the Fabric Token. This field is entirely deprecated.
	Token                *string                        `json:"token,omitempty"`
	Type                 *InterconnectionType           `json:"type,omitempty"`
	FabricProvider       *InterconnectionFabricProvider `json:"fabric_provider,omitempty"`
	CreatedAt            *time.Time                     `json:"created_at,omitempty"`
	UpdatedAt            *time.Time                     `json:"updated_at,omitempty"`
	RequestedBy          *Href                          `json:"requested_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Interconnection Interconnection

// NewInterconnection instantiates a new Interconnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterconnection() *Interconnection {
	this := Interconnection{}
	return &this
}

// NewInterconnectionWithDefaults instantiates a new Interconnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterconnectionWithDefaults() *Interconnection {
	this := Interconnection{}
	return &this
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *Interconnection) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *Interconnection) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *Interconnection) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Interconnection) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Interconnection) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Interconnection) SetDescription(v string) {
	o.Description = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *Interconnection) GetFacility() Facility {
	if o == nil || IsNil(o.Facility) {
		var ret Facility
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetFacilityOk() (*Facility, bool) {
	if o == nil || IsNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *Interconnection) HasFacility() bool {
	if o != nil && !IsNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given Facility and assigns it to the Facility field.
func (o *Interconnection) SetFacility(v Facility) {
	o.Facility = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Interconnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Interconnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Interconnection) SetId(v string) {
	o.Id = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *Interconnection) GetMetro() Metro {
	if o == nil || IsNil(o.Metro) {
		var ret Metro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetMetroOk() (*Metro, bool) {
	if o == nil || IsNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *Interconnection) HasMetro() bool {
	if o != nil && !IsNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given Metro and assigns it to the Metro field.
func (o *Interconnection) SetMetro(v Metro) {
	o.Metro = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Interconnection) GetMode() InterconnectionMode {
	if o == nil || IsNil(o.Mode) {
		var ret InterconnectionMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetModeOk() (*InterconnectionMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Interconnection) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given InterconnectionMode and assigns it to the Mode field.
func (o *Interconnection) SetMode(v InterconnectionMode) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Interconnection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Interconnection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Interconnection) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Interconnection) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Interconnection) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *Interconnection) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *Interconnection) GetPorts() []InterconnectionPort {
	if o == nil || IsNil(o.Ports) {
		var ret []InterconnectionPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetPortsOk() ([]InterconnectionPort, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *Interconnection) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []InterconnectionPort and assigns it to the Ports field.
func (o *Interconnection) SetPorts(v []InterconnectionPort) {
	o.Ports = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Interconnection) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Interconnection) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *Interconnection) SetProject(v Project) {
	o.Project = &v
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise.
func (o *Interconnection) GetRedundancy() InterconnectionRedundancy {
	if o == nil || IsNil(o.Redundancy) {
		var ret InterconnectionRedundancy
		return ret
	}
	return *o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetRedundancyOk() (*InterconnectionRedundancy, bool) {
	if o == nil || IsNil(o.Redundancy) {
		return nil, false
	}
	return o.Redundancy, true
}

// HasRedundancy returns a boolean if a field has been set.
func (o *Interconnection) HasRedundancy() bool {
	if o != nil && !IsNil(o.Redundancy) {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given InterconnectionRedundancy and assigns it to the Redundancy field.
func (o *Interconnection) SetRedundancy(v InterconnectionRedundancy) {
	o.Redundancy = &v
}

// GetServiceTokens returns the ServiceTokens field value if set, zero value otherwise.
func (o *Interconnection) GetServiceTokens() []FabricServiceToken {
	if o == nil || IsNil(o.ServiceTokens) {
		var ret []FabricServiceToken
		return ret
	}
	return o.ServiceTokens
}

// GetServiceTokensOk returns a tuple with the ServiceTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetServiceTokensOk() ([]FabricServiceToken, bool) {
	if o == nil || IsNil(o.ServiceTokens) {
		return nil, false
	}
	return o.ServiceTokens, true
}

// HasServiceTokens returns a boolean if a field has been set.
func (o *Interconnection) HasServiceTokens() bool {
	if o != nil && !IsNil(o.ServiceTokens) {
		return true
	}

	return false
}

// SetServiceTokens gets a reference to the given []FabricServiceToken and assigns it to the ServiceTokens field.
func (o *Interconnection) SetServiceTokens(v []FabricServiceToken) {
	o.ServiceTokens = v
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *Interconnection) GetAuthorizationCode() string {
	if o == nil || IsNil(o.AuthorizationCode) {
		var ret string
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *Interconnection) HasAuthorizationCode() bool {
	if o != nil && !IsNil(o.AuthorizationCode) {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given string and assigns it to the AuthorizationCode field.
func (o *Interconnection) SetAuthorizationCode(v string) {
	o.AuthorizationCode = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *Interconnection) GetSpeed() int64 {
	if o == nil || IsNil(o.Speed) {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetSpeedOk() (*int64, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *Interconnection) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *Interconnection) SetSpeed(v int64) {
	o.Speed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Interconnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Interconnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Interconnection) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Interconnection) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Interconnection) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Interconnection) SetTags(v []string) {
	o.Tags = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Interconnection) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Interconnection) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Interconnection) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Interconnection) GetType() InterconnectionType {
	if o == nil || IsNil(o.Type) {
		var ret InterconnectionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetTypeOk() (*InterconnectionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Interconnection) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given InterconnectionType and assigns it to the Type field.
func (o *Interconnection) SetType(v InterconnectionType) {
	o.Type = &v
}

// GetFabricProvider returns the FabricProvider field value if set, zero value otherwise.
func (o *Interconnection) GetFabricProvider() InterconnectionFabricProvider {
	if o == nil || IsNil(o.FabricProvider) {
		var ret InterconnectionFabricProvider
		return ret
	}
	return *o.FabricProvider
}

// GetFabricProviderOk returns a tuple with the FabricProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetFabricProviderOk() (*InterconnectionFabricProvider, bool) {
	if o == nil || IsNil(o.FabricProvider) {
		return nil, false
	}
	return o.FabricProvider, true
}

// HasFabricProvider returns a boolean if a field has been set.
func (o *Interconnection) HasFabricProvider() bool {
	if o != nil && !IsNil(o.FabricProvider) {
		return true
	}

	return false
}

// SetFabricProvider gets a reference to the given InterconnectionFabricProvider and assigns it to the FabricProvider field.
func (o *Interconnection) SetFabricProvider(v InterconnectionFabricProvider) {
	o.FabricProvider = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Interconnection) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Interconnection) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Interconnection) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Interconnection) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Interconnection) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Interconnection) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetRequestedBy returns the RequestedBy field value if set, zero value otherwise.
func (o *Interconnection) GetRequestedBy() Href {
	if o == nil || IsNil(o.RequestedBy) {
		var ret Href
		return ret
	}
	return *o.RequestedBy
}

// GetRequestedByOk returns a tuple with the RequestedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interconnection) GetRequestedByOk() (*Href, bool) {
	if o == nil || IsNil(o.RequestedBy) {
		return nil, false
	}
	return o.RequestedBy, true
}

// HasRequestedBy returns a boolean if a field has been set.
func (o *Interconnection) HasRequestedBy() bool {
	if o != nil && !IsNil(o.RequestedBy) {
		return true
	}

	return false
}

// SetRequestedBy gets a reference to the given Href and assigns it to the RequestedBy field.
func (o *Interconnection) SetRequestedBy(v Href) {
	o.RequestedBy = &v
}

func (o Interconnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Interconnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactEmail) {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Redundancy) {
		toSerialize["redundancy"] = o.Redundancy
	}
	if !IsNil(o.ServiceTokens) {
		toSerialize["service_tokens"] = o.ServiceTokens
	}
	if !IsNil(o.AuthorizationCode) {
		toSerialize["authorization_code"] = o.AuthorizationCode
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.FabricProvider) {
		toSerialize["fabric_provider"] = o.FabricProvider
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.RequestedBy) {
		toSerialize["requested_by"] = o.RequestedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Interconnection) UnmarshalJSON(data []byte) (err error) {
	varInterconnection := _Interconnection{}

	err = json.Unmarshal(data, &varInterconnection)

	if err != nil {
		return err
	}

	*o = Interconnection(varInterconnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contact_email")
		delete(additionalProperties, "description")
		delete(additionalProperties, "facility")
		delete(additionalProperties, "id")
		delete(additionalProperties, "metro")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "project")
		delete(additionalProperties, "redundancy")
		delete(additionalProperties, "service_tokens")
		delete(additionalProperties, "authorization_code")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "token")
		delete(additionalProperties, "type")
		delete(additionalProperties, "fabric_provider")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "requested_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterconnection struct {
	value *Interconnection
	isSet bool
}

func (v NullableInterconnection) Get() *Interconnection {
	return v.value
}

func (v *NullableInterconnection) Set(val *Interconnection) {
	v.value = val
	v.isSet = true
}

func (v NullableInterconnection) IsSet() bool {
	return v.isSet
}

func (v *NullableInterconnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterconnection(val *Interconnection) *NullableInterconnection {
	return &NullableInterconnection{value: val, isSet: true}
}

func (v NullableInterconnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterconnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
