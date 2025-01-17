package dypnsapi

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

// GetPhoneWithToken invokes the dypnsapi.GetPhoneWithToken API synchronously
func (client *Client) GetPhoneWithToken(request *GetPhoneWithTokenRequest) (response *GetPhoneWithTokenResponse, err error) {
	response = CreateGetPhoneWithTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetPhoneWithTokenWithChan invokes the dypnsapi.GetPhoneWithToken API asynchronously
func (client *Client) GetPhoneWithTokenWithChan(request *GetPhoneWithTokenRequest) (<-chan *GetPhoneWithTokenResponse, <-chan error) {
	responseChan := make(chan *GetPhoneWithTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPhoneWithToken(request)
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

// GetPhoneWithTokenWithCallback invokes the dypnsapi.GetPhoneWithToken API asynchronously
func (client *Client) GetPhoneWithTokenWithCallback(request *GetPhoneWithTokenRequest, callback func(response *GetPhoneWithTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPhoneWithTokenResponse
		var err error
		defer close(result)
		response, err = client.GetPhoneWithToken(request)
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

// GetPhoneWithTokenRequest is the request struct for api GetPhoneWithToken
type GetPhoneWithTokenRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ProductCode          string           `position:"Query" name:"ProductCode"`
	SpToken              string           `position:"Query" name:"SpToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// GetPhoneWithTokenResponse is the response struct for api GetPhoneWithToken
type GetPhoneWithTokenResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetPhoneWithTokenRequest creates a request to invoke GetPhoneWithToken API
func CreateGetPhoneWithTokenRequest() (request *GetPhoneWithTokenRequest) {
	request = &GetPhoneWithTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "GetPhoneWithToken", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPhoneWithTokenResponse creates a response to parse from GetPhoneWithToken response
func CreateGetPhoneWithTokenResponse() (response *GetPhoneWithTokenResponse) {
	response = &GetPhoneWithTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
