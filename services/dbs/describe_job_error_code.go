package dbs

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

// DescribeJobErrorCode invokes the dbs.DescribeJobErrorCode API synchronously
func (client *Client) DescribeJobErrorCode(request *DescribeJobErrorCodeRequest) (response *DescribeJobErrorCodeResponse, err error) {
	response = CreateDescribeJobErrorCodeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeJobErrorCodeWithChan invokes the dbs.DescribeJobErrorCode API asynchronously
func (client *Client) DescribeJobErrorCodeWithChan(request *DescribeJobErrorCodeRequest) (<-chan *DescribeJobErrorCodeResponse, <-chan error) {
	responseChan := make(chan *DescribeJobErrorCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeJobErrorCode(request)
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

// DescribeJobErrorCodeWithCallback invokes the dbs.DescribeJobErrorCode API asynchronously
func (client *Client) DescribeJobErrorCodeWithCallback(request *DescribeJobErrorCodeRequest, callback func(response *DescribeJobErrorCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeJobErrorCodeResponse
		var err error
		defer close(result)
		response, err = client.DescribeJobErrorCode(request)
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

// DescribeJobErrorCodeRequest is the request struct for api DescribeJobErrorCode
type DescribeJobErrorCodeRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	Language    string `position:"Query" name:"Language"`
	TaskId      string `position:"Query" name:"TaskId"`
	OwnerId     string `position:"Query" name:"OwnerId"`
}

// DescribeJobErrorCodeResponse is the response struct for api DescribeJobErrorCode
type DescribeJobErrorCodeResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	Item           Item   `json:"Item" xml:"Item"`
}

// CreateDescribeJobErrorCodeRequest creates a request to invoke DescribeJobErrorCode API
func CreateDescribeJobErrorCodeRequest() (request *DescribeJobErrorCodeRequest) {
	request = &DescribeJobErrorCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribeJobErrorCode", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeJobErrorCodeResponse creates a response to parse from DescribeJobErrorCode response
func CreateDescribeJobErrorCodeResponse() (response *DescribeJobErrorCodeResponse) {
	response = &DescribeJobErrorCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
