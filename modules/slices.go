package modules

import "fmt"

func Slices() {
	var a1 [3]int
	var a2 = [3]int{10, 20, 30}
	// := 要素の値が必須
	// [...] 要素数を自動で判定するため、要素数の指定が不要
	a3 := [...]int{10, 20}
	fmt.Println(a1, a2, a3)
	fmt.Printf("%v %v %v\n", a1, a2, a3)
	// len: 配列の要素数
	// cap: 配列の容量
	fmt.Printf("len(a1), cap(a1): %v, %v\n", len(a1), cap(a1))
	fmt.Printf("len(a2), cap(a2): %v, %v\n", len(a2), cap(a2))
	fmt.Printf("len(a3), cap(a3): %v, %v\n", len(a3), cap(a3))
	// 配列: 要素数を動的に変更不可、要素数の異なる配列は型も異なる
	fmt.Printf("%T %T %T\n", a1, a2, a3)
	fmt.Println("----------")

	// slice: 要素数を動的に変更可能
	// 要素数が空の配列: slice
	var s1 []int
	s2 := []int{}
	fmt.Printf("s1_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s2, len(s2), cap(s2))
	// カーリーブラケットなし: nil
	fmt.Println("s1 is nil?: ", s1 == nil)
	// カーリーブラケットあり: not nil
	fmt.Println("s2 is nil?: ", s2 == nil)
	fmt.Println("----------")

	// sliceに要素を追加
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	fmt.Println("----------")

	// sliceにsliceを追加
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Printf("s1_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	fmt.Println("----------")

	// make: capacity確保
	s4 := make([]int, 0, 2)
	fmt.Printf("s4_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s4, len(s4), cap(s4))
	// capacityを超えた追加も可能
	s4 = append(s4, 1, 2, 3, 4)
	fmt.Printf("s4_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s4, len(s4), cap(s4))
	fmt.Println("----------")

	// [0, 0, 0, 0]
	s5 := make([]int, 4, 6)
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Println("----------")

	// NOTE: メモリを切り取ると、メモリを共有する
	// s5: [0, 0, 0, 0]
	// s6: [   0, 0   ]
	// s5[1]とs5[2]がs6[0]s6[1]
	s6 := s5[1:3]
	s6[1] = 10
	// s5[2] = 10になる
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s6, len(s6), cap(s6))
	fmt.Println("----------")

	s6 = append(s6, 2)
	// s5[3] = 2になる
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s6, len(s6), cap(s6))
	fmt.Println("----------")

	// NOTE: メモリをコピーすると、メモリを共有しない
	// capacity省略: lengthに基づいて自動計算
	sc6 := make([]int, len(s5[1:3]))
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6 dst copy before: %v %v %v\n", sc6, len(sc6), cap(sc6))

	copy(sc6, s5[1:3])
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6 dst copy after: %v %v %v\n", sc6, len(sc6), cap(sc6))

	// copy後に値を書き換え
	sc6[1] = 10
	// copy元（s5）の値は変更されない
	fmt.Printf("s5 source of copy: %v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6 dst rewrite after: %v %v %v\n", sc6, len(sc6), cap(sc6))
	fmt.Println("----------")

	s5 = make([]int, 4, 6)
	// [1:3:3]指定値(3)-1のインデックスまでメモリ共有
	fs6 := s5[1:3:3]
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", fs6, len(fs6), cap(fs6))
	fmt.Println("----------")

	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	// s5[1]: 6, s5[2]: 7,
	// s5[3]: 0 (メモリ共有はインデックス2までのため、appendは共有されない)
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", fs6, len(fs6), cap(fs6))
	fmt.Println("----------")

	s5[3] = 9
	fmt.Printf("s5_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", s5, len(s5), cap(s5))
	fmt.Printf("fs6_type: %[1]T, value: %[1]v, len: %v, cap: %v\n", fs6, len(fs6), cap(fs6))
	fmt.Println("----------")

	// map(key: string, value: int)
	var m1 map[string]int
	// := 空の{}で明示的に指定
	m2 := map[string]int{}
	// m1 == nil
	fmt.Printf("m1: %v, %v\n", m1, m1 == nil)
	// m2 != nil
	fmt.Printf("m2: %v, %v\n", m2, m2 == nil)
	fmt.Println("----------")

	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	// len: 3, m2["A"] = 10
	fmt.Printf("m2: %v, len: %v, m2['A']: %v\n", m2, len(m2), m2["A"])
	delete(m2, "A")
	// len: 2, m2["A"] = 0
	fmt.Printf("m2: %v, len: %v, m2['A']: %v\n", m2, len(m2), m2["A"])
	fmt.Println("----------")

	// 存在しない値の0と存在する値の0の区別: valueの第2引数
	// 存在しない: false
	v, ok := m2["A"]
	fmt.Printf("m2['A']_v: %v, ok: %v\n", v, ok)
	// 存在する: true
	v, ok = m2["C"]
	fmt.Printf("m2['C']_v: %v, ok: %v\n", v, ok)
	fmt.Println("----------")

	// mapの取り出し(ループ処理)
	for k, v := range m2 {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
	fmt.Println("----------")
}
