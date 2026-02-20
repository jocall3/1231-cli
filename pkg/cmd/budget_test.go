// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
	"github.com/jocall3/1231-cli/internal/requestflag"
)

func TestBudgetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"budgets", "create",
		"--end-date", "2024-09-30",
		"--name", "September Living Expenses",
		"--period", "monthly",
		"--start-date", "2024-09-01",
		"--total-amount", "2800",
		"--ai-auto-populate=true",
		"--alert-threshold", "75",
		"--category", "{allocated: 1500, name: Rent}",
		"--category", "{allocated: 400, name: Groceries}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(budgetsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"budgets", "create",
		"--end-date", "2024-09-30",
		"--name", "September Living Expenses",
		"--period", "monthly",
		"--start-date", "2024-09-01",
		"--total-amount", "2800",
		"--ai-auto-populate=true",
		"--alert-threshold", "75",
		"--category.allocated", "1500",
		"--category.name", "Rent",
		"--category.allocated", "400",
		"--category.name", "Groceries",
	)
}

func TestBudgetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"budgets", "retrieve",
		"--budget-id", "budget_monthly_aug",
	)
}

func TestBudgetsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"budgets", "update",
		"--budget-id", "budget_monthly_aug",
		"--alert-threshold", "85",
		"--category", "{allocated: 550, name: Groceries}",
		"--end-date", "2024-08-31",
		"--name", "August 2024 Revised Household Budget",
		"--start-date", "2024-08-01",
		"--status", "active",
		"--total-amount", "3200",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(budgetsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"budgets", "update",
		"--budget-id", "budget_monthly_aug",
		"--alert-threshold", "85",
		"--category.allocated", "550",
		"--category.name", "Groceries",
		"--end-date", "2024-08-31",
		"--name", "August 2024 Revised Household Budget",
		"--start-date", "2024-08-01",
		"--status", "active",
		"--total-amount", "3200",
	)
}

func TestBudgetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"budgets", "list",
		"--limit", "{}",
		"--offset", "{}",
	)
}

func TestBudgetsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"budgets", "delete",
		"--budget-id", "budget_monthly_aug",
	)
}
