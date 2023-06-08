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

// DetachPolicy invokes the resourcemanager.DetachPolicy API synchronously
func (client *Client) DetachPolicy(request *DetachPolicyRequest) (response *DetachPolicyResponse, err error) {
	response = CreateDetachPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DetachPolicyWithChan invokes the resourcemanager.DetachPolicy API asynchronously
func (client *Client) DetachPolicyWithChan(request *DetachPolicyRequest) (<-chan *DetachPolicyResponse, <-chan error) {
	responseChan := make(chan *DetachPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachPolicy(request)
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

// DetachPolicyWithCallback invokes the resourcemanager.DetachPolicy API asynchronously
func (client *Client) DetachPolicyWithCallback(request *DetachPolicyRequest, callback func(response *DetachPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachPolicyResponse
		var err error
		defer close(result)
		response, err = client.DetachPolicy(request)
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

// DetachPolicyRequest is the request struct for api DetachPolicy
type DetachPolicyRequest struct {
	*requests.RpcRequest
	PolicyType      string `position:"Query" name:"PolicyType"`
	PrincipalType   string `position:"Query" name:"PrincipalType"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	PolicyName      string `position:"Query" name:"PolicyName"`
	PrincipalName   string `position:"Query" name:"PrincipalName"`
}

// DetachPolicyResponse is the response struct for api DetachPolicy
type DetachPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachPolicyRequest creates a request to invoke DetachPolicy API
func CreateDetachPolicyRequest() (request *DetachPolicyRequest) {
	request = &DetachPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "DetachPolicy", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachPolicyResponse creates a response to parse from DetachPolicy response
func CreateDetachPolicyResponse() (response *DetachPolicyResponse) {
	response = &DetachPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
