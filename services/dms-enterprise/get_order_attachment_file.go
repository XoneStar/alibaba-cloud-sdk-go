package dms_enterprise

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

// GetOrderAttachmentFile invokes the dms_enterprise.GetOrderAttachmentFile API synchronously
func (client *Client) GetOrderAttachmentFile(request *GetOrderAttachmentFileRequest) (response *GetOrderAttachmentFileResponse, err error) {
	response = CreateGetOrderAttachmentFileResponse()
	err = client.DoAction(request, response)
	return
}

// GetOrderAttachmentFileWithChan invokes the dms_enterprise.GetOrderAttachmentFile API asynchronously
func (client *Client) GetOrderAttachmentFileWithChan(request *GetOrderAttachmentFileRequest) (<-chan *GetOrderAttachmentFileResponse, <-chan error) {
	responseChan := make(chan *GetOrderAttachmentFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOrderAttachmentFile(request)
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

// GetOrderAttachmentFileWithCallback invokes the dms_enterprise.GetOrderAttachmentFile API asynchronously
func (client *Client) GetOrderAttachmentFileWithCallback(request *GetOrderAttachmentFileRequest, callback func(response *GetOrderAttachmentFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOrderAttachmentFileResponse
		var err error
		defer close(result)
		response, err = client.GetOrderAttachmentFile(request)
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

// GetOrderAttachmentFileRequest is the request struct for api GetOrderAttachmentFile
type GetOrderAttachmentFileRequest struct {
	*requests.RpcRequest
	Tid     requests.Integer `position:"Query" name:"Tid"`
	OrderId requests.Integer `position:"Query" name:"OrderId"`
}

// GetOrderAttachmentFileResponse is the response struct for api GetOrderAttachmentFile
type GetOrderAttachmentFileResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	FileUrl      string `json:"FileUrl" xml:"FileUrl"`
}

// CreateGetOrderAttachmentFileRequest creates a request to invoke GetOrderAttachmentFile API
func CreateGetOrderAttachmentFileRequest() (request *GetOrderAttachmentFileRequest) {
	request = &GetOrderAttachmentFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetOrderAttachmentFile", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOrderAttachmentFileResponse creates a response to parse from GetOrderAttachmentFile response
func CreateGetOrderAttachmentFileResponse() (response *GetOrderAttachmentFileResponse) {
	response = &GetOrderAttachmentFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
