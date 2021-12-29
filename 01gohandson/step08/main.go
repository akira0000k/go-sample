// STEP08: 品目ごとの集計

package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/tenntenn/sqlite"
)

func main() {

	// データベースへ接続
	// ドライバにはSQLiteを使って、
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
		fmt.Println("[1]入力 [2]最新10件 [3]集計 [4]終了")
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
		case 3: // 集計
			summaries, err := ab.GetSummaries()
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー:", err)
				break LOOP
			}
			showSummary(summaries)
		case 4: // 終了
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

// 集計を出力する
func showSummary(summaries []*Summary) {
	fmt.Println("===========")
	// タブ区切りで「品目 個数 合計 平均」を出力
	fmt.Printf("品目\t個数\t合計\t平均\n")
	// summariesの要素を1つずつ取り出してsに入れて繰り返す
	for _, s := range summaries {
		// TODO: 第2引数以降に品目、個数、合計、平均を渡す
		fmt.Printf("%s\t%d\t%d円\t%.2f円\n", s.Category, s.Count, s.Sum, s.Avg())
	}
	fmt.Println("===========")
}

//type Summary struct {
// 	Category string
// 	Count    int
// 	Sum      int
//}


// Akira@MBP step08 % ./step08
// [1]入力 [2]最新10件 [3]集計 [4]終了
// >2
// ===========
// ===========
// [1]入力 [2]最新10件 [3]集計 [4]終了
// >1
// 何件入力しますか>2
// 品目>coffee
// 値段>300
// 品目>tea
// 値段>200
// [1]入力 [2]最新10件 [3]集計 [4]終了
// >3
// ===========
// 品目	個数	合計	平均
// coffee	1	300円	300.00円
// tea	1	200円	200.00円
// ===========
// [1]入力 [2]最新10件 [3]集計 [4]終了
// >1
// 何件入力しますか>2
// 品目>coffee
// 値段>100
// 品目>tea
// 値段>500
// [1]入力 [2]最新10件 [3]集計 [4]終了
// >2
// ===========
// [0004] tea:500円
// [0003] coffee:100円
// [0002] tea:200円
// [0001] coffee:300円
// ===========
// [1]入力 [2]最新10件 [3]集計 [4]終了
// >3
// ===========
// 品目	個数	合計	平均
// coffee	2	400円	200.00円
// tea	2	700円	350.00円
// ===========
// [1]入力 [2]最新10件 [3]集計 [4]終了
// >4
// 終了します
// Akira@MBP step08 %
