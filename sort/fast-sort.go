/*
快速排序的练习，通过这次练习我又发现自己的盲区了，
「slice」这个数据结构，在使用make定义的时候，
如果第二个参数为0，也就是说「slice」里面没有任何元素，它的长度是0（先不考虑容量cap）
就是在这样的情况下，我往里面添加元素，只能通过append，而我在写这个代码的时候，我需要拷贝
「slice」到另一个「slice」里，那么如果我得到拷贝元素的「slice」的len是0，
那么我是拷贝不进去的，因为拷贝是：同一位置同一元素进行拷贝。
那么现在这种情况下，我没有位置，更别提元素了，So，无法拷贝。。
*/
package main

import "fmt"

//FastSort is a fast way to sort slice.
func FastSort(a []int) []int {
	if len(a) <= 1 {
		/*
			关于为什么要在这里返回一个拷贝值,
			首先，这个返回值是作为上层递归的结果值的，
			也就是说，在这个if的地方假如我返回的是a，
			那么我之后在上层递归里面如果有修改这个返回值的操作，
			那么，我这个sort就出现bug了（Panic）
		*/
		temp := make([]int, len(a))
		copy(temp, a)
		return temp
	}
	pivot := a[0]
	low := make([]int, 0, len(a))
	high := make([]int, 0, len(a))
	for index := 1; index < len(a); index++ {
		if a[index] < pivot {
			low = append(low, a[index])
		} else {
			high = append(high, a[index])
		}
	}
	low, high = FastSort(low), FastSort(high)
	//low[0] = 1
	return append(append(low, pivot), high...)
}

func test(t []int) []int {
	// temp := make([]int, len(t))
	// copy(temp, t)
	// temp[0] = 1
	t[0] = 1
	return t
}

func main() {
	//testing data.
	var b = []int{27, 38, 12, 39, 27, 16}
	//fmt.Println(FastSort(b))
	fmt.Println(b)
	newb := FastSort(b)
	fmt.Println(b)
	fmt.Println(newb)
	// fmt.Println(b)
	// newb := test(b)
	// fmt.Println(b)
	// fmt.Println(newb)
	// newb[1] = 2
	// fmt.Println(newb)
	// fmt.Println(b)
}
