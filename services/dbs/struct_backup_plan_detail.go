package dbs

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

// BackupPlanDetail is a nested struct in dbs response
type BackupPlanDetail struct {
	CrossRoleName                        string `json:"CrossRoleName" xml:"CrossRoleName"`
	SourceEndpointInstanceType           string `json:"SourceEndpointInstanceType" xml:"SourceEndpointInstanceType"`
	BackupSetDownloadDir                 string `json:"BackupSetDownloadDir" xml:"BackupSetDownloadDir"`
	SourceEndpointIpPort                 string `json:"SourceEndpointIpPort" xml:"SourceEndpointIpPort"`
	CrossAliyunId                        string `json:"CrossAliyunId" xml:"CrossAliyunId"`
	DuplicationArchivePeriod             int    `json:"DuplicationArchivePeriod" xml:"DuplicationArchivePeriod"`
	BackupPlanId                         string `json:"BackupPlanId" xml:"BackupPlanId"`
	EndTimestampForRestore               int64  `json:"EndTimestampForRestore" xml:"EndTimestampForRestore"`
	BackupPlanStatus                     string `json:"BackupPlanStatus" xml:"BackupPlanStatus"`
	BackupSetDownloadFullDataFormat      string `json:"BackupSetDownloadFullDataFormat" xml:"BackupSetDownloadFullDataFormat"`
	BackupRetentionPeriod                int    `json:"BackupRetentionPeriod" xml:"BackupRetentionPeriod"`
	OSSBucketRegion                      string `json:"OSSBucketRegion" xml:"OSSBucketRegion"`
	SourceEndpointOracleSID              string `json:"SourceEndpointOracleSID" xml:"SourceEndpointOracleSID"`
	BackupStorageType                    string `json:"BackupStorageType" xml:"BackupStorageType"`
	BackupMethod                         string `json:"BackupMethod" xml:"BackupMethod"`
	SourceEndpointRegion                 string `json:"SourceEndpointRegion" xml:"SourceEndpointRegion"`
	BackupPeriod                         string `json:"BackupPeriod" xml:"BackupPeriod"`
	SourceEndpointDatabaseName           string `json:"SourceEndpointDatabaseName" xml:"SourceEndpointDatabaseName"`
	BackupSetDownloadGatewayId           int64  `json:"BackupSetDownloadGatewayId" xml:"BackupSetDownloadGatewayId"`
	BackupPlanCreateTime                 int64  `json:"BackupPlanCreateTime" xml:"BackupPlanCreateTime"`
	InstanceClass                        string `json:"InstanceClass" xml:"InstanceClass"`
	BackupSetDownloadTargetType          string `json:"BackupSetDownloadTargetType" xml:"BackupSetDownloadTargetType"`
	DuplicationInfrequentAccessPeriod    int    `json:"DuplicationInfrequentAccessPeriod" xml:"DuplicationInfrequentAccessPeriod"`
	BackupStartTime                      string `json:"BackupStartTime" xml:"BackupStartTime"`
	ErrMessage                           string `json:"ErrMessage" xml:"ErrMessage"`
	BackupObjects                        string `json:"BackupObjects" xml:"BackupObjects"`
	BeginTimestampForRestore             int64  `json:"BeginTimestampForRestore" xml:"BeginTimestampForRestore"`
	SourceEndpointInstanceID             string `json:"SourceEndpointInstanceID" xml:"SourceEndpointInstanceID"`
	OpenBackupSetAutoDownload            bool   `json:"OpenBackupSetAutoDownload" xml:"OpenBackupSetAutoDownload"`
	BackupPlanName                       string `json:"BackupPlanName" xml:"BackupPlanName"`
	OSSBucketName                        string `json:"OSSBucketName" xml:"OSSBucketName"`
	BackupGatewayId                      int64  `json:"BackupGatewayId" xml:"BackupGatewayId"`
	SourceEndpointUserName               string `json:"SourceEndpointUserName" xml:"SourceEndpointUserName"`
	BackupSetDownloadIncrementDataFormat string `json:"BackupSetDownloadIncrementDataFormat" xml:"BackupSetDownloadIncrementDataFormat"`
	EnableBackupLog                      bool   `json:"EnableBackupLog" xml:"EnableBackupLog"`
	ResourceGroupId                      string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	DatabaseType                         string `json:"DatabaseType" xml:"DatabaseType"`
}
