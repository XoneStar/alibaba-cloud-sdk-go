package eas

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

// GetOrCreateResource invokes the eas.GetOrCreateResource API synchronously
// api document: https://help.aliyun.com/api/eas/getorcreateresource.html
func (client *Client) GetOrCreateResource(request *GetOrCreateResourceRequest) (response *GetOrCreateResourceResponse, err error) {
	response = CreateGetOrCreateResourceResponse()
	err = client.DoAction(request, response)
	return
}

// GetOrCreateResourceWithChan invokes the eas.GetOrCreateResource API asynchronously
// api document: https://help.aliyun.com/api/eas/getorcreateresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOrCreateResourceWithChan(request *GetOrCreateResourceRequest) (<-chan *GetOrCreateResourceResponse, <-chan error) {
	responseChan := make(chan *GetOrCreateResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOrCreateResource(request)
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

// GetOrCreateResourceWithCallback invokes the eas.GetOrCreateResource API asynchronously
// api document: https://help.aliyun.com/api/eas/getorcreateresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOrCreateResourceWithCallback(request *GetOrCreateResourceRequest, callback func(response *GetOrCreateResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOrCreateResourceResponse
		var err error
		defer close(result)
		response, err = client.GetOrCreateResource(request)
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

// GetOrCreateResourceRequest is the request struct for api GetOrCreateResource
type GetOrCreateResourceRequest struct {
	*requests.RoaRequest
}

// GetOrCreateResourceResponse is the response struct for api GetOrCreateResource
type GetOrCreateResourceResponse struct {
	*responses.BaseResponse
}

// CreateGetOrCreateResourceRequest creates a request to invoke GetOrCreateResource API
func CreateGetOrCreateResourceRequest() (request *GetOrCreateResourceRequest) {
	request = &GetOrCreateResourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "GetOrCreateResource", "/api/resources", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOrCreateResourceResponse creates a response to parse from GetOrCreateResource response
func CreateGetOrCreateResourceResponse() (response *GetOrCreateResourceResponse) {
	response = &GetOrCreateResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}