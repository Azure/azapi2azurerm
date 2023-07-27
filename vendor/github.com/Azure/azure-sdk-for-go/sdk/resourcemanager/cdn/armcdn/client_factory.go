//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcdn

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewManagementClient() *ManagementClient {
	subClient, _ := NewManagementClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAFDProfilesClient() *AFDProfilesClient {
	subClient, _ := NewAFDProfilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAFDCustomDomainsClient() *AFDCustomDomainsClient {
	subClient, _ := NewAFDCustomDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAFDEndpointsClient() *AFDEndpointsClient {
	subClient, _ := NewAFDEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAFDOriginGroupsClient() *AFDOriginGroupsClient {
	subClient, _ := NewAFDOriginGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAFDOriginsClient() *AFDOriginsClient {
	subClient, _ := NewAFDOriginsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRoutesClient() *RoutesClient {
	subClient, _ := NewRoutesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRuleSetsClient() *RuleSetsClient {
	subClient, _ := NewRuleSetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRulesClient() *RulesClient {
	subClient, _ := NewRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecurityPoliciesClient() *SecurityPoliciesClient {
	subClient, _ := NewSecurityPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecretsClient() *SecretsClient {
	subClient, _ := NewSecretsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewValidateClient() *ValidateClient {
	subClient, _ := NewValidateClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLogAnalyticsClient() *LogAnalyticsClient {
	subClient, _ := NewLogAnalyticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProfilesClient() *ProfilesClient {
	subClient, _ := NewProfilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEndpointsClient() *EndpointsClient {
	subClient, _ := NewEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOriginsClient() *OriginsClient {
	subClient, _ := NewOriginsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOriginGroupsClient() *OriginGroupsClient {
	subClient, _ := NewOriginGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCustomDomainsClient() *CustomDomainsClient {
	subClient, _ := NewCustomDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewResourceUsageClient() *ResourceUsageClient {
	subClient, _ := NewResourceUsageClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEdgeNodesClient() *EdgeNodesClient {
	subClient, _ := NewEdgeNodesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPoliciesClient() *PoliciesClient {
	subClient, _ := NewPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedRuleSetsClient() *ManagedRuleSetsClient {
	subClient, _ := NewManagedRuleSetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
