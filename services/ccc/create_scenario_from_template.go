package ccc

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

// CreateScenarioFromTemplate invokes the ccc.CreateScenarioFromTemplate API synchronously
func (client *Client) CreateScenarioFromTemplate(request *CreateScenarioFromTemplateRequest) (response *CreateScenarioFromTemplateResponse, err error) {
	response = CreateCreateScenarioFromTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateScenarioFromTemplateWithChan invokes the ccc.CreateScenarioFromTemplate API asynchronously
func (client *Client) CreateScenarioFromTemplateWithChan(request *CreateScenarioFromTemplateRequest) (<-chan *CreateScenarioFromTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateScenarioFromTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScenarioFromTemplate(request)
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

// CreateScenarioFromTemplateWithCallback invokes the ccc.CreateScenarioFromTemplate API asynchronously
func (client *Client) CreateScenarioFromTemplateWithCallback(request *CreateScenarioFromTemplateRequest, callback func(response *CreateScenarioFromTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScenarioFromTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateScenarioFromTemplate(request)
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

// CreateScenarioFromTemplateRequest is the request struct for api CreateScenarioFromTemplate
type CreateScenarioFromTemplateRequest struct {
	*requests.RpcRequest
	Variables   string `position:"Query" name:"Variables"`
	Description string `position:"Query" name:"Description"`
	TemplateId  string `position:"Query" name:"TemplateId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	Name        string `position:"Query" name:"Name"`
}

// CreateScenarioFromTemplateResponse is the response struct for api CreateScenarioFromTemplate
type CreateScenarioFromTemplateResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Scenario       Scenario `json:"Scenario" xml:"Scenario"`
}

// CreateCreateScenarioFromTemplateRequest creates a request to invoke CreateScenarioFromTemplate API
func CreateCreateScenarioFromTemplateRequest() (request *CreateScenarioFromTemplateRequest) {
	request = &CreateScenarioFromTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "CreateScenarioFromTemplate", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateScenarioFromTemplateResponse creates a response to parse from CreateScenarioFromTemplate response
func CreateCreateScenarioFromTemplateResponse() (response *CreateScenarioFromTemplateResponse) {
	response = &CreateScenarioFromTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}