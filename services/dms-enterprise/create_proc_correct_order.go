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

// CreateProcCorrectOrder invokes the dms_enterprise.CreateProcCorrectOrder API synchronously
func (client *Client) CreateProcCorrectOrder(request *CreateProcCorrectOrderRequest) (response *CreateProcCorrectOrderResponse, err error) {
	response = CreateCreateProcCorrectOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProcCorrectOrderWithChan invokes the dms_enterprise.CreateProcCorrectOrder API asynchronously
func (client *Client) CreateProcCorrectOrderWithChan(request *CreateProcCorrectOrderRequest) (<-chan *CreateProcCorrectOrderResponse, <-chan error) {
	responseChan := make(chan *CreateProcCorrectOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProcCorrectOrder(request)
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

// CreateProcCorrectOrderWithCallback invokes the dms_enterprise.CreateProcCorrectOrder API asynchronously
func (client *Client) CreateProcCorrectOrderWithCallback(request *CreateProcCorrectOrderRequest, callback func(response *CreateProcCorrectOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProcCorrectOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateProcCorrectOrder(request)
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

// CreateProcCorrectOrderRequest is the request struct for api CreateProcCorrectOrder
type CreateProcCorrectOrderRequest struct {
	*requests.RpcRequest
	Tid             requests.Integer            `position:"Query" name:"Tid"`
	Param           CreateProcCorrectOrderParam `position:"Query" name:"Param"  type:"Struct"`
	RelatedUserList *[]string                   `position:"Query" name:"RelatedUserList"  type:"Json"`
	AttachmentKey   string                      `position:"Query" name:"AttachmentKey"`
	Comment         string                      `position:"Query" name:"Comment"`
}

// CreateProcCorrectOrderParam is a repeated param struct in CreateProcCorrectOrderRequest
type CreateProcCorrectOrderParam struct {
	Classify               string                                       `name:"Classify"`
	RollbackSQL            string                                       `name:"RollbackSQL"`
	RollbackSqlType        string                                       `name:"RollbackSqlType"`
	DbItemList             *[]CreateProcCorrectOrderParamDbItemListItem `name:"DbItemList" type:"Repeated"`
	ExecSQL                string                                       `name:"ExecSQL"`
	RollbackAttachmentName string                                       `name:"RollbackAttachmentName"`
}

// CreateProcCorrectOrderParamDbItemListItem is a repeated param struct in CreateProcCorrectOrderRequest
type CreateProcCorrectOrderParamDbItemListItem struct {
	DbId  string `name:"DbId"`
	Logic string `name:"Logic"`
}

// CreateProcCorrectOrderResponse is the response struct for api CreateProcCorrectOrder
type CreateProcCorrectOrderResponse struct {
	*responses.BaseResponse
	RequestId         string  `json:"RequestId" xml:"RequestId"`
	Success           bool    `json:"Success" xml:"Success"`
	ErrorMessage      string  `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode         string  `json:"ErrorCode" xml:"ErrorCode"`
	CreateOrderResult []int64 `json:"CreateOrderResult" xml:"CreateOrderResult"`
}

// CreateCreateProcCorrectOrderRequest creates a request to invoke CreateProcCorrectOrder API
func CreateCreateProcCorrectOrderRequest() (request *CreateProcCorrectOrderRequest) {
	request = &CreateProcCorrectOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateProcCorrectOrder", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProcCorrectOrderResponse creates a response to parse from CreateProcCorrectOrder response
func CreateCreateProcCorrectOrderResponse() (response *CreateProcCorrectOrderResponse) {
	response = &CreateProcCorrectOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
