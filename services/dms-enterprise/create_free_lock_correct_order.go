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

// CreateFreeLockCorrectOrder invokes the dms_enterprise.CreateFreeLockCorrectOrder API synchronously
func (client *Client) CreateFreeLockCorrectOrder(request *CreateFreeLockCorrectOrderRequest) (response *CreateFreeLockCorrectOrderResponse, err error) {
	response = CreateCreateFreeLockCorrectOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFreeLockCorrectOrderWithChan invokes the dms_enterprise.CreateFreeLockCorrectOrder API asynchronously
func (client *Client) CreateFreeLockCorrectOrderWithChan(request *CreateFreeLockCorrectOrderRequest) (<-chan *CreateFreeLockCorrectOrderResponse, <-chan error) {
	responseChan := make(chan *CreateFreeLockCorrectOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFreeLockCorrectOrder(request)
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

// CreateFreeLockCorrectOrderWithCallback invokes the dms_enterprise.CreateFreeLockCorrectOrder API asynchronously
func (client *Client) CreateFreeLockCorrectOrderWithCallback(request *CreateFreeLockCorrectOrderRequest, callback func(response *CreateFreeLockCorrectOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFreeLockCorrectOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateFreeLockCorrectOrder(request)
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

// CreateFreeLockCorrectOrderRequest is the request struct for api CreateFreeLockCorrectOrder
type CreateFreeLockCorrectOrderRequest struct {
	*requests.RpcRequest
	Tid             requests.Integer                `position:"Query" name:"Tid"`
	Param           CreateFreeLockCorrectOrderParam `position:"Query" name:"Param"  type:"Struct"`
	RelatedUserList *[]string                       `position:"Query" name:"RelatedUserList"  type:"Json"`
	AttachmentKey   string                          `position:"Query" name:"AttachmentKey"`
	Comment         string                          `position:"Query" name:"Comment"`
}

// CreateFreeLockCorrectOrderParam is a repeated param struct in CreateFreeLockCorrectOrderRequest
type CreateFreeLockCorrectOrderParam struct {
	SqlType                string                                           `name:"SqlType"`
	Classify               string                                           `name:"Classify"`
	RollbackSQL            string                                           `name:"RollbackSQL"`
	RollbackSqlType        string                                           `name:"RollbackSqlType"`
	DbItemList             *[]CreateFreeLockCorrectOrderParamDbItemListItem `name:"DbItemList" type:"Repeated"`
	ExecSQL                string                                           `name:"ExecSQL"`
	ExecMode               string                                           `name:"ExecMode"`
	RollbackAttachmentName string                                           `name:"RollbackAttachmentName"`
	AttachmentName         string                                           `name:"AttachmentName"`
}

// CreateFreeLockCorrectOrderParamDbItemListItem is a repeated param struct in CreateFreeLockCorrectOrderRequest
type CreateFreeLockCorrectOrderParamDbItemListItem struct {
	DbId  string `name:"DbId"`
	Logic string `name:"Logic"`
}

// CreateFreeLockCorrectOrderResponse is the response struct for api CreateFreeLockCorrectOrder
type CreateFreeLockCorrectOrderResponse struct {
	*responses.BaseResponse
	RequestId         string  `json:"RequestId" xml:"RequestId"`
	Success           bool    `json:"Success" xml:"Success"`
	ErrorMessage      string  `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode         string  `json:"ErrorCode" xml:"ErrorCode"`
	CreateOrderResult []int64 `json:"CreateOrderResult" xml:"CreateOrderResult"`
}

// CreateCreateFreeLockCorrectOrderRequest creates a request to invoke CreateFreeLockCorrectOrder API
func CreateCreateFreeLockCorrectOrderRequest() (request *CreateFreeLockCorrectOrderRequest) {
	request = &CreateFreeLockCorrectOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateFreeLockCorrectOrder", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFreeLockCorrectOrderResponse creates a response to parse from CreateFreeLockCorrectOrder response
func CreateCreateFreeLockCorrectOrderResponse() (response *CreateFreeLockCorrectOrderResponse) {
	response = &CreateFreeLockCorrectOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
