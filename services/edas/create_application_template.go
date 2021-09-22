package edas

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

// CreateApplicationTemplate invokes the edas.CreateApplicationTemplate API synchronously
func (client *Client) CreateApplicationTemplate(request *CreateApplicationTemplateRequest) (response *CreateApplicationTemplateResponse, err error) {
	response = CreateCreateApplicationTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateApplicationTemplateWithChan invokes the edas.CreateApplicationTemplate API asynchronously
func (client *Client) CreateApplicationTemplateWithChan(request *CreateApplicationTemplateRequest) (<-chan *CreateApplicationTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateApplicationTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApplicationTemplate(request)
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

// CreateApplicationTemplateWithCallback invokes the edas.CreateApplicationTemplate API asynchronously
func (client *Client) CreateApplicationTemplateWithCallback(request *CreateApplicationTemplateRequest, callback func(response *CreateApplicationTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateApplicationTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateApplicationTemplate(request)
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

// CreateApplicationTemplateRequest is the request struct for api CreateApplicationTemplate
type CreateApplicationTemplateRequest struct {
	*requests.RoaRequest
	NasId              string           `position:"Body" name:"NasId"`
	EnableAhas         requests.Boolean `position:"Body" name:"EnableAhas"`
	SlsConfigs         string           `position:"Body" name:"SlsConfigs"`
	CommandArgs        string           `position:"Body" name:"CommandArgs"`
	Readiness          string           `position:"Body" name:"Readiness"`
	Liveness           string           `position:"Body" name:"Liveness"`
	Description        string           `position:"Body" name:"Description"`
	Envs               string           `position:"Body" name:"Envs"`
	EnvFroms           string           `position:"Body" name:"EnvFroms"`
	RequestCpu         string           `position:"Body" name:"RequestCpu"`
	RequestMem         string           `position:"Body" name:"RequestMem"`
	ShowName           string           `position:"Body" name:"ShowName"`
	LimitMem           string           `position:"Body" name:"LimitMem"`
	ConfigMountDescs   string           `position:"Body" name:"ConfigMountDescs"`
	DeployAcrossZones  requests.Boolean `position:"Body" name:"DeployAcrossZones"`
	DeployAcrossNodes  requests.Boolean `position:"Body" name:"DeployAcrossNodes"`
	PreStop            string           `position:"Body" name:"PreStop"`
	Replicas           requests.Integer `position:"Body" name:"Replicas"`
	LimitCpu           string           `position:"Body" name:"LimitCpu"`
	WebContainerConfig string           `position:"Body" name:"WebContainerConfig"`
	PackageConfig      string           `position:"Body" name:"PackageConfig"`
	IsMultilingualApp  requests.Boolean `position:"Body" name:"IsMultilingualApp"`
	NasMountDescs      string           `position:"Body" name:"NasMountDescs"`
	LocalVolumes       string           `position:"Body" name:"LocalVolumes"`
	Command            string           `position:"Body" name:"Command"`
	NasStorageType     string           `position:"Body" name:"NasStorageType"`
	ImageConfig        string           `position:"Body" name:"ImageConfig"`
	SourceConfig       string           `position:"Body" name:"SourceConfig"`
	EmptyDirs          string           `position:"Body" name:"EmptyDirs"`
	PvcMountDescs      string           `position:"Body" name:"PvcMountDescs"`
	Name               string           `position:"Body" name:"Name"`
	Attributes         string           `position:"Body" name:"Attributes"`
	RuntimeClassName   string           `position:"Body" name:"RuntimeClassName"`
	JavaStartUpConfig  string           `position:"Body" name:"JavaStartUpConfig"`
	PostStart          string           `position:"Body" name:"PostStart"`
}

// CreateApplicationTemplateResponse is the response struct for api CreateApplicationTemplate
type CreateApplicationTemplateResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateApplicationTemplateRequest creates a request to invoke CreateApplicationTemplate API
func CreateCreateApplicationTemplateRequest() (request *CreateApplicationTemplateRequest) {
	request = &CreateApplicationTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "CreateApplicationTemplate", "/pop/v5/cnedas/app_template", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateApplicationTemplateResponse creates a response to parse from CreateApplicationTemplate response
func CreateCreateApplicationTemplateResponse() (response *CreateApplicationTemplateResponse) {
	response = &CreateApplicationTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
