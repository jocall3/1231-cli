// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
	"github.com/jocall3/1231-cli/internal/requestflag"
)

func TestPaymentsInternationalInitiate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments:international", "initiate",
		"--amount", "5000",
		"--beneficiary", "{address: 'Hauptstrasse 1, 10115 Berlin, Germany', bankName: Deutsche Bank, name: Maria Schmidt, accountNumber: {}, iban: DE89370400440532013000, routingNumber: {}, swiftBic: DEUTDEFF}",
		"--purpose", "Vendor payment for Q2 services.",
		"--source-account-id", "acc_chase_checking_4567",
		"--source-currency", "USD",
		"--target-currency", "EUR",
		"--fx-rate-lock=true",
		"--fx-rate-provider", "proprietary_ai",
		"--reference", "{}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(paymentsInternationalInitiate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments:international", "initiate",
		"--amount", "5000",
		"--beneficiary.address", "Hauptstrasse 1, 10115 Berlin, Germany",
		"--beneficiary.bank-name", "Deutsche Bank",
		"--beneficiary.name", "Maria Schmidt",
		"--beneficiary.account-number", "{}",
		"--beneficiary.iban", "DE89370400440532013000",
		"--beneficiary.routing-number", "{}",
		"--beneficiary.swift-bic", "DEUTDEFF",
		"--purpose", "Vendor payment for Q2 services.",
		"--source-account-id", "acc_chase_checking_4567",
		"--source-currency", "USD",
		"--target-currency", "EUR",
		"--fx-rate-lock=true",
		"--fx-rate-provider", "proprietary_ai",
		"--reference", "{}",
	)
}

func TestPaymentsInternationalRetrieveStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments:international", "retrieve-status",
		"--payment-id", "int_pmt_xyz7890",
	)
}
