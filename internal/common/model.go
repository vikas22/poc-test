package common

var (
	DEFAULT_LIMIT = 20
)

// listOptions request to fetch admin entities
type ListOptions struct {
	Count      int    `form:"count"`
	Skip       int    `form:"skip"`
	PaymentID  string `form:"payment_id"`
	Gateway    string `form:"gateway"`
	Status     string `form:"status"`
	MerchantID string `form:"merchant_id"`
}

// Model - Base DB model
type Model struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

//// BeforeCreate - Sets the ID for model
//func (*Model) BeforeCreate(scope *gorm.Scope) (err error) {
//	rzpID := utils.NewID()
//
//	if err == nil {
//		scope.SetColumn("ID", rzpID)
//	}
//
//	return
//}
