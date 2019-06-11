package model

import "encoding/json"

// See https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules_example-events.html

// InvokingEvent is the event that triggers evaluation for a rule.
type InvokingEvent struct {
	ConfigurationItem        ConfigurationItem `json:"configurationItem"`
	NotificationCreationTime string            `json:"notificationCreationTime"`
	MessageType              string            `json:"messageType"`
	RecordVersion            string            `json:"recordVersion"`
}

// ConfigurationItem.Configuration varies according to the type of the item that has changed
// See https://docs.aws.amazon.com/config/latest/developerguide/config-item-table.html
type ConfigurationItem struct {
	Relationships                []Relationship    `json:"relationships"`
	Configuration                json.RawMessage   `json:"configuration"`
	Tags                         map[string]string `json:"tags"`
	ConfigurationItemVersion     string            `json:"configurationItemVersion"`
	ConfigurationItemCaptureTime string            `json:"configurationItemCaptureTime"`
	ConfigurationStateID         int64             `json:"configurationStateId"`
	AwsAccountID                 string            `json:"awsAccountId"`
	ConfigurationItemStatus      string            `json:"configurationItemStatus"`
	ResourceType                 string            `json:"resourceType"`
	ResourceID                   string            `json:"resourceId"`
	Arn                          string            `json:"ARN"`
	AwsRegion                    string            `json:"awsRegion"`
	AvailabilityZone             string            `json:"availabilityZone"`
	ConfigurationStateMd5Hash    string            `json:"configurationStateMd5Hash"`
	ResourceCreationTime         string            `json:"resourceCreationTime"`
}

type Relationship struct {
	ResourceID   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	Name         string `json:"name"`
}