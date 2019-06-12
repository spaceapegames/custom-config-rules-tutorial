package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/spaceapegames/custom-config-rules-tutorial/evaluators"
)

func Handler(event events.ConfigEvent) error {
	result, err := evaluators.EvaluateDbEbsOptimization(event)
	if err != nil {
		return err
	}

	return evaluators.CompleteEvaluation(result)
}

func main() {
	lambda.Start(Handler)
}