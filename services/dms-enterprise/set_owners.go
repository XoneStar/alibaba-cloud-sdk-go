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

// SetOwners invokes the dms_enterprise.SetOwners API synchronously
func (client *Client) SetOwners(request *SetOwnersRequest) (response *SetOwnersResponse, err error) {
	response = CreateSetOwnersResponse()
	err = client.DoAction(request, response)
	return
}

// SetOwnersWithChan invokes the dms_enterprise.SetOwners API asynchronously
func (client *Client) SetOwnersWithChan(request *SetOwnersRequest) (<-chan *SetOwnersResponse, <-chan error) {
	responseChan := make(chan *SetOwnersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetOwners(request)
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

// SetOwnersWithCallback invokes the dms_enterprise.SetOwners API asynchronously
func (client *Client) SetOwnersWithCallback(request *SetOwnersRequest, callback func(response *SetOwnersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetOwnersResponse
		var err error
		defer close(result)
		response, err = client.SetOwners(request)
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

// SetOwnersRequest is the request struct for api SetOwners
type SetOwnersRequest struct {
	*requests.RpcRequest
	OwnerIds   string           `position:"Query" name:"OwnerIds"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	ResourceId string           `position:"Query" name:"ResourceId"`
	OwnerType  string           `position:"Query" name:"OwnerType"`
}

// SetOwnersResponse is the response struct for api SetOwners
type SetOwnersResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateSetOwnersRequest creates a request to invoke SetOwners API
func CreateSetOwnersRequest() (request *SetOwnersRequest) {
	request = &SetOwnersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "SetOwners", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetOwnersResponse creates a response to parse from SetOwners response
func CreateSetOwnersResponse() (response *SetOwnersResponse) {
	response = &SetOwnersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
