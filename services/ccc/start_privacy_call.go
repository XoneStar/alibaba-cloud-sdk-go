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

// StartPrivacyCall invokes the ccc.StartPrivacyCall API synchronously
func (client *Client) StartPrivacyCall(request *StartPrivacyCallRequest) (response *StartPrivacyCallResponse, err error) {
	response = CreateStartPrivacyCallResponse()
	err = client.DoAction(request, response)
	return
}

// StartPrivacyCallWithChan invokes the ccc.StartPrivacyCall API asynchronously
func (client *Client) StartPrivacyCallWithChan(request *StartPrivacyCallRequest) (<-chan *StartPrivacyCallResponse, <-chan error) {
	responseChan := make(chan *StartPrivacyCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartPrivacyCall(request)
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

// StartPrivacyCallWithCallback invokes the ccc.StartPrivacyCall API asynchronously
func (client *Client) StartPrivacyCallWithCallback(request *StartPrivacyCallRequest, callback func(response *StartPrivacyCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartPrivacyCallResponse
		var err error
		defer close(result)
		response, err = client.StartPrivacyCall(request)
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

// StartPrivacyCallRequest is the request struct for api StartPrivacyCall
type StartPrivacyCallRequest struct {
	*requests.RpcRequest
	Callee     string `position:"Query" name:"Callee"`
	Caller     string `position:"Query" name:"Caller"`
	InstanceId string `position:"Query" name:"InstanceId"`
	AppId      string `position:"Query" name:"AppId"`
}

// StartPrivacyCallResponse is the response struct for api StartPrivacyCall
type StartPrivacyCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Data           string   `json:"Data" xml:"Data"`
	Params         []string `json:"Params" xml:"Params"`
}

// CreateStartPrivacyCallRequest creates a request to invoke StartPrivacyCall API
func CreateStartPrivacyCallRequest() (request *StartPrivacyCallRequest) {
	request = &StartPrivacyCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "StartPrivacyCall", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartPrivacyCallResponse creates a response to parse from StartPrivacyCall response
func CreateStartPrivacyCallResponse() (response *StartPrivacyCallResponse) {
	response = &StartPrivacyCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
