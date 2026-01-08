// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
	"github.com/jocall3/1231-cli/internal/requestflag"
)

func TestCorporatePerformSanctionScreening(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate", "perform-sanction-screening",
		"--country", "US",
		"--entity-type", "individual",
		"--name", "John Doe",
		"--address", "{city: Anytown, country: USA, state: CA, street: 123 Main St, zip: '90210'}",
		"--date-of-birth", "1970-01-01",
		"--identification-number", "{}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(corporatePerformSanctionScreening)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"corporate", "perform-sanction-screening",
		"--country", "US",
		"--entity-type", "individual",
		"--name", "John Doe",
		"--address.city", "Anytown",
		"--address.country", "USA",
		"--address.state", "CA",
		"--address.street", "123 Main St",
		"--address.zip", "90210",
		"--date-of-birth", "1970-01-01",
		"--identification-number", "{}",
	)
}
