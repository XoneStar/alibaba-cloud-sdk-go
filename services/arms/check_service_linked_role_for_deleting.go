package arms

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

// CheckServiceLinkedRoleForDeleting invokes the arms.CheckServiceLinkedRoleForDeleting API synchronously
func (client *Client) CheckServiceLinkedRoleForDeleting(request *CheckServiceLinkedRoleForDeletingRequest) (response *CheckServiceLinkedRoleForDeletingResponse, err error) {
	response = CreateCheckServiceLinkedRoleForDeletingResponse()
	err = client.DoAction(request, response)
	return
}

// CheckServiceLinkedRoleForDeletingWithChan invokes the arms.CheckServiceLinkedRoleForDeleting API asynchronously
func (client *Client) CheckServiceLinkedRoleForDeletingWithChan(request *CheckServiceLinkedRoleForDeletingRequest) (<-chan *CheckServiceLinkedRoleForDeletingResponse, <-chan error) {
	responseChan := make(chan *CheckServiceLinkedRoleForDeletingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckServiceLinkedRoleForDeleting(request)
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

// CheckServiceLinkedRoleForDeletingWithCallback invokes the arms.CheckServiceLinkedRoleForDeleting API asynchronously
func (client *Client) CheckServiceLinkedRoleForDeletingWithCallback(request *CheckServiceLinkedRoleForDeletingRequest, callback func(response *CheckServiceLinkedRoleForDeletingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckServiceLinkedRoleForDeletingResponse
		var err error
		defer close(result)
		response, err = client.CheckServiceLinkedRoleForDeleting(request)
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

// CheckServiceLinkedRoleForDeletingRequest is the request struct for api CheckServiceLinkedRoleForDeleting
type CheckServiceLinkedRoleForDeletingRequest struct {
	*requests.RpcRequest
	SPIRegionId    string `position:"Query" name:"SPIRegionId"`
	RoleArn        string `position:"Query" name:"RoleArn"`
	DeletionTaskId string `position:"Query" name:"DeletionTaskId"`
	ServiceName    string `position:"Query" name:"ServiceName"`
}

// CheckServiceLinkedRoleForDeletingResponse is the response struct for api CheckServiceLinkedRoleForDeleting
type CheckServiceLinkedRoleForDeletingResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	Deletable  bool             `json:"Deletable" xml:"Deletable"`
	RoleUsages []RoleUsagesItem `json:"RoleUsages" xml:"RoleUsages"`
}

// CreateCheckServiceLinkedRoleForDeletingRequest creates a request to invoke CheckServiceLinkedRoleForDeleting API
func CreateCheckServiceLinkedRoleForDeletingRequest() (request *CheckServiceLinkedRoleForDeletingRequest) {
	request = &CheckServiceLinkedRoleForDeletingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CheckServiceLinkedRoleForDeleting", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckServiceLinkedRoleForDeletingResponse creates a response to parse from CheckServiceLinkedRoleForDeleting response
func CreateCheckServiceLinkedRoleForDeletingResponse() (response *CheckServiceLinkedRoleForDeletingResponse) {
	response = &CheckServiceLinkedRoleForDeletingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
