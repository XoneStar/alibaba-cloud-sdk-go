package companyreg

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

// AppFailClear invokes the companyreg.AppFailClear API synchronously
func (client *Client) AppFailClear(request *AppFailClearRequest) (response *AppFailClearResponse, err error) {
	response = CreateAppFailClearResponse()
	err = client.DoAction(request, response)
	return
}

// AppFailClearWithChan invokes the companyreg.AppFailClear API asynchronously
func (client *Client) AppFailClearWithChan(request *AppFailClearRequest) (<-chan *AppFailClearResponse, <-chan error) {
	responseChan := make(chan *AppFailClearResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AppFailClear(request)
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

// AppFailClearWithCallback invokes the companyreg.AppFailClear API asynchronously
func (client *Client) AppFailClearWithCallback(request *AppFailClearRequest, callback func(response *AppFailClearResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AppFailClearResponse
		var err error
		defer close(result)
		response, err = client.AppFailClear(request)
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

// AppFailClearRequest is the request struct for api AppFailClear
type AppFailClearRequest struct {
	*requests.RpcRequest
	Period string `position:"Query" name:"Period"`
	BizId  string `position:"Query" name:"BizId"`
}

// AppFailClearResponse is the response struct for api AppFailClear
type AppFailClearResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAppFailClearRequest creates a request to invoke AppFailClear API
func CreateAppFailClearRequest() (request *AppFailClearRequest) {
	request = &AppFailClearRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "AppFailClear", "", "")
	request.Method = requests.GET
	return
}

// CreateAppFailClearResponse creates a response to parse from AppFailClear response
func CreateAppFailClearResponse() (response *AppFailClearResponse) {
	response = &AppFailClearResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}