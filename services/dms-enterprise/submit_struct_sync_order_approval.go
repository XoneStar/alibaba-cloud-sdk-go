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

// SubmitStructSyncOrderApproval invokes the dms_enterprise.SubmitStructSyncOrderApproval API synchronously
func (client *Client) SubmitStructSyncOrderApproval(request *SubmitStructSyncOrderApprovalRequest) (response *SubmitStructSyncOrderApprovalResponse, err error) {
	response = CreateSubmitStructSyncOrderApprovalResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitStructSyncOrderApprovalWithChan invokes the dms_enterprise.SubmitStructSyncOrderApproval API asynchronously
func (client *Client) SubmitStructSyncOrderApprovalWithChan(request *SubmitStructSyncOrderApprovalRequest) (<-chan *SubmitStructSyncOrderApprovalResponse, <-chan error) {
	responseChan := make(chan *SubmitStructSyncOrderApprovalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitStructSyncOrderApproval(request)
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

// SubmitStructSyncOrderApprovalWithCallback invokes the dms_enterprise.SubmitStructSyncOrderApproval API asynchronously
func (client *Client) SubmitStructSyncOrderApprovalWithCallback(request *SubmitStructSyncOrderApprovalRequest, callback func(response *SubmitStructSyncOrderApprovalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitStructSyncOrderApprovalResponse
		var err error
		defer close(result)
		response, err = client.SubmitStructSyncOrderApproval(request)
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

// SubmitStructSyncOrderApprovalRequest is the request struct for api SubmitStructSyncOrderApproval
type SubmitStructSyncOrderApprovalRequest struct {
	*requests.RpcRequest
	Tid     requests.Integer `position:"Query" name:"Tid"`
	OrderId requests.Integer `position:"Query" name:"OrderId"`
}

// SubmitStructSyncOrderApprovalResponse is the response struct for api SubmitStructSyncOrderApproval
type SubmitStructSyncOrderApprovalResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	Success            bool   `json:"Success" xml:"Success"`
	ErrorMessage       string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode          string `json:"ErrorCode" xml:"ErrorCode"`
	WorkflowInstanceId int64  `json:"WorkflowInstanceId" xml:"WorkflowInstanceId"`
}

// CreateSubmitStructSyncOrderApprovalRequest creates a request to invoke SubmitStructSyncOrderApproval API
func CreateSubmitStructSyncOrderApprovalRequest() (request *SubmitStructSyncOrderApprovalRequest) {
	request = &SubmitStructSyncOrderApprovalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "SubmitStructSyncOrderApproval", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitStructSyncOrderApprovalResponse creates a response to parse from SubmitStructSyncOrderApproval response
func CreateSubmitStructSyncOrderApprovalResponse() (response *SubmitStructSyncOrderApprovalResponse) {
	response = &SubmitStructSyncOrderApprovalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
