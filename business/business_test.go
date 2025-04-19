package business

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BusinessRuleTestSuite struct {
	suite.Suite
}

func (suite *BusinessRuleTestSuite) SetupAllSuite() {
	// open database connection
}

func (suite *BusinessRuleTestSuite) TearDownAllSuite() {
	// close database connection
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BusinessRuleTestSuite))
}

func (suite *BusinessRuleTestSuite) TestExecuteBusinessRule() {
	jsonString := "{\"nome\": \"Dóia\", \"idade\": 3, \"pedigree\": {\"raça\": \"viralata\", \"valorBRL\": 0.50}}"
	result := BusinessRule(jsonString)
	suite.Equal("Dóia", result["nome"])
	suite.Equal(3.0, result["idade"])

	if pedigree, ok := result["pedigree"].(map[string]any); ok {
		suite.Equal("viralata", pedigree["raça"])
		suite.Equal(0.5, pedigree["valorBRL"])
	} else {
		suite.FailNow("JSON structure is not the expected")
	}
}

func (suite *BusinessRuleTestSuite) TestBusinessRulePanicsWithInvalidJson() {
	suite.Assert().Panics(func() {
		BusinessRule("money for nothing and chicks for free")
	})
}
