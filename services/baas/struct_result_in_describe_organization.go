package baas

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

// ResultInDescribeOrganization is a nested struct in baas response
type ResultInDescribeOrganization struct {
	OrganizationId  string `json:"OrganizationId" xml:"OrganizationId"`
	Name            string `json:"Name" xml:"Name"`
	RegionId        string `json:"RegionId" xml:"RegionId"`
	ZoneId          string `json:"ZoneId" xml:"ZoneId"`
	CodeName        string `json:"CodeName" xml:"CodeName"`
	Domain          string `json:"Domain" xml:"Domain"`
	Description     string `json:"Description" xml:"Description"`
	OwnerBid        string `json:"OwnerBid" xml:"OwnerBid"`
	OwnerUid        int    `json:"OwnerUid" xml:"OwnerUid"`
	OwnerName       string `json:"OwnerName" xml:"OwnerName"`
	PeerCount       int    `json:"PeerCount" xml:"PeerCount"`
	UserCount       int    `json:"UserCount" xml:"UserCount"`
	ConsortiumCount int    `json:"ConsortiumCount" xml:"ConsortiumCount"`
	SpecName        string `json:"SpecName" xml:"SpecName"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	CreateTime      string `json:"CreateTime" xml:"CreateTime"`
	State           string `json:"State" xml:"State"`
}