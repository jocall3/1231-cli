// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
	"github.com/jocall3/1231-cli/internal/requestflag"
)

func TestNotificationsSettingsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:settings", "retrieve",
	)
}

func TestNotificationsSettingsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
		"--channel-preferences.inApp=true",
		"--channel-preferences.push=true",
		"--channel-preferences.sms=true",
		"--event-preferences.aiInsights=true",
		"--event-preferences.budgetAlerts=true",
		"--event-preferences.promotionalOffers=false",
		"--event-preferences.securityAlerts=true",
		"--event-preferences.transactionAlerts=true",
		"--quiet-hours.enabled=true",
		"--quiet-hours.endTime", "08:00",
		"--quiet-hours.startTime", "22:00",
	)
}
