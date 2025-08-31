package test

import (
	"testing"

	"github.com/cloudposse/test-helpers/pkg/atmos"
	helper "github.com/cloudposse/test-helpers/pkg/atmos/component-helper"
	"github.com/stretchr/testify/assert"
)

type ComponentSuite struct {
	helper.TestSuite
}

func (s *ComponentSuite) TestBasic() {
	const component = "example/basic"
	const stack = "default-test"
	const awsRegion = "us-east-2"

	var randomString = atmos.RandomString(s.T(), 10)
	inputs := map[string]interface{}{
		"organization_name": randomString,
	}

	defer s.DestroyAtmosComponent(s.T(), component, stack, &inputs)
	options, _ := s.DeployAtmosComponent(s.T(), component, stack, &inputs)
	assert.NotNil(s.T(), options)

	assert.NotEmpty(s.T(), atmos.Output(s.T(), componentInstance, "id"))
	assert.Contains(s.T(), atmos.Output(s.T(), componentInstance, "id"), randomString)
	assert.NotEmpty(s.T(), atmos.Output(s.T(), componentInstance, "description"))
	assert.NotEmpty(s.T(), atmos.Output(s.T(), componentInstance, "public_id"))
	assert.NotEmpty(s.T(), atmos.Output(s.T(), componentInstance, "user"))
	assert.NotEmpty(s.T(), atmos.Output(s.T(), componentInstance, "settings"))
	assert.NotEmpty(s.T(), atmos.Output(s.T(), componentInstance, "api_key"))

	s.DriftTest(component, stack, nil)
}

func (s *ComponentSuite) TestEnabledFlag() {
	const component = "example/disabled"
	const stack = "default-test"
	s.VerifyEnabledFlag(component, stack, nil)
}

func TestRunSuite(t *testing.T) {
	suite := new(ComponentSuite)
	helper.Run(t, suite)
}
