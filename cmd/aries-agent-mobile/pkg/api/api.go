/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

// AriesController provides Aries agent protocols tailored to mobile platforms.
type AriesController interface {

	// GetIntroduceController returns an implementation of IntroduceController
	GetIntroduceController() (IntroduceController, error)

	// GetVerifiableController returns an implementation of VerifiableController
	GetVerifiableController() (VerifiableController, error)

	// GetIssueCredentialController returns an implementation of IssueCredentialController
	GetIssueCredentialController() (IssueCredentialController, error)

	// GetPresentProofController returns an implementation of PresentProofController
	GetPresentProofController() (PresentProofController, error)

	// GetDIDExchangeController returns an implementation of DIDExchangeController
	GetDIDExchangeController() (DIDExchangeController, error)

	// GetVDRIController returns an implementation of VDRIController
	GetVDRIController() (VDRIController, error)

	// GetMediatorController returns an implementation of MediatorController
	GetMediatorController() (MediatorController, error)

	// GetMessagingController returns an implementation of MessagingController
	GetMessagingController() (MessagingController, error)

	// GetOutOfBandController returns an implementation of OutOfBandController
	GetOutOfBandController() (OutOfBandController, error)

	// GetKMSController returns an implementation of KMSController
	GetKMSController() (KMSController, error)

	// RegisterNotifier associates a notifier to relevant topics
	RegisterNotifier(n Notifier, topics string) error
}
