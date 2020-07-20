package bootstrap

import (
	"db-poc/internal/config"
	configPkg "db-poc/pkg/config"
	"db-poc/pkg/db/sql"
	"db-poc/pkg/utils"
  "fmt"
  "io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	// Config holds app config
	Config config.AppConfig
	// Db holds Database connection
	Db *sql.Db
)

// Init configures/initializes various internal pkgs
func Init() error {
	var err error

	if err = loadConfig(); err != nil {
		return err
	}

	Db, err = sql.NewDb(&Config.Database, nil)

	if err != nil {
		return err
	}

	return nil
}

func loadConfig() error {
	err := configPkg.NewDefaultConfig().Load(&Config)

	if err != nil {
		return err
	}

	return nil
}

const CARD_NUMBER_LIMIT = 1000
const MERCHANT_LIMIT = 1000

const CARD_NUMBER_FILE = "/data/cards.txt"
const MERCHANT_FILE = "/data/merchantids.txt"

func Bootstrap() {
	generateCardNumbers()
	generateMerchantIds()
}

func generateCardNumbers() {
	pwd, _ := os.Getwd()
	path := pwd + CARD_NUMBER_FILE
	fmt.Println(path)
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	i := 0
	for i < CARD_NUMBER_LIMIT {
		rand.Seed(time.Now().UTC().UnixNano())
		random := rand.Int63n(1e16)
		if _, err := f.WriteString(fmt.Sprintf("%016d", rand.Int63n(random)) + "\n"); err != nil {
			log.Println(err)
		}
		i++
	}
}

func generateMerchantIds() {
	pwd, _ := os.Getwd()
	path := pwd + MERCHANT_FILE
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	i := 0
	for i < MERCHANT_LIMIT {
		rand.Seed(time.Now().Unix())
		if _, err := f.WriteString(utils.NewID() + "\n"); err != nil {
			log.Println(err)
		}
		i++
	}
}

func GetMerchantIds() []string {
	pwd, _ := os.Getwd()
	path := pwd + MERCHANT_FILE
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Merchant ID load Error %v", err)
	}
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

func GetCardNumbers() []string {
	pwd, _ := os.Getwd()
	path := pwd + CARD_NUMBER_FILE
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Card number load Error %v", err)
	}
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	return lines
}
