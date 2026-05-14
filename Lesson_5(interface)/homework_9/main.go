package main

import "fmt"

type Displayable interface {
	Display()
}

type BankAccount struct {
	AccountNumber string
	OwnerName     string
	Balance       float64
	AccountType   string
	Currency      string
}

func (b BankAccount) Display() {
	fmt.Printf("Номер счета: %d, Владелец: %s, Баланс: %.2f, Тип счета: %s, Валюта: %s",
		b.AccountNumber, b.OwnerName, b.Balance, b.AccountType, b.Currency)
}

func addBankAccount(bankAccounts []BankAccount, bankAccount BankAccount) []BankAccount {
	return append(bankAccounts, bankAccount)
}

func filterByBalance(bankAccounts []BankAccount, minBalance float64) []BankAccount {
	var filteredBankAccounts []BankAccount
	for _, b := range bankAccounts {
		if b.Balance >= minBalance {
			filteredBankAccounts = append(filteredBankAccounts, b)
		}
	}
	return filteredBankAccounts
}

func FilterByCurrency(bankAccounts []BankAccount, currency string) []BankAccount {
	var filteredBankAccounts []BankAccount
	for _, b := range bankAccounts {
		if b.Currency == currency {
			filteredBankAccounts = append(filteredBankAccounts, b)
		}
	}
	return filteredBankAccounts
}

func Deposit(account *BankAccount, amount int) {
	if amount < 0 {
		fmt.Println("Сумма для пополнения должна быть больше ноля")
	}
	account.Balance += float64(amount)
}

func Withdraw(account *BankAccount, amount int) {
	if float64(amount) > account.Balance {
		fmt.Println("Недостаточно средств для списания")
		return
	}
	account.Balance -= float64(amount)
}

func printDisplayable(bankAccounts []Displayable) {
	for _, b := range bankAccounts {
		b.Display()
	}
}

func convertorByDisplayable(bankAccounts []BankAccount) []Displayable {
	var displayable []Displayable
	for _, b := range bankAccounts {
		displayable = append(displayable, b)
	}
	return displayable
}

func FindAccount(accounts []BankAccount, number string) *BankAccount {
	for i := range accounts {
		if accounts[i].AccountNumber == number {
			return &accounts[i]
		}
	}
	return nil
}

func main() {
	bankAccounts := []BankAccount{
		{AccountNumber: "DE1234567890", OwnerName: "Ivan Ivanov", Balance: 1500.50, AccountType: "DEBIT", Currency: "EUR"},
		{AccountNumber: "US9876543210", OwnerName: "John Smith", Balance: 3200.00, AccountType: "CREDIT", Currency: "USD"},
		{AccountNumber: "RU555666777", OwnerName: "Sergey Petrov", Balance: 87000.75, AccountType: "SAVINGS", Currency: "RUB"},
		{AccountNumber: "DE1122334455", OwnerName: "Anna Müller", Balance: 540.20, AccountType: "DEBIT", Currency: "EUR"},
		{AccountNumber: "US2233445566", OwnerName: "Emily Davis", Balance: 10200.99, AccountType: "BUSINESS", Currency: "USD"},
		{AccountNumber: "RU888999000", OwnerName: "Alexey Smirnov", Balance: 43000.00, AccountType: "DEBIT", Currency: "RUB"},
	}

	newAccount := BankAccount{AccountNumber: "DE6677889900", OwnerName: "Laura Schmidt", Balance: 7600.10, AccountType: "SAVINGS", Currency: "EUR"}
	bankAccounts = addBankAccount(bankAccounts, newAccount)

	accountForDeposit := FindAccount(bankAccounts, "ACC001")
	if accountForDeposit != nil {
		Deposit(accountForDeposit, 500)
	}

	accountForWithdraw := FindAccount(bankAccounts, "ACC002")
	if accountForWithdraw != nil {
		Withdraw(accountForWithdraw, 1000)
	}

	fmt.Println("Полный список счетов: ")
	printDisplayable(convertorByDisplayable(bankAccounts))

	var minBalance float64
	var currency string

	fmt.Print("Введите минимальный баланс, для фильтрации: ")
	fmt.Scanln(&minBalance)

	fmt.Print("Введите валюту: ")
	fmt.Scanln(&currency)

	filteredBalance := filterByBalance(bankAccounts, minBalance)
	filteredCurrency := FilterByCurrency(bankAccounts, currency)

	fmt.Printf("Счета с балансом от %.2f", minBalance)
	printDisplayable(convertorByDisplayable(filteredBalance))

	fmt.Printf("Счета с валютой %s", currency)
	printDisplayable(convertorByDisplayable(filteredCurrency))
}
