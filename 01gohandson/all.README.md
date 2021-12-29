==> step01/README.md <==
# STEP01: Goに触れる

## 新しく学ぶこと

* Goのプログラムの書き方
* Goのプログラムの実行の仕方
* 文字列の表示方法

## 動かし方

```
$ go build -v -o step01
$ ./step01
```

※ Windowsの方は以下のように後ろに.exeつけてください。
他のステップでも同様につけておいてください。

```
$ go build -v -o step01.exe
$ .¥step01.exe
```

==> step02/README.md <==
# STEP02: データの入力

## 新しく学ぶこと

* 標準パッケージの使い方
* 変数と型
* fmt.Printを使った表示
* fmt.Scanを使った入力
* fmt.Printlnを使った表示
* fmt.Printfを使った表示

## 動かし方

```
$ go build -v -o step02
$ ./step02
```

==> step03/README.md <==
# STEP03: データの記録

## 新しく学ぶこと

* 構造体
* 型の定義
* 関数
* :=による代入

## 動かし方

```
$ go build -v -o step03
$ ./step03
```

==> step04/README.md <==
# STEP04: 複数データの記録

## 新しく学ぶこと

* スライス
* スライスの初期化
* 容量の取得
* スライスの追加
* スライスの長さの取得
* スライスのi番目の要素のアクセス方法
* 関数の引数

## 動かし方

```
$ go build -v -o step04
$ ./step04
```

==> step05/README.md <==
# STEP05: ファイルへの保存

## 新しく学ぶこと

* ファイルへの保存
* defer
* エラー処理
* bufio.Scanner
* strings.Split
* strconv.Atoi

## 動かし方

```
$ go build -v -o step05
$ ./step05
```

==> step06/README.md <==
# STEP06: ブラッシュアップ

## 新しく学ぶこと

* メソッド
* 複数の戻り値
* fmt.Fprintln
* os.Stderr
* os.Exit
* ポインタ
* 無限ループ
* ラベルつきbreak
* switch
* for range
* スライス演算

## 動かし方

```
$ go build -v -o step06
$ ./step06
```

==> step07/README.md <==
# STEP07: データベースへの記録

## 新しく学ぶこと

* サードパーティパッケージの使い方
* database/sqlパッケージの使い方
 * テーブルの作成
 * INSERT
 * SELECT
* fmt.Printfの%04d

## 動かし方

```
$ go build -v -o step07
$ ./step07
```

==> step08/README.md <==
# STEP08: 品目ごとの集計

## 新しく学ぶこと

* GROUP BYと集約関数（sum, count）
* fmt.Printfの\tと%f
* float64
* キャスト

## 動かし方

```
$ go build -v -o step08
$ ./step08
```

==> step09/README.md <==
# STEP09: 一覧ページの作成

## 新しく学ぶこと

* HTTPハンドラ
* HTTPサーバの起動
* html/templateの使い方

## 動かし方

`accountbook.db`はSTEP07のプログラムで作ったものを利用しましょう。

```
$ go build -v -o step09
$ ./step09
```

==> step10/README.md <==
# STEP10: 入力ページの作成

## 新しく学ぶこと

* HTMLのform
* HTTPメソッドの取得
* POSTされたデータの取得
* リダイレクト

## 動かし方

```
$ go build -v -o step10
$ ./step10
```

==> step11/README.md <==
# STEP11: 集計ページの作成

## 新しく学ぶこと

* Google Chart API


```
$ go build -v -o step11
$ ./step11
```
