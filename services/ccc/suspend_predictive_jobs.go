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

// SuspendPredictiveJobs invokes the ccc.SuspendPredictiveJobs API synchronously
func (client *Client) SuspendPredictiveJobs(request *SuspendPredictiveJobsRequest) (response *SuspendPredictiveJobsResponse, err error) {
	response = CreateSuspendPredictiveJobsResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendPredictiveJobsWithChan invokes the ccc.SuspendPredictiveJobs API asynchronously
func (client *Client) SuspendPredictiveJobsWithChan(request *SuspendPredictiveJobsRequest) (<-chan *SuspendPredictiveJobsResponse, <-chan error) {
	responseChan := make(chan *SuspendPredictiveJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendPredictiveJobs(request)
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

// SuspendPredictiveJobsWithCallback invokes the ccc.SuspendPredictiveJobs API asynchronously
func (client *Client) SuspendPredictiveJobsWithCallback(request *SuspendPredictiveJobsRequest, callback func(response *SuspendPredictiveJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendPredictiveJobsResponse
		var err error
		defer close(result)
		response, err = client.SuspendPredictiveJobs(request)
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

// SuspendPredictiveJobsRequest is the request struct for api SuspendPredictiveJobs
type SuspendPredictiveJobsRequest struct {
	*requests.RpcRequest
	All          requests.Boolean `position:"Query" name:"All"`
	JobId        *[]string        `position:"Query" name:"JobId"  type:"Repeated"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	JobGroupId   string           `position:"Query" name:"JobGroupId"`
}

// SuspendPredictiveJobsResponse is the response struct for api SuspendPredictiveJobs
type SuspendPredictiveJobsResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateSuspendPredictiveJobsRequest creates a request to invoke SuspendPredictiveJobs API
func CreateSuspendPredictiveJobsRequest() (request *SuspendPredictiveJobsRequest) {
	request = &SuspendPredictiveJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "SuspendPredictiveJobs", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendPredictiveJobsResponse creates a response to parse from SuspendPredictiveJobs response
func CreateSuspendPredictiveJobsResponse() (response *SuspendPredictiveJobsResponse) {
	response = &SuspendPredictiveJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}