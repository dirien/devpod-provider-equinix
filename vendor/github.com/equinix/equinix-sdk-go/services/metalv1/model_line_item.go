/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the LineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItem{}

// LineItem struct for LineItem
type LineItem struct {
	Amount      *float32       `json:"amount,omitempty"`
	Currency    *string        `json:"currency,omitempty"`
	Description *string        `json:"description,omitempty"`
	Details     *string        `json:"details,omitempty"`
	Plan        *PlanIdName    `json:"plan,omitempty"`
	Unit        *string        `json:"unit,omitempty"`
	UnitPrice   *float32       `json:"unit_price,omitempty"`
	Hostname    *string        `json:"hostname,omitempty"`
	Location    *string        `json:"location,omitempty"`
	ItemType    *string        `json:"item_type,omitempty"`
	ServiceId   *string        `json:"service_id,omitempty"`
	StartDate   *string        `json:"start_date,omitempty"`
	EndDate     *string        `json:"end_date,omitempty"`
	ProjectId   *string        `json:"project_id,omitempty"`
	PlanId      *string        `json:"plan_id,omitempty"`
	Project     *ProjectIdName `json:"project,omitempty"`
	// Adjustments for the line item
	Adjustments          []LineItemAdjustment `json:"adjustments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LineItem LineItem

// NewLineItem instantiates a new LineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItem() *LineItem {
	this := LineItem{}
	return &this
}

// NewLineItemWithDefaults instantiates a new LineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemWithDefaults() *LineItem {
	this := LineItem{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *LineItem) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *LineItem) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *LineItem) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *LineItem) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *LineItem) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *LineItem) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LineItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LineItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LineItem) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *LineItem) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *LineItem) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *LineItem) SetDetails(v string) {
	o.Details = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *LineItem) GetPlan() PlanIdName {
	if o == nil || IsNil(o.Plan) {
		var ret PlanIdName
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetPlanOk() (*PlanIdName, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *LineItem) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given PlanIdName and assigns it to the Plan field.
func (o *LineItem) SetPlan(v PlanIdName) {
	o.Plan = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *LineItem) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *LineItem) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *LineItem) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *LineItem) GetUnitPrice() float32 {
	if o == nil || IsNil(o.UnitPrice) {
		var ret float32
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetUnitPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *LineItem) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given float32 and assigns it to the UnitPrice field.
func (o *LineItem) SetUnitPrice(v float32) {
	o.UnitPrice = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *LineItem) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *LineItem) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *LineItem) SetHostname(v string) {
	o.Hostname = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *LineItem) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *LineItem) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *LineItem) SetLocation(v string) {
	o.Location = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *LineItem) GetItemType() string {
	if o == nil || IsNil(o.ItemType) {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetItemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemType) {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *LineItem) HasItemType() bool {
	if o != nil && !IsNil(o.ItemType) {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *LineItem) SetItemType(v string) {
	o.ItemType = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *LineItem) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *LineItem) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *LineItem) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *LineItem) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *LineItem) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *LineItem) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *LineItem) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *LineItem) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *LineItem) SetEndDate(v string) {
	o.EndDate = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *LineItem) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *LineItem) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *LineItem) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *LineItem) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *LineItem) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *LineItem) SetPlanId(v string) {
	o.PlanId = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *LineItem) GetProject() ProjectIdName {
	if o == nil || IsNil(o.Project) {
		var ret ProjectIdName
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetProjectOk() (*ProjectIdName, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *LineItem) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given ProjectIdName and assigns it to the Project field.
func (o *LineItem) SetProject(v ProjectIdName) {
	o.Project = &v
}

// GetAdjustments returns the Adjustments field value if set, zero value otherwise.
func (o *LineItem) GetAdjustments() []LineItemAdjustment {
	if o == nil || IsNil(o.Adjustments) {
		var ret []LineItemAdjustment
		return ret
	}
	return o.Adjustments
}

// GetAdjustmentsOk returns a tuple with the Adjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetAdjustmentsOk() ([]LineItemAdjustment, bool) {
	if o == nil || IsNil(o.Adjustments) {
		return nil, false
	}
	return o.Adjustments, true
}

// HasAdjustments returns a boolean if a field has been set.
func (o *LineItem) HasAdjustments() bool {
	if o != nil && !IsNil(o.Adjustments) {
		return true
	}

	return false
}

// SetAdjustments gets a reference to the given []LineItemAdjustment and assigns it to the Adjustments field.
func (o *LineItem) SetAdjustments(v []LineItemAdjustment) {
	o.Adjustments = v
}

func (o LineItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unit_price"] = o.UnitPrice
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.ItemType) {
		toSerialize["item_type"] = o.ItemType
	}
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Adjustments) {
		toSerialize["adjustments"] = o.Adjustments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LineItem) UnmarshalJSON(data []byte) (err error) {
	varLineItem := _LineItem{}

	err = json.Unmarshal(data, &varLineItem)

	if err != nil {
		return err
	}

	*o = LineItem(varLineItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "description")
		delete(additionalProperties, "details")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "unit_price")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "location")
		delete(additionalProperties, "item_type")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "plan_id")
		delete(additionalProperties, "project")
		delete(additionalProperties, "adjustments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLineItem struct {
	value *LineItem
	isSet bool
}

func (v NullableLineItem) Get() *LineItem {
	return v.value
}

func (v *NullableLineItem) Set(val *LineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItem(val *LineItem) *NullableLineItem {
	return &NullableLineItem{value: val, isSet: true}
}

func (v NullableLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
