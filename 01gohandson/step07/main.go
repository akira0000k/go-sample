// STEP07: データベースへの記録

package main

import (
	"fmt"
	"os"
	// TODO:
	// SQLiteのドライバを使うために
	"database/sql"
	"github.com/tenntenn/sqlite"
)

func main() {

	// TODO:
	// データベースへ接続
	// ドライバにはSQLiteを使って、
	// ドライバ名はsqlite.DriverName
	// accountbook.dbというファイルでデータベース接続を行う
	db, err := sql.Open(sqlite.DriverName, "../accountbook.db")
	if err != nil {
		// 標準エラー出力（os.Stderr)にエラーメッセージを出力して終了
		fmt.Fprintln(os.Stderr, "エラー：", err)
		// ステータスコードを1で終了
		os.Exit(1)
	}

	// AccountBookをNewAccountBookを使って作成
	ab := NewAccountBook(db)

	// テーブルを作成
	if err := ab.CreateTable(); err != nil {
		// 標準エラー出力（os.Stderr)にエラーメッセージを出力して終了
		fmt.Fprintln(os.Stderr, "エラー：", err)
		// ステータスコードを1で終了
		os.Exit(1)
	}

LOOP: // 以下のループにラベル「LOOP」をつける
	for {

		// モードを選択して実行する
		var mode int
		fmt.Println("[1]入力 [2]最新10件 [3]終了")
		fmt.Printf(">")
		fmt.Scan(&mode)

		// モードによって処理を変える
		switch mode {
		case 1: // 入力
			var n int
			fmt.Print("何件入力しますか>")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				if err := ab.AddItem(inputItem()); err != nil {
					fmt.Fprintln(os.Stderr, "エラー:", err)
					break LOOP
				}
			}
		case 2: // 最新10件
			items, err := ab.GetItems(10)
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー:", err)
				break LOOP
			}
			showItems(items)
		case 3: // 終了
			fmt.Println("終了します")
			return
		}
	}
}

// Itemを入力し返す
func inputItem() *Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	return &item
}

// Itemの一覧を出力する
func showItems(items []*Item) {
	fmt.Println("===========")
	// itemsの要素を1つずつ取り出してitemに入れて繰り返す
	for _, item := range items {
		fmt.Printf("[%04d] %s:%d円\n", item.ID, item.Category, item.Price)
	}
	fmt.Println("===========")
}
// -*- mode: compilation; default-directory: "~/go/src/sample/01gohandson/step07/" -*-
// Compilation started at Sun Nov  7 00:12:08
//  
// go run main.go
// go: downloading github.com/tenntenn/sqlite v1.0.2
// go: downloading github.com/mattn/go-sqlite3 v1.10.0
// # command-line-arguments
// ./main.go:10:2: imported and not used: "github.com/tenntenn/sqlite"
//
// /Users/Akira/go/pkg/mod/github.com:
// total used in directory 0 available 123.3 GiB
// drwxr-xr-x	5 Akira	 staff	160 11	7 00:12 mattn
// drwxr-xr-x	3 Akira	 staff	 96 11	7 00:12 tenntenn

// Akira@MBP step07 % ls
// total 13000
// drwxr-xr-x  11 Akira  staff      352 11  7 21:40 ./
// drwxr-xr-x  15 Akira  staff      480 11  4 19:36 ../
// -rw-r--r--   1 Akira  staff      287 10 31 16:59 README.md
// -rw-r--r--   1 Akira  staff     2165 11  7 21:40 accountbook.go
// -rw-r--r--   1 Akira  staff     2165 11  7 21:35 accountbook.go~
// -rw-r--r--   1 Akira  staff      111 10 31 16:59 go.mod
// -rw-r--r--   1 Akira  staff     1722 10 31 16:59 go.sum
// -rwxr--r--   1 Akira  staff       35 11  7 21:36 gobuild.sh*
// -rw-r--r--   1 Akira  staff     2809 11  7 21:40 main.go
// -rw-r--r--   1 Akira  staff     2793 11  7 21:31 main.go~
// -rwxr-xr-x   1 Akira  staff  6620064 11  7 21:40 step07*
// Akira@MBP step07 % pwd
// /Users/Akira/go/src/sample/01gohandson/step07
// Akira@MBP step07 % ./step07
// [1]入力 [2]最新10件 [3]終了
// >1
// 何件入力しますか>2
// 品目>coffee
// 値段>300
// 品目>teabag
// 値段>1000
// [1]入力 [2]最新10件 [3]終了
// >2
// ===========
// [0002] teabag:1000円
// [0001] coffee:300円
// ===========
// [1]入力 [2]最新10件 [3]終了
// >3
// 終了します
// Akira@MBP step07 % ./step07 
// [1]入力 [2]最新10件 [3]終了
// >2
// ===========
// [0002] teabag:1000円
// [0001] coffee:300円
// ===========
// [1]入力 [2]最新10件 [3]終了
// >3
// 終了します
// Akira@MBP step07 %
