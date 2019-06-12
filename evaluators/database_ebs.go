package evaluators

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spaceapegames/config-rules/model"
	"time"
)

type DbEbsOptmizationEvaluation struct {
	ResourceId string
	Status string
	NotificationTime string
	ResultToken string
}

func EvaluateDbEbsOptimization(event events.ConfigEvent) (*DbEbsOptmizationEvaluation, error)  {
	
	// unmarshal the Invoking Event
	var invokingEvent model.InvokingEvent
	err := json.Unmarshal([]byte(event.InvokingEvent), &invokingEvent)
	if err != nil {
		return nil, err
	}

	// unmarshal the ConfigurationItem
	var instance ec2.Instance
	err = json.Unmarshal(invokingEvent.ConfigurationItem.Configuration, &instance)
	if err != nil {
		return nil, err
	}

	// if eventLeftScope then return immediately with NOT_APPLICABLE.
	// (this is generally used to de-scope terminated instances)
	if event.EventLeftScope {
		return &DbEbsOptmizationEvaluation{
			ResourceId: invokingEvent.ConfigurationItem.ResourceID,
			Status: configservice.ComplianceTypeNotApplicable,
			NotificationTime: invokingEvent.NotificationCreationTime,
			ResultToken: event.ResultToken,
		}, nil
	}

	if service, ok := invokingEvent.ConfigurationItem.Tags["SERVICE"]; ok && service == "database" {
		if environment, ok := invokingEvent.ConfigurationItem.Tags["ENVIRONMENT"]; ok && environment == "Production" {
			if !*instance.EbsOptimized {
				// A production database instance _without_ EBS optimization? Non-compliant mate.
				return &DbEbsOptmizationEvaluation{
					ResourceId: invokingEvent.ConfigurationItem.ResourceID,
					Status: configservice.ComplianceTypeNonCompliant,
					NotificationTime: invokingEvent.NotificationCreationTime,
					ResultToken: event.ResultToken,
				}, nil
			}
		}
	}

	// This is NOT a Production database instance, so mark as compliant
	return &DbEbsOptmizationEvaluation{
		ResourceId: invokingEvent.ConfigurationItem.ResourceID,
		Status: configservice.ComplianceTypeCompliant,
		NotificationTime: invokingEvent.NotificationCreationTime,
		ResultToken: event.ResultToken,
	}, nil

}

func CompleteEvaluation(evaluation *DbEbsOptmizationEvaluation) error {
	svc := configservice.New(session.Must(session.NewSession()))

	orderingTimestamp, err := time.Parse(time.RFC3339, evaluation.NotificationTime)
	if err != nil {
		return err
	}

	_, err = svc.PutEvaluations(
		&configservice.PutEvaluationsInput{
			Evaluations: []*configservice.Evaluation{
				{
					ComplianceResourceId:   &evaluation.ResourceId,
					ComplianceResourceType: aws.String("AWS::EC2::Instance"),
					ComplianceType:         aws.String(evaluation.Status),
					OrderingTimestamp:      &orderingTimestamp,
				},
			},
			ResultToken: aws.String(evaluation.ResultToken),
		},
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
