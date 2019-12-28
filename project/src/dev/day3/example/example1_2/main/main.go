// original codes

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func urlProcess(url string) string {

// 	result := strings.HasPrefix(url, "http://")
// 	if !result {
// 		url = fmt.Sprintf("http://%s", url)
// 	}

// 	return url
// }

// strings.TrimSpace(" sksk ") =>"sksk"
// strings.Trim("abbacba", "ab") =>"c"
// strings.TrimLeft("a","b")
// strings.TrimRight("b","c")
// "heheheworld", "he", "wo", 0

// strings.Fields("abc cde edk") ["abc", "cde", "edk]

// strings.Split("abc,cde,edk", ",") ["abc", "cde", "edk]
// strings.Join(["abc", "cde", "edk], ",") "abc,cde,edk"

// strings.Replace("str", 3) "strstrstr"
// func pathProcess(path string) string {
// 	result := strings.HasSuffix(path, "/")
// 	if !result {
// 		path = fmt.Sprintf("%s/", path)
// 	}

// 	return path
// }

// func main() {
// 	var (
// 		url  string
// 		path string
// 	)

// 	fmt.Scanf("%s%s", &url, &path)
// 	url = urlProcess(url)
// 	path = pathProcess(path)

// 	fmt.Println(url)
// 	fmt.Println(path)
// }

package main

import (
	"fmt"
	"strings"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}

	return url

}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func main() {
	var (
		url  string
		path string
	)

	// fmt.Scanf("%s%s", &url, &path)   //  在vs里面怎么调试

	fmt.Println(url)
	fmt.Println(path)

	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Println(url)
	fmt.Println(path)

	fmt.Println(strings.Fields("abc cde edk"))
	fmt.Printf("%T\n", strings.Fields("abc cde edk"))

	fmt.Println(strings.Trim("abbacba", "ab"))
	fmt.Println(strings.TrimLeft("abbacba", "ab"))
	fmt.Println(strings.Split("abc,cde,edk", ","))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))

	var arr1 = []int{1, 2, 3}
	var arr2 = new([5]int)
	slice1 := make([]int, 2)

	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	fmt.Printf("%T\n", slice1)

	// fmt.Printf("%p\n", arr1)
	// fmt.Printf("%p\n", arr2)
	// fmt.Printf("%p\n", slice1)

	// var a int = 9
	// fmt.Printf("%q\n", a)
}
