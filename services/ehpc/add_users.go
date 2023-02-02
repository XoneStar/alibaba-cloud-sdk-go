package ehpc

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

// AddUsers invokes the ehpc.AddUsers API synchronously
func (client *Client) AddUsers(request *AddUsersRequest) (response *AddUsersResponse, err error) {
	response = CreateAddUsersResponse()
	err = client.DoAction(request, response)
	return
}

// AddUsersWithChan invokes the ehpc.AddUsers API asynchronously
func (client *Client) AddUsersWithChan(request *AddUsersRequest) (<-chan *AddUsersResponse, <-chan error) {
	responseChan := make(chan *AddUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddUsers(request)
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

// AddUsersWithCallback invokes the ehpc.AddUsers API asynchronously
func (client *Client) AddUsersWithCallback(request *AddUsersRequest, callback func(response *AddUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddUsersResponse
		var err error
		defer close(result)
		response, err = client.AddUsers(request)
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

// AddUsersRequest is the request struct for api AddUsers
type AddUsersRequest struct {
	*requests.RpcRequest
	ClusterId string           `position:"Query" name:"ClusterId"`
	Async     requests.Boolean `position:"Query" name:"Async"`
	User      *[]AddUsersUser  `position:"Query" name:"User"  type:"Repeated"`
}

// AddUsersUser is a repeated param struct in AddUsersRequest
type AddUsersUser struct {
	Password string `name:"Password"`
	Name     string `name:"Name"`
	Group    string `name:"Group"`
}

// AddUsersResponse is the response struct for api AddUsers
type AddUsersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddUsersRequest creates a request to invoke AddUsers API
func CreateAddUsersRequest() (request *AddUsersRequest) {
	request = &AddUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "AddUsers", "", "")
	request.Method = requests.GET
	return
}

// CreateAddUsersResponse creates a response to parse from AddUsers response
func CreateAddUsersResponse() (response *AddUsersResponse) {
	response = &AddUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
