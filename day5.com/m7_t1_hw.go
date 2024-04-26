package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == "GET" {
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
<title>Invoice Form</title>
</head>
<body>
<h2>Invoice Input Form</h2>
<form method="post">
Item Name: <input type="text" name="itemName"><br>
Quantity: <input type="number" name="quantity"><br>
Customer Name: <input type="text" name="customerName"><br>
Contact Phone: <input type="text" name="contactPhone"><br>
Postal Code: <input type="text" name="postalCode"><br>
City: <input type="text" name="city"><br>
Street: <input type="text" name="street"><br>
House: <input type="text" name="house"><br>
Apartment: <input type="text" name="apartment"><br>
<input type="submit" value="Submit">
</form>
</body>
</html>`)
	} else if r.Method == "POST" {
		r.ParseForm()
		itemName := r.FormValue("itemName")
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))
		customerName := r.FormValue("customerName")
		contactPhone := r.FormValue("contactPhone")
		postalCode := r.FormValue("postalCode")
		city := r.FormValue("city")
		street := r.FormValue("street")
		house := r.FormValue("house")
		apartment := r.FormValue("apartment")

		invoice, err := NewInvoice(itemName, quantity, customerName, contactPhone, postalCode, city, street, house, apartment)
		if err != nil {
			fmt.Fprintf(w, "Error creating invoice: %v", err)
			return
		}

		// Вывод информации о накладной
		fmt.Fprintf(w, "Invoice created successfully:<br>")
		fmt.Fprintf(w, "Item Name: %s<br>Quantity: %d<br>Customer Name: %s<br>Contact Phone: %s<br>Address: %s, %s, %s, %s, %s",
			invoice.ItemName, invoice.Quantity, invoice.CustomerName, invoice.ContactPhone, invoice.DeliveryAddress.PostalCode,
			invoice.DeliveryAddress.City, invoice.DeliveryAddress.Street, invoice.DeliveryAddress.House, invoice.DeliveryAddress.Apartment)
	}
}

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
