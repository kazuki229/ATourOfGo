package main

import (
	"fmt"
	"math"
	"strings"
)

// golang.org/x/tour/picのようにgo

// 構造体
type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	p := &i         // iのポインタ
	fmt.Println(*p) // ポインタpの値を参照
	*p = 21         // ポインタpの値をセット
	fmt.Println(i)  // iの値が更新される

	p = &j         // jのポインタ
	*p = *p / 37   // pの値を代入
	fmt.Println(j) // jの値が更新される

	v := Vertex{1, 2}
	// フィールドへは"."でアクセス
	v.X = 4
	fmt.Println(v.X)

	// 構造体のポインタも使用可能
	p2 := &v
	p2.X = 1e9
	fmt.Println(v)

	// 構造体リテラル
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p3 = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p3, v2, v3)

	// 配列の宣言
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 配列の初期化(固定長)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// スライス(可変長)
	var s []int = primes[1:4]
	fmt.Println(s)

	// スライスは配列への参照
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a2 := names[0:2]
	b2 := names[1:3]
	fmt.Println(a2, b2)

	// namesのスライスであるb2を更新すると
	// namesが更新される
	b2[0] = "XXX"
	fmt.Println(a2, b2)
	fmt.Println(names)

	// 以下の配列で
	// var a [10]int32
	// 下記のスライスは等価
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]

	s2 := []int{2, 3, 5, 7, 11, 13}

	s2 = s2[1:4]
	fmt.Println(s2)

	s2 = s2[:2]
	fmt.Println(s2)

	s2 = s2[1:]
	fmt.Println(s2)

	// スライスはlen(長さ)とcap(容量)を持つ
	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice2(s3)

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	printSlice2(s3)

	// Extend its length.
	s3 = s3[:4]
	printSlice2(s3)

	// Drop its first two values.
	s3 = s3[2:]
	printSlice2(s3)

	// ズライスのゼロ値はnil
	var s4 []int
	fmt.Println(s4, len(s4), cap(s4))
	if s4 == nil {
		fmt.Println("nil!")
	}

	// スライスはmake関数で初期化可能
	a3 := make([]int, 5)
	printSlice("a", a3)

	// (型, len, cap)
	b3 := make([]int, 0, 5)
	printSlice("b", b3)

	c3 := b3[:2]
	printSlice("c", c3)

	d3 := c3[2:5]
	printSlice("d", d3)

	// スライスは任意の型を含められる
	// スライスも
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to slice
	var s5 []int
	printSlice2(s5) // lent=0 cap=0 []

	s5 = append(s5, 0)
	printSlice2(s5) // len=1 cap=1 [0]

	s5 = append(s5, 1)
	printSlice2(s5) // len=2 cap=2 [0 1]

	// appendしたときにlenがcapを超えるとcapは元のcapの2倍となる
	s5 = append(s5, 2, 3, 4)
	printSlice2(s5) // len=5 cap=6 [0 1 2 3 4]

	// Rance
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// forループで利用するrangeはスライスやマップを一つずつ反復処理するために用いる
	// スライスをrangeで繰り返す場合、1つめの変数i2はインデックス、v4はインデックスの場所の要素をコピーとなる
	for i2, v4 := range pow {
		fmt.Printf("2**%d = %d\n", i2, v4)
	}

	pow2 := make([]int, 10)
	// インデックスだけ必要な場合はvalueを省略可能
	for i3 := range pow2 {
		pow2[i3] = 1 << uint(i3) // == 2**i
	}

	// valueのみ必要な場合はindexを_で置き換える
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	//pic.Show(Pic)

	// Maps
	// 連想配列
	// [キーの型]値の型
	var m map[string]Vertex2
	// make関数は指定した型の初期化されたmapを返す
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals
	// リテラルの場合はキーが必要
	var m2 = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	// Map literals continued
	// mapにわたすトップレベルの型が単純なものである場合、リテラルから型推定できるので型名を省略可能
	var m3 = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	// Mutating Maps
	m4 := make(map[string]int)
	// 要素の挿入
	m4["Answer"] = 42
	fmt.Println("The value: ", m4["Answer"]) // 42

	// 要素の更新
	m4["Answer"] = 48
	fmt.Println("The value: ", m4["Answer"]) // 48

	// 要素の削除
	delete(m4, "Answer")
	fmt.Println("The value: ", m4["Answer"]) // 0

	// 要素の取得
	// 2つめはキーが歩かないかをboolで返す
	v5, ok := m4["Answer"]
	fmt.Println("The value: ", v5, "Present?", ok) //0 false

	// Exercise: Maps

	// Function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	// 関数を変数として扱えるので、関数の引数に取れる
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Function Closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// Exercise: Fibonacci closure
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//func Pic(dx, dy int) [][]uint8 {
//	return
//}

type Vertex2 struct {
	Lat, Long float64
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
