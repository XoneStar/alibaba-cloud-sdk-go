package sae

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateApplicationScalingRule invokes the sae.CreateApplicationScalingRule API synchronously
func (client *Client) CreateApplicationScalingRule(request *CreateApplicationScalingRuleRequest) (response *CreateApplicationScalingRuleResponse, err error) {
	response = CreateCreateApplicationScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateApplicationScalingRuleWithChan invokes the sae.CreateApplicationScalingRule API asynchronously
func (client *Client) CreateApplicationScalingRuleWithChan(request *CreateApplicationScalingRuleRequest) (<-chan *CreateApplicationScalingRuleResponse, <-chan error) {
	responseChan := make(chan *CreateApplicationScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApplicationScalingRule(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateApplicationScalingRuleWithCallback invokes the sae.CreateApplicationScalingRule API asynchronously
func (client *Client) CreateApplicationScalingRuleWithCallback(request *CreateApplicationScalingRuleRequest, callback func(response *CreateApplicationScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateApplicationScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateApplicationScalingRule(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateApplicationScalingRuleRequest is the request struct for api CreateApplicationScalingRule
type CreateApplicationScalingRuleRequest struct {
	*requests.RoaRequest
	ScalingRuleName   string `position:"Query" name:"ScalingRuleName"`
	MinReadyInstances string `position:"Query" name:"MinReadyInstances"`
	ScalingRuleTimer  string `position:"Query" name:"ScalingRuleTimer"`
	ScalingRuleMetric string `position:"Query" name:"ScalingRuleMetric"`
	AppId             string `position:"Query" name:"AppId"`
	ScalingRuleType   string `position:"Query" name:"ScalingRuleType"`
}

// CreateApplicationScalingRuleResponse is the response struct for api CreateApplicationScalingRule
type CreateApplicationScalingRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateApplicationScalingRuleRequest creates a request to invoke CreateApplicationScalingRule API
func CreateCreateApplicationScalingRuleRequest() (request *CreateApplicationScalingRuleRequest) {
	request = &CreateApplicationScalingRuleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "CreateApplicationScalingRule", "/pop/v1/sam/scale/applicationScalingRule", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateApplicationScalingRuleResponse creates a response to parse from CreateApplicationScalingRule response
func CreateCreateApplicationScalingRuleResponse() (response *CreateApplicationScalingRuleResponse) {
	response = &CreateApplicationScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
