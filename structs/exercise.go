package structs

import (
	"fmt"
	"os"
)

/*
	Your Tasks
		1) Create a new type / data structure for storing product data (e.g. about a book)
				The data structure should contain these fields:
				- ID
				- Title / Name
				- Short description
				- price (number without currency, we'll just assume USD)
		2) Create concrete instances of this data type in two different ways:
				- Directly inside the main function
				- Inside of main, by using a "creation helper function"
				(use any values for title etc. of your choice)
				Output (print) the created data structures in the command line (in the main function)
		3) Add a "connected function" that outputs the data + call that function from inside "main"
		4) Change the program to fetch user input values for the different data fields
				and create only one concrete instance with that entered data.
				Output that instance data (via the connected function) then.
		5) Bonus: Add a connected "store" function that writes that data into a file
				The filename should be the unique id, the function should be called at the
				end of main.
*/

type Product struct {
	Id               string
	Title            string
	ShortDescription string
	Price            float64
}

func (product *Product) DisplayProduct() {
	fmt.Printf("Book id: %v title:%v with description %v with price :$ %.2f\n\n ", product.Id, product.Title, product.ShortDescription, product.Price)
}

func (product *Product) Store() {
	file, _ := os.Create("structs/" + product.Id + ".txt")
	content := fmt.Sprintf("Book id: %v \ntitle:%v \ndescription %v \nprice :$ %.2f\n\n ", product.Id, product.Title, product.ShortDescription, product.Price)
	file.WriteString(content)
	file.Close()
}

func NewProduct(id string, title string, shortDesc string, price float64) *Product {
	product := Product{
		Id:               id,
		Title:            title,
		ShortDescription: shortDesc,
		Price:            price,
	}
	return &product
}
