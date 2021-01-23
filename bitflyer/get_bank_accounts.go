package bitflyer

import (
	"context"
)

// BankAccount represents a bank account that registered to your bitFlyer account.
type BankAccount struct {
	ID            int64  `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	BankName      string `json:"bank_name"`
	BranchName    string `json:"branch_name"`
	AccountType   string `json:"account_type"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

// GetBankAccountListOutput represent an output of GetBankAccountList method.
type GetBankAccountListOutput struct {
	BankAccounts []*BankAccount
}

// GetBankAccountList gets summary of bank accounts.
func (c *Client) GetBankAccountList(ctx context.Context) (*GetBankAccountListOutput, error) {

	req, err := c.NewRequest(ctx, "GET", "getbankaccounts", nil, nil, true)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	output := GetBankAccountListOutput{}
	if err := decodeBody(res, &output.BankAccounts); err != nil {
		return nil, err
	}

	return &output, nil
}
