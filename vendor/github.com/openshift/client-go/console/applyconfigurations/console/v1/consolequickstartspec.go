// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	authorizationv1 "k8s.io/api/authorization/v1"
)

// ConsoleQuickStartSpecApplyConfiguration represents a declarative configuration of the ConsoleQuickStartSpec type for use
// with apply.
type ConsoleQuickStartSpecApplyConfiguration struct {
	DisplayName           *string                                   `json:"displayName,omitempty"`
	Icon                  *string                                   `json:"icon,omitempty"`
	Tags                  []string                                  `json:"tags,omitempty"`
	DurationMinutes       *int                                      `json:"durationMinutes,omitempty"`
	Description           *string                                   `json:"description,omitempty"`
	Prerequisites         []string                                  `json:"prerequisites,omitempty"`
	Introduction          *string                                   `json:"introduction,omitempty"`
	Tasks                 []ConsoleQuickStartTaskApplyConfiguration `json:"tasks,omitempty"`
	Conclusion            *string                                   `json:"conclusion,omitempty"`
	NextQuickStart        []string                                  `json:"nextQuickStart,omitempty"`
	AccessReviewResources []authorizationv1.ResourceAttributes      `json:"accessReviewResources,omitempty"`
}

// ConsoleQuickStartSpecApplyConfiguration constructs a declarative configuration of the ConsoleQuickStartSpec type for use with
// apply.
func ConsoleQuickStartSpec() *ConsoleQuickStartSpecApplyConfiguration {
	return &ConsoleQuickStartSpecApplyConfiguration{}
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithDisplayName(value string) *ConsoleQuickStartSpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithIcon sets the Icon field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Icon field is set to the value of the last call.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithIcon(value string) *ConsoleQuickStartSpecApplyConfiguration {
	b.Icon = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithTags(values ...string) *ConsoleQuickStartSpecApplyConfiguration {
	for i := range values {
		b.Tags = append(b.Tags, values[i])
	}
	return b
}

// WithDurationMinutes sets the DurationMinutes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DurationMinutes field is set to the value of the last call.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithDurationMinutes(value int) *ConsoleQuickStartSpecApplyConfiguration {
	b.DurationMinutes = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithDescription(value string) *ConsoleQuickStartSpecApplyConfiguration {
	b.Description = &value
	return b
}

// WithPrerequisites adds the given value to the Prerequisites field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Prerequisites field.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithPrerequisites(values ...string) *ConsoleQuickStartSpecApplyConfiguration {
	for i := range values {
		b.Prerequisites = append(b.Prerequisites, values[i])
	}
	return b
}

// WithIntroduction sets the Introduction field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Introduction field is set to the value of the last call.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithIntroduction(value string) *ConsoleQuickStartSpecApplyConfiguration {
	b.Introduction = &value
	return b
}

// WithTasks adds the given value to the Tasks field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tasks field.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithTasks(values ...*ConsoleQuickStartTaskApplyConfiguration) *ConsoleQuickStartSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTasks")
		}
		b.Tasks = append(b.Tasks, *values[i])
	}
	return b
}

// WithConclusion sets the Conclusion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conclusion field is set to the value of the last call.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithConclusion(value string) *ConsoleQuickStartSpecApplyConfiguration {
	b.Conclusion = &value
	return b
}

// WithNextQuickStart adds the given value to the NextQuickStart field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NextQuickStart field.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithNextQuickStart(values ...string) *ConsoleQuickStartSpecApplyConfiguration {
	for i := range values {
		b.NextQuickStart = append(b.NextQuickStart, values[i])
	}
	return b
}

// WithAccessReviewResources adds the given value to the AccessReviewResources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AccessReviewResources field.
func (b *ConsoleQuickStartSpecApplyConfiguration) WithAccessReviewResources(values ...authorizationv1.ResourceAttributes) *ConsoleQuickStartSpecApplyConfiguration {
	for i := range values {
		b.AccessReviewResources = append(b.AccessReviewResources, values[i])
	}
	return b
}
