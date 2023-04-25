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

// ListAudioFiles invokes the ccc.ListAudioFiles API synchronously
func (client *Client) ListAudioFiles(request *ListAudioFilesRequest) (response *ListAudioFilesResponse, err error) {
	response = CreateListAudioFilesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAudioFilesWithChan invokes the ccc.ListAudioFiles API asynchronously
func (client *Client) ListAudioFilesWithChan(request *ListAudioFilesRequest) (<-chan *ListAudioFilesResponse, <-chan error) {
	responseChan := make(chan *ListAudioFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAudioFiles(request)
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

// ListAudioFilesWithCallback invokes the ccc.ListAudioFiles API asynchronously
func (client *Client) ListAudioFilesWithCallback(request *ListAudioFilesRequest, callback func(response *ListAudioFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAudioFilesResponse
		var err error
		defer close(result)
		response, err = client.ListAudioFiles(request)
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

// ListAudioFilesRequest is the request struct for api ListAudioFiles
type ListAudioFilesRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListAudioFilesResponse is the response struct for api ListAudioFiles
type ListAudioFilesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string               `json:"RequestId" xml:"RequestId"`
	Code           string               `json:"Code" xml:"Code"`
	Message        string               `json:"Message" xml:"Message"`
	Data           DataInListAudioFiles `json:"Data" xml:"Data"`
}

// CreateListAudioFilesRequest creates a request to invoke ListAudioFiles API
func CreateListAudioFilesRequest() (request *ListAudioFilesRequest) {
	request = &ListAudioFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListAudioFiles", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAudioFilesResponse creates a response to parse from ListAudioFiles response
func CreateListAudioFilesResponse() (response *ListAudioFilesResponse) {
	response = &ListAudioFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
