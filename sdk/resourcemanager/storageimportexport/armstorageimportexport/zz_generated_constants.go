//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorageimportexport

const (
	module  = "armstorageimportexport"
	version = "v0.1.0"
)

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

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// DriveState - The drive's current state.
type DriveState string

const (
	DriveStateCompleted         DriveState = "Completed"
	DriveStateCompletedMoreInfo DriveState = "CompletedMoreInfo"
	DriveStateNeverReceived     DriveState = "NeverReceived"
	DriveStateReceived          DriveState = "Received"
	DriveStateShippedBack       DriveState = "ShippedBack"
	DriveStateSpecified         DriveState = "Specified"
	DriveStateTransferring      DriveState = "Transferring"
)

// PossibleDriveStateValues returns the possible values for the DriveState const type.
func PossibleDriveStateValues() []DriveState {
	return []DriveState{
		DriveStateCompleted,
		DriveStateCompletedMoreInfo,
		DriveStateNeverReceived,
		DriveStateReceived,
		DriveStateShippedBack,
		DriveStateSpecified,
		DriveStateTransferring,
	}
}

// ToPtr returns a *DriveState pointing to the current value.
func (c DriveState) ToPtr() *DriveState {
	return &c
}

// EncryptionKekType - The type of kek encryption key
type EncryptionKekType string

const (
	EncryptionKekTypeCustomerManaged  EncryptionKekType = "CustomerManaged"
	EncryptionKekTypeMicrosoftManaged EncryptionKekType = "MicrosoftManaged"
)

// PossibleEncryptionKekTypeValues returns the possible values for the EncryptionKekType const type.
func PossibleEncryptionKekTypeValues() []EncryptionKekType {
	return []EncryptionKekType{
		EncryptionKekTypeCustomerManaged,
		EncryptionKekTypeMicrosoftManaged,
	}
}

// ToPtr returns a *EncryptionKekType pointing to the current value.
func (c EncryptionKekType) ToPtr() *EncryptionKekType {
	return &c
}

// IdentityType - The type of identity
type IdentityType string

const (
	IdentityTypeNone           IdentityType = "None"
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
	IdentityTypeUserAssigned   IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeUserAssigned,
	}
}

// ToPtr returns a *IdentityType pointing to the current value.
func (c IdentityType) ToPtr() *IdentityType {
	return &c
}