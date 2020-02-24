package hbase

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

// DescribeInstanceType invokes the hbase.DescribeInstanceType API synchronously
// api document: https://help.aliyun.com/api/hbase/describeinstancetype.html
func (client *Client) DescribeInstanceType(request *DescribeInstanceTypeRequest) (response *DescribeInstanceTypeResponse, err error) {
	response = CreateDescribeInstanceTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceTypeWithChan invokes the hbase.DescribeInstanceType API asynchronously
// api document: https://help.aliyun.com/api/hbase/describeinstancetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceTypeWithChan(request *DescribeInstanceTypeRequest) (<-chan *DescribeInstanceTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceType(request)
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

// DescribeInstanceTypeWithCallback invokes the hbase.DescribeInstanceType API asynchronously
// api document: https://help.aliyun.com/api/hbase/describeinstancetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceTypeWithCallback(request *DescribeInstanceTypeRequest, callback func(response *DescribeInstanceTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceType(request)
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

// DescribeInstanceTypeRequest is the request struct for api DescribeInstanceType
type DescribeInstanceTypeRequest struct {
	*requests.RpcRequest
	InstanceType string `position:"Query" name:"InstanceType"`
}

// DescribeInstanceTypeResponse is the response struct for api DescribeInstanceType
type DescribeInstanceTypeResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	InstanceTypeSpecList InstanceTypeSpecList `json:"InstanceTypeSpecList" xml:"InstanceTypeSpecList"`
}

// CreateDescribeInstanceTypeRequest creates a request to invoke DescribeInstanceType API
func CreateDescribeInstanceTypeRequest() (request *DescribeInstanceTypeRequest) {
	request = &DescribeInstanceTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeInstanceType", "", "")
	return
}

// CreateDescribeInstanceTypeResponse creates a response to parse from DescribeInstanceType response
func CreateDescribeInstanceTypeResponse() (response *DescribeInstanceTypeResponse) {
	response = &DescribeInstanceTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
