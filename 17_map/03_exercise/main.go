package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text
	var (

		// #1
		// Key        : Last name
		// Element    : Phone number
		phones = map[string]string{
			"bowen": "202-555-0179",
			"dulin": "03.37.77.63.06",
			"greco": "03489940240",
		}

		// #2
		// Key        : Product ID
		// Element    : Available / Unavailable
		products = map[int]bool{
			617841573: true,
			879401371: false,
			576872813: true,
		}

		// #3
		// Key        : Last name
		// Element    : Phone numbers
		multiPhones = map[string][]string{
			"bowen": {"202-555-0179"},
			"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
			"greco": {"03489940240", "03489900120"},
		}

		// #4
		// Key        : Customer ID
		// Element Key:
		//   Key: Product ID Element: Quantity
		basket = map[int]map[int]int{
			100: {617841573: 4, 576872813: 2},
			101: {576872813: 5, 657473833: 20},
			102: {},
		}
	)

	// ---------------------------------------------------------
	// EXERCISE II: Populate and Lookup
	//
	//  Add elements to the maps that you've declared in the
	//  first exercise, and try them by looking up for the keys.
	//
	//  Either use the `make()` or `map literals`.
	//
	//  After completing the exercise, remove the data and check
	//  that your program still works.
	//
	//
	//  1. Phone numbers by last name
	//     --------------------------
	//     bowen  202-555-0179
	//     dulin  03.37.77.63.06
	//     greco  03489940240
	//
	//     Print the dulin's phone number.

	who, phone := "dulin", "N/A"

	if value, ok := phones["dulin"]; ok {
		phone = value
	}
	fmt.Printf("Phone Number of %s: %s\n", who, phone)

	//  2. Product availability by Product ID
	//     ----------------
	//     617841573 true
	//     879401371 false
	//     576872813 true
	//
	//     Is Product ID 879401371 available?
	id, status := 87940371, "available"

	if !products[id] {
		status = "not " + status
	}
	fmt.Printf("Product ID %d is %s\n", id, status)

	//  3. Multiple phone numbers by last name
	//     ------------------------------------------------------
	//     bowen  [202-555-0179]
	//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
	//     greco  [03489940240 03489900120]
	//
	//     What is Greco's second phone number?
	who, phone = "greco", "N/A"

	if phones := multiPhones[who]; len(phones) >= 2 {
		phone = phones[1]
	}
	fmt.Printf("%s's second phone number: %s\n", who, phone)

	//  4. Shopping basket by Customer ID
	//     -------------------------------
	//     100 [617841573:4 576872813:2]
	//     101 [576872813:5 657473833:20]
	//     102 []
	//
	//     How many of 576872813 the customer 101 is going to buy?
	//                (Product ID)  (Customer ID)
	cid, pid := 101, 576872813

	fmt.Printf("Customer #%d is going to buy %d from Product #%d.\n", cid, basket[cid][pid], pid)

	// EXPECTED OUTPUT
	//
	//   1. Run the solution to see the output
	//   2. Here is the output with empty maps:
	//
	//      dulin's phone number: N/A
	//      Product ID #879401371 is not available
	//      greco's 2nd phone number: N/A
	//      Customer #101 is going to buy 5 from Product ID #576872813.
	//
	// ---------------------------------------------------------

	fmt.Println()
	fmt.Printf("phones     : %#v\n", phones)
	fmt.Printf("products   : %#v\n", products)
	fmt.Printf("multiPhones: %#v\n", multiPhones)
	fmt.Printf("basket     : %#v\n", basket)
}
