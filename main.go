package main

import "fmt"

func main() {
	fmt.Println("Data Collections:Arrays, Slices & Maps ")
	// array based on same type
	prices := [4]float64{10.2, 5.7, 12, 54.5}
	fmt.Println("Array: ", prices)
	nums := [5]int{1, 2, 3, 4, 5}
	for i, item := range nums {
		fmt.Println(i, item)
	}

}

/*
func main() {

	//var ch1 = make(chan string)
	//go channels.Sell(ch1)
	//go channels.Buy(ch1)
	//time.Sleep(2 * time.Second)

	//var wg sync.WaitGroup
	//wg.Add(2)
	//ch := make(chan int, 10)
	//go channels.SellRange(ch, &wg)
	//go channels.BuyRange(ch, &wg)
	//wg.Wait()

	//ch := make(chan int, 10)
	//ch <- 10
	//ch <- 11
	//ch <- 12
	//val, ok := <-ch
	//fmt.Println(val, ok)
	//close(ch)
	//val, ok = <-ch
	//fmt.Println(val, ok)
	//val, ok = <-ch
	//fmt.Println(val, ok)
	//val, ok = <-ch
	//fmt.Println(val, ok)

	//var wg sync.WaitGroup
	//wg.Add(2)
	//ch2 := make(chan int, 3)
	//go channels.SellBuff(ch2, &wg)
	//wg.Wait()

	//start := time.Now()
	//routines.Calulate()
	//time.Sleep(2 * time.Second)
	//elapsed := time.Since(start)

	//routines.Wg.Wait()
	//fmt.Println("Function Elapsed :", elapsed)

	//go routines.Start()
	//time.Sleep(1 * time.Second)

	// anonymous function
	//func() {
	//	fmt.Println("In anonymous method")
	//	fmt.Println(runtime.NumCPU())
	//}()
	//time.Sleep(1 * time.Second)
}
*/

/*
func main()  {
	fmt.Println("Struct")
	firstName := structs.GetData("First Name:")
	lastName := structs.GetData("Last Name:")
	birthDate := structs.GetData("Birth Date (mm/dd/YYYY):")

	var user1 structs.User = structs.User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}

	//var user2 structs.User = structs.User{firstName, lastName, birthDate, time.Now()}
	//var user3 = *structs.NewUser(firstName, lastName, birthDate)
	user1.OutPutUserDetails()
	//fmt.Println(user1)
	//fmt.Println(user3)

	fmt.Println("\n Struct Exercise")
	fmt.Println("----------------")
	product1 := structs.Product{
		Id:               "E45E478",
		Title:            "book title",
		ShortDescription: "book short description",
		Price:            300.56,
	}
	fmt.Println("product 1 :", product1)
	product1.DisplayProduct()

	product2 := *structs.NewProduct("445KE", "book title 2", "short description", 200.98)
	fmt.Println("product 2 :", product2)

	title := structs.GetData("Book title:")
	shortdesc := structs.GetData("Description:")
	price := structs.GetData("Price:")
	priceInFloat, _ := strconv.ParseFloat(price, 64)
	product3 := structs.Product{
		Id:               "E489KQ",
		Title:            title,
		ShortDescription: shortdesc,
		Price:            priceInFloat,
	}
	product3.DisplayProduct()
	product3.Store()
}
*/

/*
func main() {
	fmt.Println("pointers!")
	//fmt.Println(pointers.Age)
	//fmt.Println(pointers.PointerAge)
	//fmt.Println(*pointers.PointerAge)
	//*pointers.PointerAge = *pointers.PointerAge + 1
	//fmt.Println(pointers.PointerAge)
	//fmt.Println(*pointers.PointerAge)
	//fmt.Println(pointers.DoubleAge)
	//fmt.Println(pointers.Age)

}
*/
