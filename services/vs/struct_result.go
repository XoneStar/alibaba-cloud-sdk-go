package vs

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

// Result is a nested struct in vs response
type Result struct {
	Region           string   `json:"Region" xml:"Region"`
	TemplateId       string   `json:"TemplateId" xml:"TemplateId"`
	Name             string   `json:"Name" xml:"Name"`
	DirectoryId      string   `json:"DirectoryId" xml:"DirectoryId"`
	Error            string   `json:"Error" xml:"Error"`
	InstanceType     string   `json:"InstanceType" xml:"InstanceType"`
	ParentPlatformId string   `json:"ParentPlatformId" xml:"ParentPlatformId"`
	InstanceId       string   `json:"InstanceId" xml:"InstanceId"`
	DeviceId         string   `json:"DeviceId" xml:"DeviceId"`
	GroupId          string   `json:"GroupId" xml:"GroupId"`
	Id               string   `json:"Id" xml:"Id"`
	TemplateType     string   `json:"TemplateType" xml:"TemplateType"`
	Streams          []Stream `json:"Streams" xml:"Streams"`
}