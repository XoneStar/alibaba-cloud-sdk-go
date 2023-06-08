package resourcemanager

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

// ListTargetAttachmentsForControlPolicy invokes the resourcemanager.ListTargetAttachmentsForControlPolicy API synchronously
func (client *Client) ListTargetAttachmentsForControlPolicy(request *ListTargetAttachmentsForControlPolicyRequest) (response *ListTargetAttachmentsForControlPolicyResponse, err error) {
	response = CreateListTargetAttachmentsForControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ListTargetAttachmentsForControlPolicyWithChan invokes the resourcemanager.ListTargetAttachmentsForControlPolicy API asynchronously
func (client *Client) ListTargetAttachmentsForControlPolicyWithChan(request *ListTargetAttachmentsForControlPolicyRequest) (<-chan *ListTargetAttachmentsForControlPolicyResponse, <-chan error) {
	responseChan := make(chan *ListTargetAttachmentsForControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTargetAttachmentsForControlPolicy(request)
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

// ListTargetAttachmentsForControlPolicyWithCallback invokes the resourcemanager.ListTargetAttachmentsForControlPolicy API asynchronously
func (client *Client) ListTargetAttachmentsForControlPolicyWithCallback(request *ListTargetAttachmentsForControlPolicyRequest, callback func(response *ListTargetAttachmentsForControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTargetAttachmentsForControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.ListTargetAttachmentsForControlPolicy(request)
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

// ListTargetAttachmentsForControlPolicyRequest is the request struct for api ListTargetAttachmentsForControlPolicy
type ListTargetAttachmentsForControlPolicyRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PolicyId   string           `position:"Query" name:"PolicyId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListTargetAttachmentsForControlPolicyResponse is the response struct for api ListTargetAttachmentsForControlPolicy
type ListTargetAttachmentsForControlPolicyResponse struct {
	*responses.BaseResponse
	TotalCount        int               `json:"TotalCount" xml:"TotalCount"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	PageSize          int               `json:"PageSize" xml:"PageSize"`
	PageNumber        int               `json:"PageNumber" xml:"PageNumber"`
	TargetAttachments TargetAttachments `json:"TargetAttachments" xml:"TargetAttachments"`
}

// CreateListTargetAttachmentsForControlPolicyRequest creates a request to invoke ListTargetAttachmentsForControlPolicy API
func CreateListTargetAttachmentsForControlPolicyRequest() (request *ListTargetAttachmentsForControlPolicyRequest) {
	request = &ListTargetAttachmentsForControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ListTargetAttachmentsForControlPolicy", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTargetAttachmentsForControlPolicyResponse creates a response to parse from ListTargetAttachmentsForControlPolicy response
func CreateListTargetAttachmentsForControlPolicyResponse() (response *ListTargetAttachmentsForControlPolicyResponse) {
	response = &ListTargetAttachmentsForControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
