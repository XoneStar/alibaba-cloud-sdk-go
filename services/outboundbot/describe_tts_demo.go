package outboundbot

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

// DescribeTTSDemo invokes the outboundbot.DescribeTTSDemo API synchronously
func (client *Client) DescribeTTSDemo(request *DescribeTTSDemoRequest) (response *DescribeTTSDemoResponse, err error) {
	response = CreateDescribeTTSDemoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTTSDemoWithChan invokes the outboundbot.DescribeTTSDemo API asynchronously
func (client *Client) DescribeTTSDemoWithChan(request *DescribeTTSDemoRequest) (<-chan *DescribeTTSDemoResponse, <-chan error) {
	responseChan := make(chan *DescribeTTSDemoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTTSDemo(request)
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

// DescribeTTSDemoWithCallback invokes the outboundbot.DescribeTTSDemo API asynchronously
func (client *Client) DescribeTTSDemoWithCallback(request *DescribeTTSDemoRequest, callback func(response *DescribeTTSDemoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTTSDemoResponse
		var err error
		defer close(result)
		response, err = client.DescribeTTSDemo(request)
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

// DescribeTTSDemoRequest is the request struct for api DescribeTTSDemo
type DescribeTTSDemoRequest struct {
	*requests.RpcRequest
	Voice      string           `position:"Query" name:"Voice"`
	Volume     requests.Integer `position:"Query" name:"Volume"`
	ScriptId   string           `position:"Query" name:"ScriptId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	Text       string           `position:"Query" name:"Text"`
	SpeechRate requests.Integer `position:"Query" name:"SpeechRate"`
	PitchRate  requests.Integer `position:"Query" name:"PitchRate"`
}

// DescribeTTSDemoResponse is the response struct for api DescribeTTSDemo
type DescribeTTSDemoResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	AuditionUrl    string `json:"AuditionUrl" xml:"AuditionUrl"`
}

// CreateDescribeTTSDemoRequest creates a request to invoke DescribeTTSDemo API
func CreateDescribeTTSDemoRequest() (request *DescribeTTSDemoRequest) {
	request = &DescribeTTSDemoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DescribeTTSDemo", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTTSDemoResponse creates a response to parse from DescribeTTSDemo response
func CreateDescribeTTSDemoResponse() (response *DescribeTTSDemoResponse) {
	response = &DescribeTTSDemoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
