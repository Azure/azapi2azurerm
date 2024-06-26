//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpanngfw

const (
	moduleName    = "armpanngfw"
	moduleVersion = "v1.0.0"
)

type ActionEnum string

const (
	ActionEnumAllow           ActionEnum = "Allow"
	ActionEnumDenyResetBoth   ActionEnum = "DenyResetBoth"
	ActionEnumDenyResetServer ActionEnum = "DenyResetServer"
	ActionEnumDenySilent      ActionEnum = "DenySilent"
)

// PossibleActionEnumValues returns the possible values for the ActionEnum const type.
func PossibleActionEnumValues() []ActionEnum {
	return []ActionEnum{
		ActionEnumAllow,
		ActionEnumDenyResetBoth,
		ActionEnumDenyResetServer,
		ActionEnumDenySilent,
	}
}

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

type AdvSecurityObjectTypeEnum string

const (
	AdvSecurityObjectTypeEnumFeeds     AdvSecurityObjectTypeEnum = "feeds"
	AdvSecurityObjectTypeEnumURLCustom AdvSecurityObjectTypeEnum = "urlCustom"
)

// PossibleAdvSecurityObjectTypeEnumValues returns the possible values for the AdvSecurityObjectTypeEnum const type.
func PossibleAdvSecurityObjectTypeEnumValues() []AdvSecurityObjectTypeEnum {
	return []AdvSecurityObjectTypeEnum{
		AdvSecurityObjectTypeEnumFeeds,
		AdvSecurityObjectTypeEnumURLCustom,
	}
}

// BillingCycle - Billing cycle
type BillingCycle string

const (
	BillingCycleMONTHLY BillingCycle = "MONTHLY"
	BillingCycleWEEKLY  BillingCycle = "WEEKLY"
)

// PossibleBillingCycleValues returns the possible values for the BillingCycle const type.
func PossibleBillingCycleValues() []BillingCycle {
	return []BillingCycle{
		BillingCycleMONTHLY,
		BillingCycleWEEKLY,
	}
}

// BooleanEnum - Boolean Enum
type BooleanEnum string

const (
	BooleanEnumFALSE BooleanEnum = "FALSE"
	BooleanEnumTRUE  BooleanEnum = "TRUE"
)

// PossibleBooleanEnumValues returns the possible values for the BooleanEnum const type.
func PossibleBooleanEnumValues() []BooleanEnum {
	return []BooleanEnum{
		BooleanEnumFALSE,
		BooleanEnumTRUE,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DNSProxy - DNS Proxy
type DNSProxy string

const (
	DNSProxyDISABLED DNSProxy = "DISABLED"
	DNSProxyENABLED  DNSProxy = "ENABLED"
)

// PossibleDNSProxyValues returns the possible values for the DNSProxy const type.
func PossibleDNSProxyValues() []DNSProxy {
	return []DNSProxy{
		DNSProxyDISABLED,
		DNSProxyENABLED,
	}
}

type DecryptionRuleTypeEnum string

const (
	DecryptionRuleTypeEnumNone                  DecryptionRuleTypeEnum = "None"
	DecryptionRuleTypeEnumSSLInboundInspection  DecryptionRuleTypeEnum = "SSLInboundInspection"
	DecryptionRuleTypeEnumSSLOutboundInspection DecryptionRuleTypeEnum = "SSLOutboundInspection"
)

// PossibleDecryptionRuleTypeEnumValues returns the possible values for the DecryptionRuleTypeEnum const type.
func PossibleDecryptionRuleTypeEnumValues() []DecryptionRuleTypeEnum {
	return []DecryptionRuleTypeEnum{
		DecryptionRuleTypeEnumNone,
		DecryptionRuleTypeEnumSSLInboundInspection,
		DecryptionRuleTypeEnumSSLOutboundInspection,
	}
}

// DefaultMode - Type for Default Mode for rules creation
type DefaultMode string

const (
	DefaultModeFIREWALL DefaultMode = "FIREWALL"
	DefaultModeIPS      DefaultMode = "IPS"
	DefaultModeNONE     DefaultMode = "NONE"
)

// PossibleDefaultModeValues returns the possible values for the DefaultMode const type.
func PossibleDefaultModeValues() []DefaultMode {
	return []DefaultMode{
		DefaultModeFIREWALL,
		DefaultModeIPS,
		DefaultModeNONE,
	}
}

// EgressNat - Egress NAT
type EgressNat string

const (
	EgressNatDISABLED EgressNat = "DISABLED"
	EgressNatENABLED  EgressNat = "ENABLED"
)

// PossibleEgressNatValues returns the possible values for the EgressNat const type.
func PossibleEgressNatValues() []EgressNat {
	return []EgressNat{
		EgressNatDISABLED,
		EgressNatENABLED,
	}
}

// EnabledDNSType - Enabled DNS type values
type EnabledDNSType string

const (
	EnabledDNSTypeAZURE  EnabledDNSType = "AZURE"
	EnabledDNSTypeCUSTOM EnabledDNSType = "CUSTOM"
)

// PossibleEnabledDNSTypeValues returns the possible values for the EnabledDNSType const type.
func PossibleEnabledDNSTypeValues() []EnabledDNSType {
	return []EnabledDNSType{
		EnabledDNSTypeAZURE,
		EnabledDNSTypeCUSTOM,
	}
}

// HealthStatus - Status Codes for the Firewall
type HealthStatus string

const (
	HealthStatusGREEN        HealthStatus = "GREEN"
	HealthStatusINITIALIZING HealthStatus = "INITIALIZING"
	HealthStatusRED          HealthStatus = "RED"
	HealthStatusYELLOW       HealthStatus = "YELLOW"
)

// PossibleHealthStatusValues returns the possible values for the HealthStatus const type.
func PossibleHealthStatusValues() []HealthStatus {
	return []HealthStatus{
		HealthStatusGREEN,
		HealthStatusINITIALIZING,
		HealthStatusRED,
		HealthStatusYELLOW,
	}
}

// LogOption - Log options possible
type LogOption string

const (
	LogOptionINDIVIDUALDESTINATION LogOption = "INDIVIDUAL_DESTINATION"
	LogOptionSAMEDESTINATION       LogOption = "SAME_DESTINATION"
)

// PossibleLogOptionValues returns the possible values for the LogOption const type.
func PossibleLogOptionValues() []LogOption {
	return []LogOption{
		LogOptionINDIVIDUALDESTINATION,
		LogOptionSAMEDESTINATION,
	}
}

// LogType - Possible log types
type LogType string

const (
	LogTypeAUDIT      LogType = "AUDIT"
	LogTypeDECRYPTION LogType = "DECRYPTION"
	LogTypeDLP        LogType = "DLP"
	LogTypeTHREAT     LogType = "THREAT"
	LogTypeTRAFFIC    LogType = "TRAFFIC"
	LogTypeWILDFIRE   LogType = "WILDFIRE"
)

// PossibleLogTypeValues returns the possible values for the LogType const type.
func PossibleLogTypeValues() []LogType {
	return []LogType{
		LogTypeAUDIT,
		LogTypeDECRYPTION,
		LogTypeDLP,
		LogTypeTHREAT,
		LogTypeTRAFFIC,
		LogTypeWILDFIRE,
	}
}

// ManagedIdentityType - The kind of managed identity assigned to this resource.
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                  ManagedIdentityType = "None"
	ManagedIdentityTypeSystemAndUserAssigned ManagedIdentityType = "SystemAssigned,UserAssigned"
	ManagedIdentityTypeSystemAssigned        ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeUserAssigned          ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeNone,
		ManagedIdentityTypeSystemAndUserAssigned,
		ManagedIdentityTypeSystemAssigned,
		ManagedIdentityTypeUserAssigned,
	}
}

// MarketplaceSubscriptionStatus - Marketplace Subscription Status
type MarketplaceSubscriptionStatus string

const (
	MarketplaceSubscriptionStatusFulfillmentRequested    MarketplaceSubscriptionStatus = "FulfillmentRequested"
	MarketplaceSubscriptionStatusNotStarted              MarketplaceSubscriptionStatus = "NotStarted"
	MarketplaceSubscriptionStatusPendingFulfillmentStart MarketplaceSubscriptionStatus = "PendingFulfillmentStart"
	MarketplaceSubscriptionStatusSubscribed              MarketplaceSubscriptionStatus = "Subscribed"
	MarketplaceSubscriptionStatusSuspended               MarketplaceSubscriptionStatus = "Suspended"
	MarketplaceSubscriptionStatusUnsubscribed            MarketplaceSubscriptionStatus = "Unsubscribed"
)

// PossibleMarketplaceSubscriptionStatusValues returns the possible values for the MarketplaceSubscriptionStatus const type.
func PossibleMarketplaceSubscriptionStatusValues() []MarketplaceSubscriptionStatus {
	return []MarketplaceSubscriptionStatus{
		MarketplaceSubscriptionStatusFulfillmentRequested,
		MarketplaceSubscriptionStatusNotStarted,
		MarketplaceSubscriptionStatusPendingFulfillmentStart,
		MarketplaceSubscriptionStatusSubscribed,
		MarketplaceSubscriptionStatusSuspended,
		MarketplaceSubscriptionStatusUnsubscribed,
	}
}

// NetworkType - NetworkType Enum
type NetworkType string

const (
	NetworkTypeVNET NetworkType = "VNET"
	NetworkTypeVWAN NetworkType = "VWAN"
)

// PossibleNetworkTypeValues returns the possible values for the NetworkType const type.
func PossibleNetworkTypeValues() []NetworkType {
	return []NetworkType{
		NetworkTypeVNET,
		NetworkTypeVWAN,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProtocolType - Protocol Enum
type ProtocolType string

const (
	ProtocolTypeTCP ProtocolType = "TCP"
	ProtocolTypeUDP ProtocolType = "UDP"
)

// PossibleProtocolTypeValues returns the possible values for the ProtocolType const type.
func PossibleProtocolTypeValues() []ProtocolType {
	return []ProtocolType{
		ProtocolTypeTCP,
		ProtocolTypeUDP,
	}
}

// ProvisioningState - Provisioning state of the firewall resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ReadOnlyProvisioningState - Provisioning state of the firewall resource.
type ReadOnlyProvisioningState string

const (
	ReadOnlyProvisioningStateDeleted   ReadOnlyProvisioningState = "Deleted"
	ReadOnlyProvisioningStateFailed    ReadOnlyProvisioningState = "Failed"
	ReadOnlyProvisioningStateSucceeded ReadOnlyProvisioningState = "Succeeded"
)

// PossibleReadOnlyProvisioningStateValues returns the possible values for the ReadOnlyProvisioningState const type.
func PossibleReadOnlyProvisioningStateValues() []ReadOnlyProvisioningState {
	return []ReadOnlyProvisioningState{
		ReadOnlyProvisioningStateDeleted,
		ReadOnlyProvisioningStateFailed,
		ReadOnlyProvisioningStateSucceeded,
	}
}

// ScopeType - Rulestack Type
type ScopeType string

const (
	ScopeTypeGLOBAL ScopeType = "GLOBAL"
	ScopeTypeLOCAL  ScopeType = "LOCAL"
)

// PossibleScopeTypeValues returns the possible values for the ScopeType const type.
func PossibleScopeTypeValues() []ScopeType {
	return []ScopeType{
		ScopeTypeGLOBAL,
		ScopeTypeLOCAL,
	}
}

type SecurityServicesTypeEnum string

const (
	SecurityServicesTypeEnumAntiSpyware      SecurityServicesTypeEnum = "antiSpyware"
	SecurityServicesTypeEnumAntiVirus        SecurityServicesTypeEnum = "antiVirus"
	SecurityServicesTypeEnumDNSSubscription  SecurityServicesTypeEnum = "dnsSubscription"
	SecurityServicesTypeEnumFileBlocking     SecurityServicesTypeEnum = "fileBlocking"
	SecurityServicesTypeEnumIPsVulnerability SecurityServicesTypeEnum = "ipsVulnerability"
	SecurityServicesTypeEnumURLFiltering     SecurityServicesTypeEnum = "urlFiltering"
)

// PossibleSecurityServicesTypeEnumValues returns the possible values for the SecurityServicesTypeEnum const type.
func PossibleSecurityServicesTypeEnumValues() []SecurityServicesTypeEnum {
	return []SecurityServicesTypeEnum{
		SecurityServicesTypeEnumAntiSpyware,
		SecurityServicesTypeEnumAntiVirus,
		SecurityServicesTypeEnumDNSSubscription,
		SecurityServicesTypeEnumFileBlocking,
		SecurityServicesTypeEnumIPsVulnerability,
		SecurityServicesTypeEnumURLFiltering,
	}
}

// ServerStatus - Connectivity Status for Panorama Server
type ServerStatus string

const (
	ServerStatusDOWN ServerStatus = "DOWN"
	ServerStatusUP   ServerStatus = "UP"
)

// PossibleServerStatusValues returns the possible values for the ServerStatus const type.
func PossibleServerStatusValues() []ServerStatus {
	return []ServerStatus{
		ServerStatusDOWN,
		ServerStatusUP,
	}
}

// StateEnum - Enabled or Disabled Enum
type StateEnum string

const (
	StateEnumDISABLED StateEnum = "DISABLED"
	StateEnumENABLED  StateEnum = "ENABLED"
)

// PossibleStateEnumValues returns the possible values for the StateEnum const type.
func PossibleStateEnumValues() []StateEnum {
	return []StateEnum{
		StateEnumDISABLED,
		StateEnumENABLED,
	}
}

// UsageType - Usage Type
type UsageType string

const (
	UsageTypeCOMMITTED UsageType = "COMMITTED"
	UsageTypePAYG      UsageType = "PAYG"
)

// PossibleUsageTypeValues returns the possible values for the UsageType const type.
func PossibleUsageTypeValues() []UsageType {
	return []UsageType{
		UsageTypeCOMMITTED,
		UsageTypePAYG,
	}
}
