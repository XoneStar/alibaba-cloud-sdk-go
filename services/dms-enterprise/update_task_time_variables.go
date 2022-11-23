package dms_enterprise

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

// UpdateTaskTimeVariables invokes the dms_enterprise.UpdateTaskTimeVariables API synchronously
func (client *Client) UpdateTaskTimeVariables(request *UpdateTaskTimeVariablesRequest) (response *UpdateTaskTimeVariablesResponse, err error) {
	response = CreateUpdateTaskTimeVariablesResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTaskTimeVariablesWithChan invokes the dms_enterprise.UpdateTaskTimeVariables API asynchronously
func (client *Client) UpdateTaskTimeVariablesWithChan(request *UpdateTaskTimeVariablesRequest) (<-chan *UpdateTaskTimeVariablesResponse, <-chan error) {
	responseChan := make(chan *UpdateTaskTimeVariablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTaskTimeVariables(request)
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

// UpdateTaskTimeVariablesWithCallback invokes the dms_enterprise.UpdateTaskTimeVariables API asynchronously
func (client *Client) UpdateTaskTimeVariablesWithCallback(request *UpdateTaskTimeVariablesRequest, callback func(response *UpdateTaskTimeVariablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTaskTimeVariablesResponse
		var err error
		defer close(result)
		response, err = client.UpdateTaskTimeVariables(request)
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

// UpdateTaskTimeVariablesRequest is the request struct for api UpdateTaskTimeVariables
type UpdateTaskTimeVariablesRequest struct {
	*requests.RpcRequest
	Tid           requests.Integer `position:"Query" name:"Tid"`
	NodeId        string           `position:"Query" name:"NodeId"`
	TimeVariables string           `position:"Query" name:"TimeVariables"`
}

// UpdateTaskTimeVariablesResponse is the response struct for api UpdateTaskTimeVariables
type UpdateTaskTimeVariablesResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateUpdateTaskTimeVariablesRequest creates a request to invoke UpdateTaskTimeVariables API
func CreateUpdateTaskTimeVariablesRequest() (request *UpdateTaskTimeVariablesRequest) {
	request = &UpdateTaskTimeVariablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "UpdateTaskTimeVariables", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTaskTimeVariablesResponse creates a response to parse from UpdateTaskTimeVariables response
func CreateUpdateTaskTimeVariablesResponse() (response *UpdateTaskTimeVariablesResponse) {
	response = &UpdateTaskTimeVariablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
