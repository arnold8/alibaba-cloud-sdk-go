package r_kvstore

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

// DescribeRdsVSwitchs invokes the r_kvstore.DescribeRdsVSwitchs API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describerdsvswitchs.html
func (client *Client) DescribeRdsVSwitchs(request *DescribeRdsVSwitchsRequest) (response *DescribeRdsVSwitchsResponse, err error) {
	response = CreateDescribeRdsVSwitchsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRdsVSwitchsWithChan invokes the r_kvstore.DescribeRdsVSwitchs API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describerdsvswitchs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRdsVSwitchsWithChan(request *DescribeRdsVSwitchsRequest) (<-chan *DescribeRdsVSwitchsResponse, <-chan error) {
	responseChan := make(chan *DescribeRdsVSwitchsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRdsVSwitchs(request)
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

// DescribeRdsVSwitchsWithCallback invokes the r_kvstore.DescribeRdsVSwitchs API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describerdsvswitchs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRdsVSwitchsWithCallback(request *DescribeRdsVSwitchsRequest, callback func(response *DescribeRdsVSwitchsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRdsVSwitchsResponse
		var err error
		defer close(result)
		response, err = client.DescribeRdsVSwitchs(request)
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

// DescribeRdsVSwitchsRequest is the request struct for api DescribeRdsVSwitchs
type DescribeRdsVSwitchsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeRdsVSwitchsResponse is the response struct for api DescribeRdsVSwitchs
type DescribeRdsVSwitchsResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	VSwitches VSwitches `json:"VSwitches" xml:"VSwitches"`
}

// CreateDescribeRdsVSwitchsRequest creates a request to invoke DescribeRdsVSwitchs API
func CreateDescribeRdsVSwitchsRequest() (request *DescribeRdsVSwitchsRequest) {
	request = &DescribeRdsVSwitchsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeRdsVSwitchs", "", "")
	return
}

// CreateDescribeRdsVSwitchsResponse creates a response to parse from DescribeRdsVSwitchs response
func CreateDescribeRdsVSwitchsResponse() (response *DescribeRdsVSwitchsResponse) {
	response = &DescribeRdsVSwitchsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
