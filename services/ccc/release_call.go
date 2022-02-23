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

// ReleaseCall invokes the ccc.ReleaseCall API synchronously
func (client *Client) ReleaseCall(request *ReleaseCallRequest) (response *ReleaseCallResponse, err error) {
	response = CreateReleaseCallResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseCallWithChan invokes the ccc.ReleaseCall API asynchronously
func (client *Client) ReleaseCallWithChan(request *ReleaseCallRequest) (<-chan *ReleaseCallResponse, <-chan error) {
	responseChan := make(chan *ReleaseCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseCall(request)
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

// ReleaseCallWithCallback invokes the ccc.ReleaseCall API asynchronously
func (client *Client) ReleaseCallWithCallback(request *ReleaseCallRequest, callback func(response *ReleaseCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseCallResponse
		var err error
		defer close(result)
		response, err = client.ReleaseCall(request)
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

// ReleaseCallRequest is the request struct for api ReleaseCall
type ReleaseCallRequest struct {
	*requests.RpcRequest
	UserId     string `position:"Query" name:"UserId"`
	DeviceId   string `position:"Query" name:"DeviceId"`
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	ChannelId  string `position:"Query" name:"ChannelId"`
}

// ReleaseCallResponse is the response struct for api ReleaseCall
type ReleaseCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateReleaseCallRequest creates a request to invoke ReleaseCall API
func CreateReleaseCallRequest() (request *ReleaseCallRequest) {
	request = &ReleaseCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ReleaseCall", "", "")
	request.Method = requests.POST
	return
}

// CreateReleaseCallResponse creates a response to parse from ReleaseCall response
func CreateReleaseCallResponse() (response *ReleaseCallResponse) {
	response = &ReleaseCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}