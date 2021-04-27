package smc

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

// CutOverReplicationJob invokes the smc.CutOverReplicationJob API synchronously
func (client *Client) CutOverReplicationJob(request *CutOverReplicationJobRequest) (response *CutOverReplicationJobResponse, err error) {
	response = CreateCutOverReplicationJobResponse()
	err = client.DoAction(request, response)
	return
}

// CutOverReplicationJobWithChan invokes the smc.CutOverReplicationJob API asynchronously
func (client *Client) CutOverReplicationJobWithChan(request *CutOverReplicationJobRequest) (<-chan *CutOverReplicationJobResponse, <-chan error) {
	responseChan := make(chan *CutOverReplicationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CutOverReplicationJob(request)
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

// CutOverReplicationJobWithCallback invokes the smc.CutOverReplicationJob API asynchronously
func (client *Client) CutOverReplicationJobWithCallback(request *CutOverReplicationJobRequest, callback func(response *CutOverReplicationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CutOverReplicationJobResponse
		var err error
		defer close(result)
		response, err = client.CutOverReplicationJob(request)
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

// CutOverReplicationJobRequest is the request struct for api CutOverReplicationJob
type CutOverReplicationJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
	SyncData             requests.Boolean `position:"Query" name:"SyncData"`
}

// CutOverReplicationJobResponse is the response struct for api CutOverReplicationJob
type CutOverReplicationJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCutOverReplicationJobRequest creates a request to invoke CutOverReplicationJob API
func CreateCutOverReplicationJobRequest() (request *CutOverReplicationJobRequest) {
	request = &CutOverReplicationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("smc", "2019-06-01", "CutOverReplicationJob", "smc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCutOverReplicationJobResponse creates a response to parse from CutOverReplicationJob response
func CreateCutOverReplicationJobResponse() (response *CutOverReplicationJobResponse) {
	response = &CutOverReplicationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}