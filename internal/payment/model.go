package payment

import "time"

// Authentication - Store all the Authentication data for card
type Payment struct {
  PaymentId            string `json:"payment_id"`
	MerchantId            string
	Amount                int64
	Currency              string
	BaseAmount            int64
	Method                string
	Status                string
	TwoFactorAuth         string
	orderId               string
	International         bool
	AmountAuthorized      int64
	AmountRefunded        int64
	BaseAmountRefunded    int64
	RefundStatus          string
	Description           string
	CardId                string
	ErrorCode             string
	InternalErrorCode     string
	ErrorDescription      string
	CustomerId            string
	GlobalCustomerId      string
	AppToken              string
	TokenId               string
	Email                 string
	Contact               string
	TransactionId         string
	AuthorizedAt          int64
	AutoCaptured          bool
	CapturedAt            int64
	Gateway               string
	TerminalId            string
	AuthenticationGateway string
	GatewayCaptured       bool
	CallbackUrl           string
	Fee                   int64
	Mdr                   int64
	Tax                   int64
	Recurring             bool
	Save                  bool
	LateAuthorized        bool
	ConvertCurrency       bool
	CreatedAt             int64
  PartitionAt             time.Time
}

// TableName - Get the authentication table name (impl gorm tabler interface)
func (*Payment) TableName() string {
	return "payments"
}
