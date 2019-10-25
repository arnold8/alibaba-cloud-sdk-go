package yundun_dbaudit

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

// RefundInstance invokes the yundun_dbaudit.RefundInstance API synchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/refundinstance.html
func (client *Client) RefundInstance(request *RefundInstanceRequest) (response *RefundInstanceResponse, err error) {
	response = CreateRefundInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RefundInstanceWithChan invokes the yundun_dbaudit.RefundInstance API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/refundinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefundInstanceWithChan(request *RefundInstanceRequest) (<-chan *RefundInstanceResponse, <-chan error) {
	responseChan := make(chan *RefundInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefundInstance(request)
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

// RefundInstanceWithCallback invokes the yundun_dbaudit.RefundInstance API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/refundinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefundInstanceWithCallback(request *RefundInstanceRequest, callback func(response *RefundInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefundInstanceResponse
		var err error
		defer close(result)
		response, err = client.RefundInstance(request)
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

// RefundInstanceRequest is the request struct for api RefundInstance
type RefundInstanceRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	Name     requests.Integer `position:"Query" name:"Name"`
	Id       requests.Integer `position:"Query" name:"id"`
}

// RefundInstanceResponse is the response struct for api RefundInstance
type RefundInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRefundInstanceRequest creates a request to invoke RefundInstance API
func CreateRefundInstanceRequest() (request *RefundInstanceRequest) {
	request = &RefundInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-dbaudit", "2018-10-29", "RefundInstance", "dbaudit", "openAPI")
	return
}

// CreateRefundInstanceResponse creates a response to parse from RefundInstance response
func CreateRefundInstanceResponse() (response *RefundInstanceResponse) {
	response = &RefundInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
