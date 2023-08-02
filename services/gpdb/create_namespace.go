package gpdb

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

// CreateNamespace invokes the gpdb.CreateNamespace API synchronously
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
	response = CreateCreateNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNamespaceWithChan invokes the gpdb.CreateNamespace API asynchronously
func (client *Client) CreateNamespaceWithChan(request *CreateNamespaceRequest) (<-chan *CreateNamespaceResponse, <-chan error) {
	responseChan := make(chan *CreateNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNamespace(request)
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

// CreateNamespaceWithCallback invokes the gpdb.CreateNamespace API asynchronously
func (client *Client) CreateNamespaceWithCallback(request *CreateNamespaceRequest, callback func(response *CreateNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNamespaceResponse
		var err error
		defer close(result)
		response, err = client.CreateNamespace(request)
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

// CreateNamespaceRequest is the request struct for api CreateNamespace
type CreateNamespaceRequest struct {
	*requests.RpcRequest
	ManagerAccount         string           `position:"Query" name:"ManagerAccount"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ManagerAccountPassword string           `position:"Query" name:"ManagerAccountPassword"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	NamespacePassword      string           `position:"Query" name:"NamespacePassword"`
	Namespace              string           `position:"Query" name:"Namespace"`
}

// CreateNamespaceResponse is the response struct for api CreateNamespace
type CreateNamespaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateCreateNamespaceRequest creates a request to invoke CreateNamespace API
func CreateCreateNamespaceRequest() (request *CreateNamespaceRequest) {
	request = &CreateNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "CreateNamespace", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateNamespaceResponse creates a response to parse from CreateNamespace response
func CreateCreateNamespaceResponse() (response *CreateNamespaceResponse) {
	response = &CreateNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
