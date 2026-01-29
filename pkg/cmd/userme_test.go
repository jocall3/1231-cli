// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
	"github.com/jocall3/1231-cli/internal/requestflag"
)

func TestUsersMeRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me", "retrieve",
	)
}

func TestUsersMeUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me", "update",
		"--address", "{city: Anytown, country: USA, state: CA, street: 123 Main St, zip: '90210'}",
		"--name", "Quantum Visionary Pro",
		"--phone", "+1-555-999-0000",
		"--preferences", "{aiInteractionMode: balanced, dataSharingConsent: true, notificationChannels: {email: true, inApp: true, push: true, sms: false}, preferredLanguage: en-US, theme: Dark-Quantum, transactionGrouping: category}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(usersMeUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:me", "update",
		"--address.city", "Anytown",
		"--address.country", "USA",
		"--address.state", "CA",
		"--address.street", "123 Main St",
		"--address.zip", "90210",
		"--name", "Quantum Visionary Pro",
		"--phone", "+1-555-999-0000",
		"--preferences.ai-interaction-mode", "balanced",
		"--preferences.data-sharing-consent=true",
		"--preferences.notification-channels", "{email: true, inApp: true, push: true, sms: false}",
		"--preferences.preferred-language", "en-US",
		"--preferences.theme", "Dark-Quantum",
		"--preferences.transaction-grouping", "category",
	)
}
