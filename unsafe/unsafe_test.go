package unsafe

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceHeader(t *testing.T) {
	old := make([]int, 1, 10)
	old[0] = 10

	new := append(old, 20)
	t.Log(new, old)

	// 用unsafe获取old的引用数组
	oldHeader := (*reflect.SliceHeader)(unsafe.Pointer(&old))
	oldArr := (*[10]int)(unsafe.Pointer(oldHeader.Data))

	// 下面就会打印出&[10 20 0 0 0 0 0 0 0 0],可以看到底层数组已经变了,只是old看不到第二个元素
	t.Log(oldArr)

	// 我们手动把old里的Len改成2再看看,发现再打印old的时候已经能看到第二个元素20了
	oldHeader.Len = 2
	t.Log(old)
}
