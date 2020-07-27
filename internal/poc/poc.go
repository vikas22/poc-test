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
	pCore, _ := paymentPkg.GetCore()
	cCore, _ := cardsPkg.GetCore()
	cRepo := cardsPkg.GetRepo()
	merchantIdSize := len(merchantIds) - 1
	cardSize := len(cardIds) - 1
	op := "poc"

	now := time.Now()
	cardNumber := cardIds[utils.RandomInt(cardSize)]
	merchantId := merchantIds[utils.RandomInt(merchantIdSize)]
	vaultToken := base64.StdEncoding.EncodeToString([]byte(cardNumber))
	cardFetchTime := time.Now()
	card := &cardsPkg.Card{}
	cRepo.FindExistingCard(vaultToken, merchantId, card)
	fmt.Println("cardID", card.ID)
	if card.ID == "" {
		prom_metrics.IncOperation(op+"_card_fetch_not_found", true)
		prom_metrics.DbRequestDuration(op+"_card_fetch_not_found", true, cardFetchTime)
	} else {
		prom_metrics.IncOperation(op+"_card_fetch_found", true)
		prom_metrics.DbRequestDuration(op+"_card_fetch_found", true, cardFetchTime)
	}

	cardId := ""

	if card.ID == "" {
		cardWriteTime := time.Now()
		newCard := cardsPkg.Card{VaultToken: vaultToken, MerchantId: merchantId, Id: utils.NewID()}
		cCore.CreateCard(newCard)
		cardId = newCard.ID
		prom_metrics.IncOperation(op+"_card_write", true)
		prom_metrics.DbRequestDuration(op+"_card_write", false, cardWriteTime)
		op += "_with_new_card"
	} else {
		cardId = card.ID
		op += "_with_existing_card"
	}

	paymentOp := "payment_write"
	paymentWriteTime := time.Now()
	payment := paymentPkg.Payment{CardId: cardId, PaymentId: utils.NewID(), PartitionAt: utils.GetPartitionAt()}
	err := pCore.CreatePayment(payment)

	prom_metrics.IncOperation(paymentOp, true)
	prom_metrics.DbRequestDuration(paymentOp, true, paymentWriteTime)
	fmt.Println(op)
	if err != nil {
		log.Println(err)
		prom_metrics.IncOperation(op, false)
		prom_metrics.DbRequestDuration(op, false, now)
	} else {
		prom_metrics.IncOperation(op, true)
		prom_metrics.DbRequestDuration(op, true, now)
	}
}
