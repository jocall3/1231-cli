// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
	"github.com/jocall3/1231-cli/internal/requestflag"
)

func TestNotificationsSettingsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:settings", "retrieve",
	)
}

func TestNotificationsSettingsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:settings", "update",
		"--channel-preferences", "{email: true, inApp: true, push: true, sms: true}",
		"--event-preferences", "{aiInsights: true, budgetAlerts: true, promotionalOffers: false, securityAlerts: true, transactionAlerts: true}",
		"--quiet-hours", "{enabled: true, endTime: '08:00', startTime: '22:00'}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(notificationsSettingsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:settings", "update",
		"--channel-preferences.email=true",
		"--channel-preferences.in-app=true",
		"--channel-preferences.push=true",
		"--channel-preferences.sms=true",
		"--event-preferences.ai-insights=true",
		"--event-preferences.budget-alerts=true",
		"--event-preferences.promotional-offers=false",
		"--event-preferences.security-alerts=true",
		"--event-preferences.transaction-alerts=true",
		"--quiet-hours.enabled=true",
		"--quiet-hours.end-time", "08:00",
		"--quiet-hours.start-time", "22:00",
	)
}
