package controller

import (
	"fmt"

	"github.com/yichouchou/ginPlus/bind"
)

//todo 在controller结构体上添加请求类型，请求内容类型，字符类型等等，鉴权类型等
type Hello struct {
	Name string
	Age  int
}

// [name string, password string, age int]
// @GET /block1
func (s *Hello) Hi1(parm1 string, parm2 string, parm3 int, hiValue bind.ReqTest, hi *bind.ReqTest) (commentHi1 string, errHi1 error) {
	fmt.Println(parm1 + parm2)
	fmt.Println(parm3)
	fmt.Println(hi)
	fmt.Println(hiValue)
	return "ni hao", nil
}

// @POST /block2
func (s *Hello) Hi2(hiValue bind.ReqTest, hi *bind.ReqTest) (commentHi2 bind.ReqTest, errHi2 error) {
	fmt.Println(hi)
	fmt.Println(hiValue)
	return bind.ReqTest{}, nil
}

// @GET /block3
func (s *Hello) Hi3(name, password string, age, year int) (commentHi3 int, errHi3 error) {
	fmt.Println("---get请求，无参数注解，多基本数据类型已经调通")
	return 10, nil
}

// {List []bind.ReqTest}
// @POST /block4
func (s *Hello) Hi4(reqList []bind.ReqTest) (index int, errHi4 error) {
	fmt.Println("这是post请求接受结构体数组")
	fmt.Println("---post请求，有参数注解，结构体值数组已经调通")

	for i := range reqList {
		fmt.Println(reqList[i])
	}
	return 4, nil
}

// [reqList *bind.ReqTest]
// @GET /block5
func (s *Hello) Hi5(reqList *bind.ReqTest) (index int, errHi5 error) {
	fmt.Println(reqList)
	fmt.Println("这是get请求接受结构体指针")
	fmt.Println("---这是get请求接受结构体指针已经调通")
	return 5, nil
}

// {List []*bind.ReqTest}
// @POST /block6
func (s *Hello) Hi6(reqList []*bind.ReqTest) (index int, errHi4 error) {
	fmt.Println("这是post请求接受指针结构体数组")
	fmt.Println("---post请求，有参数注解，结构体指针数组已经调通")
	for i := range reqList {
		fmt.Println(reqList[i])
	}
	return 6, nil
}

// [reqList bind.ReqTest]
// @GET /block7
func (s *Hello) Hi7(reqList bind.ReqTest) (index int, errHi5 error) {
	fmt.Println(reqList)
	fmt.Println("这是get请求接受结构体指针")
	fmt.Println("---这是get请求接受结构体指针已经调通")
	return 5, nil
}

//
//func init() {
//	annotation.SetVersion(1625327764)
//	annotation.AddGenOne("Hello.Hello", "/block", []string{"post"})
//	_= map[string]map[string]string{}
//}
//
////annotation.AddGenOne("Hello.Hello", "/block", []string{"post"})
//
//// []string{"post"} 这样的描述太过于简洁，我希望能够把入参的名称和类型都注册上去，把返回值类型也注册上去-这个不是绝对的
////													入参1类型：string  入参名字:name
////													入参2类型：string  入参名字:password
////													入参3类型：string  入参名字:age
////												          * * *
////													以什么协议传输：josn -默认
////大致的类型： _= map[string]map[string]string{}   post:
//
////然后在请求进入的时候，根据这个map 获取到相关的参数，

// Example ...
type Example struct {
}

