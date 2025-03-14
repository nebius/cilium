package vpc

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

// QueueListItem is a nested struct in vpc response
type QueueListItem struct {
	QueueName        string         `json:"QueueName" xml:"QueueName"`
	QueueDescription string         `json:"QueueDescription" xml:"QueueDescription"`
	QosId            string         `json:"QosId" xml:"QosId"`
	QueueId          string         `json:"QueueId" xml:"QueueId"`
	QueueType        string         `json:"QueueType" xml:"QueueType"`
	BandwidthPercent string         `json:"BandwidthPercent" xml:"BandwidthPercent"`
	Status           string         `json:"Status" xml:"Status"`
	RuleList         []RuleListItem `json:"RuleList" xml:"RuleList"`
}
