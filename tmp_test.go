package cast

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndirect(t *testing.T) {
	s := "hello"
	n := 1
	c := 1.1
	type enh struct {
		name string
	}
	enhh := enh{name: "555"}
	fmt.Println(indirect(s))
	fmt.Println(indirect(n))
	fmt.Println(indirect(c))
	fmt.Println(indirect(enhh))	

	// 创建一个多重指针
	var x int = 42
	ptr1 := &x       // 指向 x 的指针
	ptr2 := &ptr1    // 指向 ptr1 的指针
	ptr3 := &ptr2    // 指向 ptr2 的指针
	ptr4 := &ptr3    // 指向 ptr3 的指针

	// 使用反射获取多重指针的反射值
	v := reflect.ValueOf(ptr4)

	// 判断反射值是否是指针类型并且不是 nil
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		// 解引用指针并更新反射值
		v = v.Elem()
		fmt.Printf("%T\n", v)
	}
	fmt.Printf("%T\n", v.Interface())
	// 打印解引用后的反射值
	fmt.Println("Dereferenced value:", v.Interface())
}