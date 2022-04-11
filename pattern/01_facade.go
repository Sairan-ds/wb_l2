package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"fmt"
)

/* Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе
классов, библиотеке или фреймворку. */
/* Фасад — это простой интерфейс для работы со сложной подсистемой, содержащей множество классов. Фасад может
иметь урезанный интерфейс, не имеющий 100% функциональности, которой можно достичь, используя сложную подсистему напрямую.
Но он предоставляет именно те фичи, которые нужны клиенту, и скрывает все остальные. */

/* Ниже предоставлен пример реализации фасада. 



*/
//_______________________________________________________________________
// Wallet Facade

// Структура фасада
type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}
// Фасад должен будет позаботиться о том, чтобы правильно инициализировать объекты подсистемы. 
func NewWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Account created")
	return walletFacacde
}
// Добавление денег в кошелек, где пройдете проверка аккаунта на валидность, пин кода, добавление на счет, отправка уведомление об успешности и запись 
// в бухгалтерской книге
func (w *walletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}
// Снятие денег с кошелька, где пройдете проверка аккаунта на валидность, пин кода, наличие доступных средств и при наличии снятие, 
//отправка уведомление об успешности и запись в бухгалтерской книге
func (w *walletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}


//_______________________________________________________________________
// Account

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

//_______________________________________________________________________
// Ledger

type ledger struct {
}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
}

//_______________________________________________________________________
// Notification

type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

//_______________________________________________________________________
// SecurityCode

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

//_______________________________________________________________________
// Wallet

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

