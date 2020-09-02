package cloudcallcenter

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

// ContactFlow is a nested struct in cloudcallcenter response
type ContactFlow struct {
	ContactFlowDescription string                                            `json:"ContactFlowDescription" xml:"ContactFlowDescription"`
	InstanceId             string                                            `json:"InstanceId" xml:"InstanceId"`
	ContactFlowName        string                                            `json:"ContactFlowName" xml:"ContactFlowName"`
	ContactFlowId          string                                            `json:"ContactFlowId" xml:"ContactFlowId"`
	AppliedVersion         string                                            `json:"AppliedVersion" xml:"AppliedVersion"`
	Type                   string                                            `json:"Type" xml:"Type"`
	Versions               VersionsInBeginContactFlowVersionModification     `json:"Versions" xml:"Versions"`
	PhoneNumbers           PhoneNumbersInBeginContactFlowVersionModification `json:"PhoneNumbers" xml:"PhoneNumbers"`
}