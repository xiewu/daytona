/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daytonaapiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UsageOverview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageOverview{}

// UsageOverview struct for UsageOverview
type UsageOverview struct {
	TotalCpuQuota            float32 `json:"totalCpuQuota"`
	TotalGpuQuota            float32 `json:"totalGpuQuota"`
	TotalMemoryQuota         float32 `json:"totalMemoryQuota"`
	TotalDiskQuota           float32 `json:"totalDiskQuota"`
	TotalWorkspaceQuota      float32 `json:"totalWorkspaceQuota"`
	ConcurrentWorkspaceQuota float32 `json:"concurrentWorkspaceQuota"`
	CurrentCpuUsage          float32 `json:"currentCpuUsage"`
	CurrentMemoryUsage       float32 `json:"currentMemoryUsage"`
	CurrentDiskUsage         float32 `json:"currentDiskUsage"`
	CurrentWorkspaces        float32 `json:"currentWorkspaces"`
	ConcurrentWorkspaces     float32 `json:"concurrentWorkspaces"`
	CurrentImageNumber       float32 `json:"currentImageNumber"`
	ImageQuota               float32 `json:"imageQuota"`
	TotalImageSizeQuota      float32 `json:"totalImageSizeQuota"`
	TotalImageSizeUsed       float32 `json:"totalImageSizeUsed"`
	MaxVolumes               float32 `json:"maxVolumes"`
	UsedVolumes              float32 `json:"usedVolumes"`
}

type _UsageOverview UsageOverview

// NewUsageOverview instantiates a new UsageOverview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageOverview(totalCpuQuota float32, totalGpuQuota float32, totalMemoryQuota float32, totalDiskQuota float32, totalWorkspaceQuota float32, concurrentWorkspaceQuota float32, currentCpuUsage float32, currentMemoryUsage float32, currentDiskUsage float32, currentWorkspaces float32, concurrentWorkspaces float32, currentImageNumber float32, imageQuota float32, totalImageSizeQuota float32, totalImageSizeUsed float32, maxVolumes float32, usedVolumes float32) *UsageOverview {
	this := UsageOverview{}
	this.TotalCpuQuota = totalCpuQuota
	this.TotalGpuQuota = totalGpuQuota
	this.TotalMemoryQuota = totalMemoryQuota
	this.TotalDiskQuota = totalDiskQuota
	this.TotalWorkspaceQuota = totalWorkspaceQuota
	this.ConcurrentWorkspaceQuota = concurrentWorkspaceQuota
	this.CurrentCpuUsage = currentCpuUsage
	this.CurrentMemoryUsage = currentMemoryUsage
	this.CurrentDiskUsage = currentDiskUsage
	this.CurrentWorkspaces = currentWorkspaces
	this.ConcurrentWorkspaces = concurrentWorkspaces
	this.CurrentImageNumber = currentImageNumber
	this.ImageQuota = imageQuota
	this.TotalImageSizeQuota = totalImageSizeQuota
	this.TotalImageSizeUsed = totalImageSizeUsed
	this.MaxVolumes = maxVolumes
	this.UsedVolumes = usedVolumes
	return &this
}

// NewUsageOverviewWithDefaults instantiates a new UsageOverview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageOverviewWithDefaults() *UsageOverview {
	this := UsageOverview{}
	return &this
}

// GetTotalCpuQuota returns the TotalCpuQuota field value
func (o *UsageOverview) GetTotalCpuQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCpuQuota
}

// GetTotalCpuQuotaOk returns a tuple with the TotalCpuQuota field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetTotalCpuQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCpuQuota, true
}

// SetTotalCpuQuota sets field value
func (o *UsageOverview) SetTotalCpuQuota(v float32) {
	o.TotalCpuQuota = v
}

// GetTotalGpuQuota returns the TotalGpuQuota field value
func (o *UsageOverview) GetTotalGpuQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalGpuQuota
}

// GetTotalGpuQuotaOk returns a tuple with the TotalGpuQuota field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetTotalGpuQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalGpuQuota, true
}

// SetTotalGpuQuota sets field value
func (o *UsageOverview) SetTotalGpuQuota(v float32) {
	o.TotalGpuQuota = v
}

// GetTotalMemoryQuota returns the TotalMemoryQuota field value
func (o *UsageOverview) GetTotalMemoryQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalMemoryQuota
}

// GetTotalMemoryQuotaOk returns a tuple with the TotalMemoryQuota field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetTotalMemoryQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMemoryQuota, true
}

// SetTotalMemoryQuota sets field value
func (o *UsageOverview) SetTotalMemoryQuota(v float32) {
	o.TotalMemoryQuota = v
}

// GetTotalDiskQuota returns the TotalDiskQuota field value
func (o *UsageOverview) GetTotalDiskQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalDiskQuota
}

// GetTotalDiskQuotaOk returns a tuple with the TotalDiskQuota field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetTotalDiskQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDiskQuota, true
}

// SetTotalDiskQuota sets field value
func (o *UsageOverview) SetTotalDiskQuota(v float32) {
	o.TotalDiskQuota = v
}

// GetTotalWorkspaceQuota returns the TotalWorkspaceQuota field value
func (o *UsageOverview) GetTotalWorkspaceQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalWorkspaceQuota
}

// GetTotalWorkspaceQuotaOk returns a tuple with the TotalWorkspaceQuota field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetTotalWorkspaceQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalWorkspaceQuota, true
}

// SetTotalWorkspaceQuota sets field value
func (o *UsageOverview) SetTotalWorkspaceQuota(v float32) {
	o.TotalWorkspaceQuota = v
}

// GetConcurrentWorkspaceQuota returns the ConcurrentWorkspaceQuota field value
func (o *UsageOverview) GetConcurrentWorkspaceQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ConcurrentWorkspaceQuota
}

// GetConcurrentWorkspaceQuotaOk returns a tuple with the ConcurrentWorkspaceQuota field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetConcurrentWorkspaceQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConcurrentWorkspaceQuota, true
}

// SetConcurrentWorkspaceQuota sets field value
func (o *UsageOverview) SetConcurrentWorkspaceQuota(v float32) {
	o.ConcurrentWorkspaceQuota = v
}

// GetCurrentCpuUsage returns the CurrentCpuUsage field value
func (o *UsageOverview) GetCurrentCpuUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentCpuUsage
}

// GetCurrentCpuUsageOk returns a tuple with the CurrentCpuUsage field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetCurrentCpuUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentCpuUsage, true
}

// SetCurrentCpuUsage sets field value
func (o *UsageOverview) SetCurrentCpuUsage(v float32) {
	o.CurrentCpuUsage = v
}

// GetCurrentMemoryUsage returns the CurrentMemoryUsage field value
func (o *UsageOverview) GetCurrentMemoryUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentMemoryUsage
}

// GetCurrentMemoryUsageOk returns a tuple with the CurrentMemoryUsage field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetCurrentMemoryUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentMemoryUsage, true
}

// SetCurrentMemoryUsage sets field value
func (o *UsageOverview) SetCurrentMemoryUsage(v float32) {
	o.CurrentMemoryUsage = v
}

// GetCurrentDiskUsage returns the CurrentDiskUsage field value
func (o *UsageOverview) GetCurrentDiskUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentDiskUsage
}

// GetCurrentDiskUsageOk returns a tuple with the CurrentDiskUsage field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetCurrentDiskUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentDiskUsage, true
}

// SetCurrentDiskUsage sets field value
func (o *UsageOverview) SetCurrentDiskUsage(v float32) {
	o.CurrentDiskUsage = v
}

// GetCurrentWorkspaces returns the CurrentWorkspaces field value
func (o *UsageOverview) GetCurrentWorkspaces() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentWorkspaces
}

// GetCurrentWorkspacesOk returns a tuple with the CurrentWorkspaces field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetCurrentWorkspacesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentWorkspaces, true
}

// SetCurrentWorkspaces sets field value
func (o *UsageOverview) SetCurrentWorkspaces(v float32) {
	o.CurrentWorkspaces = v
}

// GetConcurrentWorkspaces returns the ConcurrentWorkspaces field value
func (o *UsageOverview) GetConcurrentWorkspaces() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ConcurrentWorkspaces
}

// GetConcurrentWorkspacesOk returns a tuple with the ConcurrentWorkspaces field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetConcurrentWorkspacesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConcurrentWorkspaces, true
}

// SetConcurrentWorkspaces sets field value
func (o *UsageOverview) SetConcurrentWorkspaces(v float32) {
	o.ConcurrentWorkspaces = v
}

// GetCurrentImageNumber returns the CurrentImageNumber field value
func (o *UsageOverview) GetCurrentImageNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentImageNumber
}

// GetCurrentImageNumberOk returns a tuple with the CurrentImageNumber field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetCurrentImageNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentImageNumber, true
}

// SetCurrentImageNumber sets field value
func (o *UsageOverview) SetCurrentImageNumber(v float32) {
	o.CurrentImageNumber = v
}

// GetImageQuota returns the ImageQuota field value
func (o *UsageOverview) GetImageQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ImageQuota
}

// GetImageQuotaOk returns a tuple with the ImageQuota field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetImageQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageQuota, true
}

// SetImageQuota sets field value
func (o *UsageOverview) SetImageQuota(v float32) {
	o.ImageQuota = v
}

// GetTotalImageSizeQuota returns the TotalImageSizeQuota field value
func (o *UsageOverview) GetTotalImageSizeQuota() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalImageSizeQuota
}

// GetTotalImageSizeQuotaOk returns a tuple with the TotalImageSizeQuota field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetTotalImageSizeQuotaOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalImageSizeQuota, true
}

// SetTotalImageSizeQuota sets field value
func (o *UsageOverview) SetTotalImageSizeQuota(v float32) {
	o.TotalImageSizeQuota = v
}

// GetTotalImageSizeUsed returns the TotalImageSizeUsed field value
func (o *UsageOverview) GetTotalImageSizeUsed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalImageSizeUsed
}

// GetTotalImageSizeUsedOk returns a tuple with the TotalImageSizeUsed field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetTotalImageSizeUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalImageSizeUsed, true
}

// SetTotalImageSizeUsed sets field value
func (o *UsageOverview) SetTotalImageSizeUsed(v float32) {
	o.TotalImageSizeUsed = v
}

// GetMaxVolumes returns the MaxVolumes field value
func (o *UsageOverview) GetMaxVolumes() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxVolumes
}

// GetMaxVolumesOk returns a tuple with the MaxVolumes field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetMaxVolumesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxVolumes, true
}

// SetMaxVolumes sets field value
func (o *UsageOverview) SetMaxVolumes(v float32) {
	o.MaxVolumes = v
}

// GetUsedVolumes returns the UsedVolumes field value
func (o *UsageOverview) GetUsedVolumes() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UsedVolumes
}

// GetUsedVolumesOk returns a tuple with the UsedVolumes field value
// and a boolean to check if the value has been set.
func (o *UsageOverview) GetUsedVolumesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedVolumes, true
}

// SetUsedVolumes sets field value
func (o *UsageOverview) SetUsedVolumes(v float32) {
	o.UsedVolumes = v
}

func (o UsageOverview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageOverview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["totalCpuQuota"] = o.TotalCpuQuota
	toSerialize["totalGpuQuota"] = o.TotalGpuQuota
	toSerialize["totalMemoryQuota"] = o.TotalMemoryQuota
	toSerialize["totalDiskQuota"] = o.TotalDiskQuota
	toSerialize["totalWorkspaceQuota"] = o.TotalWorkspaceQuota
	toSerialize["concurrentWorkspaceQuota"] = o.ConcurrentWorkspaceQuota
	toSerialize["currentCpuUsage"] = o.CurrentCpuUsage
	toSerialize["currentMemoryUsage"] = o.CurrentMemoryUsage
	toSerialize["currentDiskUsage"] = o.CurrentDiskUsage
	toSerialize["currentWorkspaces"] = o.CurrentWorkspaces
	toSerialize["concurrentWorkspaces"] = o.ConcurrentWorkspaces
	toSerialize["currentImageNumber"] = o.CurrentImageNumber
	toSerialize["imageQuota"] = o.ImageQuota
	toSerialize["totalImageSizeQuota"] = o.TotalImageSizeQuota
	toSerialize["totalImageSizeUsed"] = o.TotalImageSizeUsed
	toSerialize["maxVolumes"] = o.MaxVolumes
	toSerialize["usedVolumes"] = o.UsedVolumes
	return toSerialize, nil
}

func (o *UsageOverview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"totalCpuQuota",
		"totalGpuQuota",
		"totalMemoryQuota",
		"totalDiskQuota",
		"totalWorkspaceQuota",
		"concurrentWorkspaceQuota",
		"currentCpuUsage",
		"currentMemoryUsage",
		"currentDiskUsage",
		"currentWorkspaces",
		"concurrentWorkspaces",
		"currentImageNumber",
		"imageQuota",
		"totalImageSizeQuota",
		"totalImageSizeUsed",
		"maxVolumes",
		"usedVolumes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUsageOverview := _UsageOverview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsageOverview)

	if err != nil {
		return err
	}

	*o = UsageOverview(varUsageOverview)

	return err
}

type NullableUsageOverview struct {
	value *UsageOverview
	isSet bool
}

func (v NullableUsageOverview) Get() *UsageOverview {
	return v.value
}

func (v *NullableUsageOverview) Set(val *UsageOverview) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageOverview) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageOverview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageOverview(val *UsageOverview) *NullableUsageOverview {
	return &NullableUsageOverview{value: val, isSet: true}
}

func (v NullableUsageOverview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageOverview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
