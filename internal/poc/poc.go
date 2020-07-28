package poc

import (
	"db-poc/internal/bootstrap"
	cardsPkg "db-poc/internal/card"
	paymentPkg "db-poc/internal/payment"
	"db-poc/pkg/prom_metrics"
	"db-poc/pkg/utils"
	"encoding/base64"
  "fmt"
  "log"
	"sync"
	"time"
)

const (
	//_, b, _, _      = runtime.Caller(0)
	//basePath        = filepath.Dir(b)
	paymentDataFile = "data/payment.json"
	cardDataFile    = "data/card.json"
)

func RunPoc(records, concurrency int64) {
	var wg sync.WaitGroup
	maxPaymentPerMerchant := records / concurrency

	log.Printf("\n=== Put() %d cards with %d concurrency", records, concurrency)
	cardNumbers := bootstrap.GetCardNumbers()
	merchantIds := bootstrap.GetMerchantIds()

	for thread := 0; thread < int(concurrency); thread++ {
		go writePaymentsHot(&wg, thread, maxPaymentPerMerchant, cardNumbers, merchantIds)
	}
	wg.Wait()
}

func writePaymentsHot(wg *sync.WaitGroup, thread int, maxPayment int64, cardIds, merchantIds []string) {
	for pId := 1; pId <= int(maxPayment); pId++ {
		wg.Add(1)
		go test(wg, cardIds, merchantIds)
	}
}

func test(wg *sync.WaitGroup, cardIds, merchantIds []string) {
	defer wg.Done()
	Test2(cardIds, merchantIds)
}


func Test2(cardIds, merchantIds []string) {
  defer prom_metrics.DbRequestDuration("transaction", true, time.Now())
  now := time.Now()
  var op *string
  s := ""
  op = &s

	pCore, _ := paymentPkg.GetCore()
	cCore, _ := cardsPkg.GetCore()
	cRepo := cardsPkg.GetRepo()
	merchantIdSize := len(merchantIds) - 1
	cardSize := len(cardIds) - 1

	cardNumber := cardIds[utils.RandomInt(cardSize)]
	merchantId := merchantIds[utils.RandomInt(merchantIdSize)]
	vaultToken := base64.StdEncoding.EncodeToString([]byte(cardNumber))
	card := &cardsPkg.Card{}
	cRepo.FindExistingCard(vaultToken, merchantId, card)
	fmt.Println("cardID", card.ID)
	if card.ID == "" {
		go prom_metrics.IncOperation("card_fetch_not_found", true)
	} else {
    go prom_metrics.IncOperation("card_fetch_found", true)
	}

	cardId := ""

	if card.ID == "" {
		newCard := cardsPkg.Card{VaultToken: vaultToken, MerchantId: merchantId, Id: utils.NewID()}
		cCore.CreateCard(newCard)
		cardId = newCard.ID
		*op += "new_card"
		go prom_metrics.IncOperation("card_write", true)
	} else {
    *op += "existing_card"
		cardId = card.ID
	}

	payment := paymentPkg.Payment{CardId: cardId, PaymentId: utils.NewID(), PartitionAt: utils.GetPartitionAt()}
	pCore.CreatePayment(payment)
  go prom_metrics.IncOperation("payment_create", true)

  go prom_metrics.IncOperation("transaction_"+*op, true)
  go  prom_metrics.DbRequestDuration("transaction", true, now)
}
