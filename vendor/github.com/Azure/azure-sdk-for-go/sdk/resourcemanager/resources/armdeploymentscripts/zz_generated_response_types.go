//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentscripts

// ClientCreateResponse contains the response from method Client.Create.
type ClientCreateResponse struct {
	DeploymentScriptClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClientCreateResponse.
func (c *ClientCreateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDeploymentScriptClassification(data)
	if err != nil {
		return err
	}
	c.DeploymentScriptClassification = res
	return nil
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetLogsDefaultResponse contains the response from method Client.GetLogsDefault.
type ClientGetLogsDefaultResponse struct {
	ScriptLog
}

// ClientGetLogsResponse contains the response from method Client.GetLogs.
type ClientGetLogsResponse struct {
	ScriptLogsList
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	DeploymentScriptClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClientGetResponse.
func (c *ClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDeploymentScriptClassification(data)
	if err != nil {
		return err
	}
	c.DeploymentScriptClassification = res
	return nil
}

// ClientListByResourceGroupResponse contains the response from method Client.ListByResourceGroup.
type ClientListByResourceGroupResponse struct {
	DeploymentScriptListResult
}

// ClientListBySubscriptionResponse contains the response from method Client.ListBySubscription.
type ClientListBySubscriptionResponse struct {
	DeploymentScriptListResult
}

// ClientUpdateResponse contains the response from method Client.Update.
type ClientUpdateResponse struct {
	DeploymentScriptClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClientUpdateResponse.
func (c *ClientUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDeploymentScriptClassification(data)
	if err != nil {
		return err
	}
	c.DeploymentScriptClassification = res
	return nil
}
