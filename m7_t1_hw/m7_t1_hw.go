package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Address struct {
	PostalCode string
	City       string
	Street     string
	House      string
	Apartment  string
}

type Invoice struct {
	ItemName        string
	Quantity        int
	CustomerName    string
	ContactPhone    string
	DeliveryAddress Address
}

func NewInvoice(itemName string, quantity int, customerName, contactPhone, postalCode, city, street, house, apartment string) (*Invoice, error) {
	invoice := &Invoice{
		ItemName:     itemName,
		Quantity:     quantity,
		CustomerName: customerName,
		ContactPhone: contactPhone,
		DeliveryAddress: Address{
			PostalCode: postalCode,
			City:       city,
			Street:     street,
			House:      house,
			Apartment:  apartment,
		},
	}
	if err := invoice.Validate(); err != nil {
		return nil, err
	}
	return invoice, nil
}

func (inv *Invoice) Validate() error {
	if _, err := validateItemName(inv.ItemName); err != nil {
		return err
	}
	if _, err := validateQuantity(strconv.Itoa(inv.Quantity)); err != nil {
		return err
	}
	if _, err := validateCustomerName(inv.CustomerName); err != nil {
		return err
	}
	if _, err := validateContactPhone(inv.ContactPhone); err != nil {
		return err
	}
	if _, err := validatePostalCode(inv.DeliveryAddress.PostalCode); err != nil {
		return err
	}
	if _, err := validateAddressField(inv.DeliveryAddress.City, "city"); err != nil {
		return err
	}
	if _, err := validateAddressField(inv.DeliveryAddress.Street, "street"); err != nil {
		return err
	}
	if _, err := validateAddressField(inv.DeliveryAddress.House, "house"); err != nil {
		return err
	}
	if _, err := validateAddressField(inv.DeliveryAddress.Apartment, "apartment"); err != nil {
		return err
	}
	return nil
}

func (inv *Invoice) Display() {
	fmt.Printf("Item Name: %s\n", inv.ItemName)
	fmt.Printf("Quantity: %d\n", inv.Quantity)
	fmt.Printf("Customer Name: %s\n", inv.CustomerName)
	fmt.Printf("Contact Phone: %s\n", inv.ContactPhone)
	fmt.Printf("Address: %s, %s, %s, %s, %s\n", inv.DeliveryAddress.PostalCode, inv.DeliveryAddress.City, inv.DeliveryAddress.Street, inv.DeliveryAddress.House, inv.DeliveryAddress.Apartment)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var itemName, customerName, contactPhone, postalCode, city, street, house, apartment string
	var quantity int
	var err error

	for {
		fmt.Println("Enter item name (1-100 characters):")
		input, _ := readLine(reader)
		itemName, err = validateItemName(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter quantity (positive integer):")
		input, _ := readLine(reader)
		quantity, err = validateQuantity(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter customer name (only letters):")
		input, _ := readLine(reader)
		customerName, err = validateCustomerName(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter contact phone (exactly 10 digits):")
		input, _ := readLine(reader)
		contactPhone, err = validateContactPhone(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter postal code (exactly 6 digits):")
		input, _ := readLine(reader)
		postalCode, err = validatePostalCode(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter city:")
		input, _ := readLine(reader)
		city, err = validateAddressField(input, "city")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter street:")
		input, _ := readLine(reader)
		street, err = validateAddressField(input, "street")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter house number:")
		input, _ := readLine(reader)
		house, err = validateAddressField(input, "house")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	for {
		fmt.Println("Enter apartment number:")
		input, _ := readLine(reader)
		apartment, err = validateAddressField(input, "apartment")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	invoice, err := NewInvoice(itemName, quantity, customerName, contactPhone, postalCode, city, street, house, apartment)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Invoice created successfully:")
	invoice.Display()
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	return strings.TrimSpace(line), err
}

func validateItemName(itemName string) (string, error) {
	if itemName == "" || len(itemName) > 100 {
		return "", fmt.Errorf("invalid item name, must be between 1 and 100 characters")
	}
	return itemName, nil
}

func validateQuantity(input string) (int, error) {
	quantity, err := strconv.Atoi(input)
	if err != nil || quantity < 1 {
		return 0, fmt.Errorf("quantity must be a positive integer")
	}
	return quantity, nil
}

func validateCustomerName(name string) (string, error) {
	if match, _ := regexp.MatchString(`^[a-zA-Zа-яА-Я\s]+$`, name); !match {
		return "", fmt.Errorf("customer name must only contain letters")
	}
	return name, nil
}

func validateContactPhone(phone string) (string, error) {
	if match, _ := regexp.MatchString(`^\d{10}$`, phone); !match {
		return "", fmt.Errorf("phone number must be exactly 10 digits")
	}
	return phone, nil
}

func validatePostalCode(code string) (string, error) {
	if match, _ := regexp.MatchString(`^\d{6}$`, code); !match {
		return "", fmt.Errorf("postal code must be exactly 6 digits")
	}
	return code, nil
}

func validateAddressField(field, fieldName string) (string, error) {
	if strings.TrimSpace(field) == "" {
		return "", fmt.Errorf("%s must not be empty", fieldName)
	}
	return field, nil
}
