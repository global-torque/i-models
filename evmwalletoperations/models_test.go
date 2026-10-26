package evmwalletoperations

import "testing"

func TestWalletSessionReservationConsumptionIsInternal(t *testing.T) {
	t.Parallel()

	operation := WalletOperation{
		WalletSessionReservationConsumed: true,
	}
	if got := operation.GetValueByTag(
		"wallet_session_reservation_consumed",
	); got != true {
		t.Fatalf(
			"GetValueByTag(wallet_session_reservation_consumed) = %#v, want true",
			got,
		)
	}
	if _, exposed := operation.ToJSON()["wallet_session_reservation_consumed"]; exposed {
		t.Fatal(
			"WalletOperation.ToJSON() exposed internal reservation consumption",
		)
	}
}

func TestTransactionFromAddressProjection(t *testing.T) {
	t.Parallel()

	fromAddress := "0x1111111111111111111111111111111111111111"
	operation := WalletOperation{TransactionFromAddress: &fromAddress}

	if got := operation.GetValueByTag(
		"transaction_from_address",
	); got != operation.TransactionFromAddress {
		t.Fatalf(
			"GetValueByTag(transaction_from_address) = %#v, want %#v",
			got,
			operation.TransactionFromAddress,
		)
	}
	if got := operation.ToJSON()["transaction_from_address"]; got != operation.TransactionFromAddress {
		t.Fatalf(
			"ToJSON transaction_from_address = %#v, want %#v",
			got,
			operation.TransactionFromAddress,
		)
	}
}

func TestWalletOperationProgressEvidenceProjection(t *testing.T) {
	t.Parallel()

	receiptStatus := "success"
	actionStatus := "verified"
	operation := WalletOperation{
		ReceiptExecutionStatus: &receiptStatus,
		ExpectedActionStatus:   &actionStatus,
	}

	for field, want := range map[string]any{
		"receipt_execution_status":     operation.ReceiptExecutionStatus,
		"expected_action_status":       operation.ExpectedActionStatus,
		"expected_action_failure_code": operation.ExpectedActionFailureCode,
	} {
		if got := operation.GetValueByTag(field); got != want {
			t.Fatalf("GetValueByTag(%s) = %#v, want %#v", field, got, want)
		}
		if got := operation.ToJSON()[field]; got != want {
			t.Fatalf("ToJSON[%s] = %#v, want %#v", field, got, want)
		}
	}
}
