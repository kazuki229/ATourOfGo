package main

import "fmt"

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
	printSlice(s3)

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	printSlice(s3)

	// Extend its length.
	s3 = s3[:4]
	printSlice(s3)

	// Drop its first two values.
	s3 = s3[2:]
	printSlice(s3)

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
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
