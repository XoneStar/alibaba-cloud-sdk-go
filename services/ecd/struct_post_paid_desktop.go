package ecd

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

// PostPaidDesktop is a nested struct in ecd response
type PostPaidDesktop struct {
	CreateDuration   string   `json:"CreateDuration" xml:"CreateDuration"`
	EndUserId        string   `json:"EndUserId" xml:"EndUserId"`
	DesktopStatus    string   `json:"DesktopStatus" xml:"DesktopStatus"`
	CreateTime       string   `json:"CreateTime" xml:"CreateTime"`
	ReleaseTime      string   `json:"ReleaseTime" xml:"ReleaseTime"`
	DesktopName      string   `json:"DesktopName" xml:"DesktopName"`
	ConnectionStatus string   `json:"ConnectionStatus" xml:"ConnectionStatus"`
	DesktopId        string   `json:"DesktopId" xml:"DesktopId"`
	EndUserName      string   `json:"EndUserName" xml:"EndUserName"`
	ManagementFlag   string   `json:"ManagementFlag" xml:"ManagementFlag"`
	ResetTime        string   `json:"ResetTime" xml:"ResetTime"`
	ImageId          string   `json:"ImageId" xml:"ImageId"`
	ImageName        string   `json:"ImageName" xml:"ImageName"`
	SystemDiskSize   int      `json:"SystemDiskSize" xml:"SystemDiskSize"`
	OsType           string   `json:"OsType" xml:"OsType"`
	GpuDriverVersion string   `json:"GpuDriverVersion" xml:"GpuDriverVersion"`
	DiskType         string   `json:"DiskType" xml:"DiskType"`
	EndUserIds       []string `json:"EndUserIds" xml:"EndUserIds"`
	EndUserNames     []string `json:"EndUserNames" xml:"EndUserNames"`
}