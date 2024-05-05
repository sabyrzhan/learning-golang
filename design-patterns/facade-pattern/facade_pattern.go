package facade_pattern

import (
	"errors"
	"fmt"
)

/*
Facade is implement using P2P payment system example.
Money transfer from one account to another account is not a simple process and the workflow is complicated.
To make a transaction various subsystems are involved along the way such as:
- AccountService for validation account and balance validation
- CurrencyService for converting from foreign currency to local currency
- BankingService to transfer money from one account to another account
- DatabaseService to persist any transaction related data
- NotificationService to send confirmation or failed transaction notifications to user
- 3DSecureService to validate OTP code used as confirmation for payment transaction
*/

// ======== ACCOUNT SERVICE ===========

type AccountService interface {
	Exists(user string) bool
}

type AccountAPI struct {}
func (a *AccountAPI) Exists(user string) bool {
	return true
}

// ======== CURRENCY SERVICE ===========

type Currency int
const (
	USD Currency = iota
	EUR
	RUR
	KZT
)

type CurrencyService interface {
	ConvertToKZT(amount float32, currency Currency) float32
}

type CurrencyAPI struct {}
func (c *CurrencyAPI) ConvertToKZT(amount float32, currency Currency) float32 {
	return amount
}

// ======== BANKING SERVICE ===========

type BankingService interface {
	TransferMoney(fromUser string, toUser string, amount float32) bool
}
type BankingAPI struct {}
func (b *BankingAPI) TransferMoney(fromUser string, toUser string, amount float32) bool {
	return true
}

// ======== DATABASE SERVICE ===========

type DatabaseRepository interface {
	SaveTransaction(fromUser string, toUser string, amount float32) bool
}
type DB struct {}
func (db *DB) SaveTransaction(fromUser string, toUser string, amount float32) bool {
	return true
}

// ======== NOTIFICATION SERVICE ===========

type NotificationService interface {
	SendNotification(user string, message string) bool
}
type NotificationAPI struct {}
func (n *NotificationAPI) SendNotification(user string, message string) bool {
	return true
}

// ======== USER 3D SECURITY SERVICE ===========

type User3DSecurityService interface {
	verify(otpCode string) bool
}
type User3DSecurityAPI struct {}
func (u *User3DSecurityAPI) verify(otpCode string) bool {
	return true
}

// ======== P2P PAYMENT FACADE  ===========

type P2PPayment struct {
	accountService 			AccountService
	currencyService 		CurrencyService
	bankingService  		BankingService
	databaseRepository 		DatabaseRepository
	notificationService 	NotificationService
	user3DSecurityService 	User3DSecurityService
}

type P2PPaymentFacade interface{
	SendPayment(fromUser string, toUser string, amount float32, currency Currency, otpCode string) error
}

func NewP2PPaymentFacade() P2PPaymentFacade {
	return &P2PPayment {
		&AccountAPI{},
		&CurrencyAPI{},
		&BankingAPI{},
		&DB{},
		&NotificationAPI{},
		&User3DSecurityAPI{},
	}
}

func (p *P2PPayment) SendPayment(fromUser string, toUser string, amount float32, currency Currency, otpCode string) error {
	if !p.accountService.Exists(fromUser) {
		return errors.New("from user does not exist")
	}
	fmt.Println(fmt.Sprintf("From user validated successfully"))

	if !p.accountService.Exists(toUser) {
		return errors.New("to user does not exist")
	}

	fmt.Println(fmt.Sprintf("To user validated successfully"))

	if amount < 0 {
		return errors.New("amount must be greater than zero")
	}

	fmt.Println(fmt.Sprintf("Amount validated successfully"))

	kztAmount := amount
	if currency != KZT {
		kztAmount = p.currencyService.ConvertToKZT(amount, currency)

	}
	if kztAmount <= 100 {
		return errors.New("amount must be greater than than or equal to 100 KZT")
	}

	fmt.Println(fmt.Sprintf("Converted amount is %f KZT and is valid", kztAmount))

	if !p.user3DSecurityService.verify(otpCode) {
		return errors.New("invalid OTP code")
	}

	fmt.Println(fmt.Sprintf("3D sceure OTP code is valid"))

	bankingResult := p.bankingService.TransferMoney(fromUser, toUser, amount)
	if !bankingResult {
		return errors.New("failed to transfer money")
	}

	fmt.Println(fmt.Sprintf("Money successfully transferred from banking API"))

	saveResult := p.databaseRepository.SaveTransaction(fromUser, toUser, amount)
	if !saveResult {
		return errors.New("failed to save transaction")
	}

	fmt.Println(fmt.Sprintf("Trnasaction successfully saved to DB"))

	p.notificationService.SendNotification(fromUser, fmt.Sprintf("Money with amount %f transferred successfully", amount))

	fmt.Println(fmt.Sprintf("Notification send successully!"))

	fmt.Println(fmt.Sprintf("P2P send finished!"))


	return nil
}


