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

// ListSkillGroupStates invokes the ccc.ListSkillGroupStates API synchronously
func (client *Client) ListSkillGroupStates(request *ListSkillGroupStatesRequest) (response *ListSkillGroupStatesResponse, err error) {
	response = CreateListSkillGroupStatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListSkillGroupStatesWithChan invokes the ccc.ListSkillGroupStates API asynchronously
func (client *Client) ListSkillGroupStatesWithChan(request *ListSkillGroupStatesRequest) (<-chan *ListSkillGroupStatesResponse, <-chan error) {
	responseChan := make(chan *ListSkillGroupStatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSkillGroupStates(request)
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

// ListSkillGroupStatesWithCallback invokes the ccc.ListSkillGroupStates API asynchronously
func (client *Client) ListSkillGroupStatesWithCallback(request *ListSkillGroupStatesRequest, callback func(response *ListSkillGroupStatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSkillGroupStatesResponse
		var err error
		defer close(result)
		response, err = client.ListSkillGroupStates(request)
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

// ListSkillGroupStatesRequest is the request struct for api ListSkillGroupStates
type ListSkillGroupStatesRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	SkillGroupIds string           `position:"Query" name:"SkillGroupIds"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListSkillGroupStatesResponse is the response struct for api ListSkillGroupStates
type ListSkillGroupStatesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string                     `json:"Code" xml:"Code"`
	Message        string                     `json:"Message" xml:"Message"`
	RequestId      string                     `json:"RequestId" xml:"RequestId"`
	Success        bool                       `json:"Success" xml:"Success"`
	Data           DataInListSkillGroupStates `json:"Data" xml:"Data"`
}

// CreateListSkillGroupStatesRequest creates a request to invoke ListSkillGroupStates API
func CreateListSkillGroupStatesRequest() (request *ListSkillGroupStatesRequest) {
	request = &ListSkillGroupStatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListSkillGroupStates", "CCC", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListSkillGroupStatesResponse creates a response to parse from ListSkillGroupStates response
func CreateListSkillGroupStatesResponse() (response *ListSkillGroupStatesResponse) {
	response = &ListSkillGroupStatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
