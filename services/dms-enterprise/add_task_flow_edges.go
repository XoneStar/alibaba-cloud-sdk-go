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

// AddTaskFlowEdges invokes the dms_enterprise.AddTaskFlowEdges API synchronously
func (client *Client) AddTaskFlowEdges(request *AddTaskFlowEdgesRequest) (response *AddTaskFlowEdgesResponse, err error) {
	response = CreateAddTaskFlowEdgesResponse()
	err = client.DoAction(request, response)
	return
}

// AddTaskFlowEdgesWithChan invokes the dms_enterprise.AddTaskFlowEdges API asynchronously
func (client *Client) AddTaskFlowEdgesWithChan(request *AddTaskFlowEdgesRequest) (<-chan *AddTaskFlowEdgesResponse, <-chan error) {
	responseChan := make(chan *AddTaskFlowEdgesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTaskFlowEdges(request)
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

// AddTaskFlowEdgesWithCallback invokes the dms_enterprise.AddTaskFlowEdges API asynchronously
func (client *Client) AddTaskFlowEdgesWithCallback(request *AddTaskFlowEdgesRequest, callback func(response *AddTaskFlowEdgesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTaskFlowEdgesResponse
		var err error
		defer close(result)
		response, err = client.AddTaskFlowEdges(request)
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

// AddTaskFlowEdgesRequest is the request struct for api AddTaskFlowEdges
type AddTaskFlowEdgesRequest struct {
	*requests.RpcRequest
	DagId requests.Integer         `position:"Query" name:"DagId"`
	Tid   requests.Integer         `position:"Query" name:"Tid"`
	Edges *[]AddTaskFlowEdgesEdges `position:"Query" name:"Edges"  type:"Json"`
}

// AddTaskFlowEdgesEdges is a repeated param struct in AddTaskFlowEdgesRequest
type AddTaskFlowEdgesEdges struct {
	NodeEnd  string `name:"NodeEnd"`
	NodeFrom string `name:"NodeFrom"`
}

// AddTaskFlowEdgesResponse is the response struct for api AddTaskFlowEdges
type AddTaskFlowEdgesResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateAddTaskFlowEdgesRequest creates a request to invoke AddTaskFlowEdges API
func CreateAddTaskFlowEdgesRequest() (request *AddTaskFlowEdgesRequest) {
	request = &AddTaskFlowEdgesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "AddTaskFlowEdges", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddTaskFlowEdgesResponse creates a response to parse from AddTaskFlowEdges response
func CreateAddTaskFlowEdgesResponse() (response *AddTaskFlowEdgesResponse) {
	response = &AddTaskFlowEdgesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
