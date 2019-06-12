STACK_NAME="custom-config-rule-tutorial"
S3_BUCKET ?= "custom-config-rule-tutorial"

build:
	GOOS=linux go build -o ./bin/ebsOptimizationCheck ./handlers/ebsOptimizationCheck/ \
		&& zip -r main.zip bin/
.PHONY: build

package:
	sam package --template-file template.yml --output-template-file packaged.yml --s3-bucket $(S3_BUCKET)
.PHONY: package

deploy: build package
	sam deploy --template-file ./packaged.yml --stack-name $(STACK_NAME) --capabilities CAPABILITY_IAM
.PHONY: deploy

destroy:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
.PHONY: destroy

