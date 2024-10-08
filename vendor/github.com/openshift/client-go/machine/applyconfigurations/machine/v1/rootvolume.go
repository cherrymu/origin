// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// RootVolumeApplyConfiguration represents a declarative configuration of the RootVolume type for use
// with apply.
type RootVolumeApplyConfiguration struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	VolumeType       *string `json:"volumeType,omitempty"`
}

// RootVolumeApplyConfiguration constructs a declarative configuration of the RootVolume type for use with
// apply.
func RootVolume() *RootVolumeApplyConfiguration {
	return &RootVolumeApplyConfiguration{}
}

// WithAvailabilityZone sets the AvailabilityZone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailabilityZone field is set to the value of the last call.
func (b *RootVolumeApplyConfiguration) WithAvailabilityZone(value string) *RootVolumeApplyConfiguration {
	b.AvailabilityZone = &value
	return b
}

// WithVolumeType sets the VolumeType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeType field is set to the value of the last call.
func (b *RootVolumeApplyConfiguration) WithVolumeType(value string) *RootVolumeApplyConfiguration {
	b.VolumeType = &value
	return b
}
