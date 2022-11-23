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

// ModifyDesensitizationStrategy invokes the dms_enterprise.ModifyDesensitizationStrategy API synchronously
func (client *Client) ModifyDesensitizationStrategy(request *ModifyDesensitizationStrategyRequest) (response *ModifyDesensitizationStrategyResponse, err error) {
	response = CreateModifyDesensitizationStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDesensitizationStrategyWithChan invokes the dms_enterprise.ModifyDesensitizationStrategy API asynchronously
func (client *Client) ModifyDesensitizationStrategyWithChan(request *ModifyDesensitizationStrategyRequest) (<-chan *ModifyDesensitizationStrategyResponse, <-chan error) {
	responseChan := make(chan *ModifyDesensitizationStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDesensitizationStrategy(request)
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

// ModifyDesensitizationStrategyWithCallback invokes the dms_enterprise.ModifyDesensitizationStrategy API asynchronously
func (client *Client) ModifyDesensitizationStrategyWithCallback(request *ModifyDesensitizationStrategyRequest, callback func(response *ModifyDesensitizationStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDesensitizationStrategyResponse
		var err error
		defer close(result)
		response, err = client.ModifyDesensitizationStrategy(request)
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

// ModifyDesensitizationStrategyRequest is the request struct for api ModifyDesensitizationStrategy
type ModifyDesensitizationStrategyRequest struct {
	*requests.RpcRequest
	IsReset    requests.Boolean `position:"Query" name:"IsReset"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	TableName  string           `position:"Query" name:"TableName"`
	SchemaName string           `position:"Query" name:"SchemaName"`
	IsLogic    requests.Boolean `position:"Query" name:"IsLogic"`
	ColumnName string           `position:"Query" name:"ColumnName"`
	DbId       requests.Integer `position:"Query" name:"DbId"`
	RuleId     requests.Integer `position:"Query" name:"RuleId"`
}

// ModifyDesensitizationStrategyResponse is the response struct for api ModifyDesensitizationStrategy
type ModifyDesensitizationStrategyResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Result       bool   `json:"Result" xml:"Result"`
}

// CreateModifyDesensitizationStrategyRequest creates a request to invoke ModifyDesensitizationStrategy API
func CreateModifyDesensitizationStrategyRequest() (request *ModifyDesensitizationStrategyRequest) {
	request = &ModifyDesensitizationStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ModifyDesensitizationStrategy", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDesensitizationStrategyResponse creates a response to parse from ModifyDesensitizationStrategy response
func CreateModifyDesensitizationStrategyResponse() (response *ModifyDesensitizationStrategyResponse) {
	response = &ModifyDesensitizationStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
