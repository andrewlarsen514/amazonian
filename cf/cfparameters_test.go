package cf

import "testing"

func TestCreateCloudformationParametersKey(t *testing.T) {
	parameters := make(map[string]string)

	parameters["key"] = "value"

	cfParams := createCloudformationParameters(parameters)

	param := cfParams

	println("cf paramets ", param)
	cfParamsLength := len(cfParams)
	mapLength := len(parameters)

	if cfParamsLength != mapLength {
		t.Log("Created parameter length: ", cfParamsLength, " is not equal to mapLength: ", mapLength)
		t.Fail()
	}

	paramValue := param[0]

	if *paramValue.ParameterKey != "key" {
		t.Log("paramater key ", paramValue.GoString(), " is not correct")
		t.Fail()
	}
}
func TestCreateCloudformationParametersValue(t *testing.T) {
	parameters := make(map[string]string)

	parameters["key"] = "value"

	cfParams := createCloudformationParameters(parameters)

	param := cfParams

	println("cf paramets ", param)
	cfParamsLength := len(cfParams)
	mapLength := len(parameters)

	if cfParamsLength != mapLength {
		t.Log("Created parameter length: ", cfParamsLength, " is not equal to mapLength: ", mapLength)
		t.Fail()
	}

	paramValue := param[0]

	if *paramValue.ParameterValue != "value" {
		t.Log("paramater key ", paramValue.GoString(), " is not correct")
		t.Fail()
	}
}

func TestCreateCloudformationParameters2ParametersBothKeysAreCorrect(t *testing.T) {
	parameters := make(map[string]string)

	parameters["key"] = "value"
	parameters["key1"] = "value1"

	cfParams := createCloudformationParameters(parameters)

	param := cfParams

	cfParamsLength := len(cfParams)
	mapLength := len(parameters)

	if cfParamsLength != mapLength {
		t.Log("Created parameter length: ", cfParamsLength, " is not equal to mapLength: ", mapLength)
		t.Fail()
	}

	paramValue1 := param[0]
	paramValue2 := param[1]

	if *paramValue1.ParameterKey != "key" && *paramValue1.ParameterKey != "key1" {
		t.Log("paramater key ", paramValue1.GoString(), " is not correct")
		t.Fail()
	}
	if *paramValue2.ParameterKey != "key" && *paramValue2.ParameterKey != "key1" {
		t.Log("paramater key1 ", paramValue2.GoString(), " is not correct")
		t.Fail()
	}
}
func TestCreateCloudformationParameters2ParametersBothValueAreCorrect(t *testing.T) {
	parameters := make(map[string]string)

	parameters["key"] = "value"
	parameters["key1"] = "value1"

	cfParams := createCloudformationParameters(parameters)

	param := cfParams

	println("cf paramets ", param)
	cfParamsLength := len(cfParams)
	mapLength := len(parameters)

	if cfParamsLength != mapLength {
		t.Log("Created parameter length: ", cfParamsLength, " is not equal to mapLength: ", mapLength)
		t.Fail()
	}

	paramValue1 := param[0]
	paramValue2 := param[1]

	if *paramValue1.ParameterValue != "value" && *paramValue1.ParameterValue != "value1" {
		t.Log("paramater key ", paramValue1.GoString(), " is not correct")
		t.Fail()
	}
	if *paramValue2.ParameterValue != "value" && *paramValue2.ParameterValue != "value1" {
		t.Log("paramater key ", paramValue2.GoString(), " is not correct")
		t.Fail()
	}
}

//CreateClusterParameters tests
func TestCreateClusterParameters(t *testing.T) {
	cluster := EcsCluster{DomainName: "DOMAINAME"}

	params := CreateClusterParameters(cluster)

	if *params[0].ParameterKey != "DomainName" {
		t.Log("paramkey ", params[0].ParameterKey, " did not get set to correct constant value")
	}
	if *params[0].ParameterValue != "DOMAINAME" {
		t.Log("paramvalue ", params[0].ParameterValue, " did not get set to correct constant value")
	}
}
