package main

import "fmt"

func main(){
	fmt.Print("Enter a number to begin with: ")
	var num int64
	var value byte
	var n uint8

	for {
		_, err := fmt.Scan(&num)
		if err != nil{
			fmt.Println("Wrong input! Try again: ")
		} else {
			break
		}
	}

	for {
		fmt.Printf("Your current number: %d, which is %064b\n", num, uint64(num))
		fmt.Println("_______________________________________________________________________________")
		fmt.Print("Choose a digit to insert (0 or 1): ")
		_, err := fmt.Scan(&value)
		if err != nil || !(value == 0 || value ==1){
			fmt.Println("Wrong input! You can only insert 0 or 1. Try again.")
			continue
		}

		fmt.Print("Choose bit to insert (from 0 to 63): ")
		_, err = fmt.Scan(&n)
		if err != nil || n > 63{
			fmt.Println("Wrong input! You can only insert number from 0 to 63 Try again.")
			continue
		}

		switch value{
		case 1:
			num = num | (1 << n)
		default:
			num = num &^ (1 << n)
		}
	}
}