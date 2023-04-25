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

// ListInstancesOfUser invokes the ccc.ListInstancesOfUser API synchronously
func (client *Client) ListInstancesOfUser(request *ListInstancesOfUserRequest) (response *ListInstancesOfUserResponse, err error) {
	response = CreateListInstancesOfUserResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstancesOfUserWithChan invokes the ccc.ListInstancesOfUser API asynchronously
func (client *Client) ListInstancesOfUserWithChan(request *ListInstancesOfUserRequest) (<-chan *ListInstancesOfUserResponse, <-chan error) {
	responseChan := make(chan *ListInstancesOfUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstancesOfUser(request)
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

// ListInstancesOfUserWithCallback invokes the ccc.ListInstancesOfUser API asynchronously
func (client *Client) ListInstancesOfUserWithCallback(request *ListInstancesOfUserRequest, callback func(response *ListInstancesOfUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstancesOfUserResponse
		var err error
		defer close(result)
		response, err = client.ListInstancesOfUser(request)
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

// ListInstancesOfUserRequest is the request struct for api ListInstancesOfUser
type ListInstancesOfUserRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListInstancesOfUserResponse is the response struct for api ListInstancesOfUser
type ListInstancesOfUserResponse struct {
	*responses.BaseResponse
	Code           string                    `json:"Code" xml:"Code"`
	HttpStatusCode int                       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                    `json:"Message" xml:"Message"`
	RequestId      string                    `json:"RequestId" xml:"RequestId"`
	Data           DataInListInstancesOfUser `json:"Data" xml:"Data"`
}

// CreateListInstancesOfUserRequest creates a request to invoke ListInstancesOfUser API
func CreateListInstancesOfUserRequest() (request *ListInstancesOfUserRequest) {
	request = &ListInstancesOfUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListInstancesOfUser", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListInstancesOfUserResponse creates a response to parse from ListInstancesOfUser response
func CreateListInstancesOfUserResponse() (response *ListInstancesOfUserResponse) {
	response = &ListInstancesOfUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
