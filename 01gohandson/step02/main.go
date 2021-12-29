// STEP02: データの入力

package main

// fmtパッケージをインポートする
import "fmt"

func main() {
	// TODO: 品目を入れる変数を定義
	var category string
	
	// 値段を入れる変数を定義
	var price int

	// 「品目>」と表示する
	fmt.Print("品目>")
	// 入力した結果をcategoryに入れる
	fmt.Scan(&category)

	// TODO: 「値段>」と表示する
	fmt.Print("値段>")
	// TODO: 入力した結果をpriceに入れる
	fmt.Scan(&price)
	
	// 「===========」と出力して改行する
	fmt.Println("===========")

	// 品目に「コーヒー」、値段に「100」と入力した場合に
	// 「コーヒーに100円使いました」と表示する
	fmt.Printf("%sに%d円使いました\n", category, price)

	// 「===========」と出力して改行する
	fmt.Println("===========")
}
//Akira@MBP step02 % go run main.go
//品目>coffee
//値段>300
//===========
//coffeeに300円使いました
//===========
//Akira@MBP step02 %
