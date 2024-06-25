package db

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func CreateDatabaseAndTable() error {
	// Criação do banco de dados "testdb" se não existir
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS testdb")
	if err != nil {
		log.Printf("Error creating database: %v", err)
		return err
	}

	// Uso do banco de dados "testdb"
	_, err = db.Exec("USE testdb")
	if err != nil {
		log.Printf("Error using database: %v", err)
		return err
	}

	// Criação da tabela "myTable" se não existir
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS myTable (
			PartnerId VARCHAR(255),
			PartnerName VARCHAR(255),
			CustomerId VARCHAR(255),
			CustomerName VARCHAR(255),
			CustomerDomainName VARCHAR(255),
			CustomerCountry VARCHAR(255),
			MpnId INT,
			Tier2MpnId INT,
			InvoiceNumber VARCHAR(255),
			ProductId VARCHAR(255),
			SkuId VARCHAR(255),
			AvailabilityId INT,
			SkuName VARCHAR(255),
			ProductName VARCHAR(255),
			PublisherName VARCHAR(255),
			PublisherId INT,
			SubscriptionDescription VARCHAR(255),
			SubscriptionId VARCHAR(255),
			ChargeStartDate DATE,
			ChargeEndDate DATE,
			UsageDate DATE,
			MeterType VARCHAR(255),
			MeterCategory VARCHAR(255),
			MeterId VARCHAR(255),
			MeterSubCategory VARCHAR(255),
			MeterName VARCHAR(255),
			MeterRegion VARCHAR(255),
			Unit INT,
			ResourceLocation VARCHAR(255),
			ConsumedService VARCHAR(255),
			ResourceGroup VARCHAR(255),
			ResourceURI VARCHAR(255),
			ChargeType VARCHAR(255),
			UnitPrice DECIMAL(10, 2),
			Quantity DECIMAL(10, 2),
			UnitType VARCHAR(255),
			BillingPreTaxTotal DECIMAL(10, 2),
			BillingCurrency VARCHAR(255),
			PricingPreTaxTotal DECIMAL(10, 2),
			PricingCurrency VARCHAR(255),
			ServiceInfo1 VARCHAR(255),
			ServiceInfo2 VARCHAR(255),
			Tags VARCHAR(255),
			AdditionalInfo VARCHAR(255),
			EffectiveUnitPrice DECIMAL(10, 2),
			PCToBCExchangeRate DECIMAL(10, 2),
			PCToBCExchangeRateDate DATE,
			EntitlementId VARCHAR(255),
			EntitlementDescription VARCHAR(255),
			PartnerEarnedCreditPercentage DECIMAL(10, 2),
			CreditPercentage DECIMAL(10, 2),
			CreditType VARCHAR(255),
			BenefitOrderId VARCHAR(255),
			BenefitId VARCHAR(255),
			BenefitType VARCHAR(255)
		)
	`)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return err
	}

	return nil
}

func InsertDataFromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		query := scanner.Text()
		query = strings.TrimSpace(query)

		if query != "" {
			_, err := db.Exec(query)
			if err != nil {
				log.Printf("Error executing query: %v", err)
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file: %v", err)
		return err
	}

	return nil
}
