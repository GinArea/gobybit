// [Transfer Data Endpoints] https://bybit-exchange.github.io/docs/account_asset/#t-transfer_api
package account

// [Create Internal Transfer] https://bybit-exchange.github.io/docs/account_asset/#t-createinternaltransfer
type CreateInternalTransfer struct {
	TransferID      string   `param:"transfer_id"`
	Coin            Currency `param:"coin"`
	Amount          string   `param:"amount"`
	FromAccountType string   `param:"from_account_type"`
	ToAccountType   string   `param:"to_account_type"`
}

func (this CreateInternalTransfer) Do(client *Client) (string, bool) {
	type result struct {
		TransferID string `json:"transfer_id"`
	}
	r, ok := Post[result](client, "transfer", this)
	return r.TransferID, ok
}

func (this *Client) CreateInternalTransfer(v CreateInternalTransfer) (string, bool) {
	return v.Do(this)
}

// [Query Internal Transfer List] https://bybit-exchange.github.io/docs/account_asset/#t-querytransferlist
type QueryInternalTransferList struct {
	TransferID *string         `param:"transfer_id"`
	Coin       *Currency       `param:"coin"`
	Status     *TransferStatus `param:"status"`
	StartTime  *int            `param:"start_time"`
	EndTime    *int            `param:"end_time"`
	Direction  *PageDirection  `param:"direction"`
	Limit      *int            `param:"limit"`
	Cursor     *string         `param:"cursor"`
}

func (this QueryInternalTransferList) Do(client *Client) (InternalTransfers, bool) {
	return Get[InternalTransfers](client, "transfer/list", this)
}

type InternalTransfers struct {
	List   []InternalTransfer `json:"list"`
	Cursor string             `json:"cursor"`
}

type InternalTransfer struct {
	TransferID      string         `json:"transfer_id"`
	Coin            Currency       `json:"coin"`
	Amount          string         `json:"amount"`
	FromAccountType AccountType    `json:"from_account_type"`
	ToAccountType   AccountType    `json:"to_account_type"`
	Timestamp       string         `json:"timestamp"`
	Status          TransferStatus `json:"status"`
}

func (this *Client) QueryInternalTransferList(v QueryInternalTransferList) (InternalTransfers, bool) {
	return v.Do(this)
}
