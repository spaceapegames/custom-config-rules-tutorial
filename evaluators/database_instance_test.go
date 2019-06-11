package evaluators

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestEvaluateDatabaseTerminationProtection(t *testing.T) {
	EbsOptimized, err := ioutil.ReadFile("../fixtures/sampleConfigItem.json")
	if err != nil {
		t.Fatal(err)
	}
	NonEbsOptimized, err := ioutil.ReadFile("../fixtures/sampleConfigItemNonOptimized.json")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name  string
		event events.ConfigEvent
		want  string
	}{
		{
			"eventLeftScope",
			events.ConfigEvent{
				EventLeftScope: true,
				InvokingEvent:  string(NonEbsOptimized),
			},
			configservice.ComplianceTypeNotApplicable,
		},
		{
			"nonEBSOptimized",
			events.ConfigEvent{
				EventLeftScope: false,
				InvokingEvent:  string(NonEbsOptimized),
			},
			configservice.ComplianceTypeNonCompliant,
		},
		{
			"EBSOptimized",
			events.ConfigEvent{
				EventLeftScope: false,
				InvokingEvent:  string(EbsOptimized),
			},
			configservice.ComplianceTypeCompliant,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := EvaluateDatabaseTerminationProtection(c.event)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, c.want, result.Status)
		})
	}
}
