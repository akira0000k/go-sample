// STEP03: データの記録

package main

import "fmt"

// TODO:
// 品目と値段を一緒に扱うために
// Itemという構造体の型を定義する
// Categoryという品目を入れる文字列型のフィールドを持つ
// Priceという値段を入れる整数型のフィールドを持つ
type Item struct {
	Category string
	Price int
}

func main() {

	// TODO:
	// inputItemという関数を呼び出し
	// 結果をitemという変数に入れる
	item := inputItem()
	
	fmt.Println("===========")

	// TODO:
	// 品目に「コーヒー」、値段に「100」と入力した場合に
	// 「コーヒーに100円使いました」と表示する
	if item.Category == "coffee" {
		if item.Price == 100 {
			fmt.Println("コーヒーに100円使いました")
		}
	}
	fmt.Println("===========")
}

// 入力を行う関数
// 入力したItemを返す
func inputItem() Item {
	// Item型のitemという名前の変数を定義する
	var item Item

	fmt.Print("品目>")
	// TODO: 入力した値をitemのCategoryフィールドに入れる
	fmt.Scan(&item.Category)
	
	fmt.Print("値段>")
	// 入力した値をitemのPriceフィールドに入れる
	fmt.Scan(&item.Price)

	// TODO: itemを返す
	return item
}
// Akira@MBP step03 % go run main.go
// 品目>coffee
// 値段>100
// ===========
// コーヒーに100円使いました
// ===========
// Akira@MBP step03 % go build main.go
// Akira@MBP step03 % ls
// total 4280
// drwxr-xr-x   8 Akira  staff      256 11  5 16:45 ./
// drwxr-xr-x  15 Akira  staff      480 11  4 19:36 ../
// -rw-r--r--   1 Akira  staff       18 10 31 16:59 .gitignore
// -rw-r--r--   1 Akira  staff      173 10 31 16:59 README.md
// -rw-r--r--   1 Akira  staff       68 10 31 16:59 go.mod
// -rwxr-xr-x   1 Akira  staff  2167296 11  5 16:45 main*
// -rw-r--r--   1 Akira  staff     1246 11  5 16:44 main.go
// -rw-r--r--   1 Akira  staff     1017 10 31 16:59 main.go~
// Akira@MBP step03 % ./main
// 品目>spaghetti
// 値段>300
// ===========
// ===========
// Akira@MBP step03 % ./main
// 品目>coffee
// 値段>200
// ===========
// ===========
// Akira@MBP step03 % ./main
// 品目>coffee
// 値段>100
// ===========
// コーヒーに100円使いました
// ===========
// Akira@MBP step03 %
