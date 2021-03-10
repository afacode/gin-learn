package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a, _ := gormadapter.NewAdapter("mysql", "root:zxcv1234@tcp(127.0.0.1:3306)/gorbca", true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("./model.conf", a)

	sub := "zhangsan" // the user that wants to access a resource.
	obj := "data1"    // the resource that is going to be accessed.
	act := "read"     // the operation that the user performs on the resource.

	// 增 AddPolicy
	// added, err := e.AddPolicy("zhangsan", "data2", "read")
	// fmt.Println(added)
	// fmt.Println(err)
	// 查
	filteredPolicy := e.GetFilteredPolicy(0, "zhangsan")
	fmt.Println(filteredPolicy)

	// 跟新
	// updated, err := e.UpdatePolicy([]string{"zhangsan", "data1", "read"}, []string{"zhangsan", "data3", "write"})
	// fmt.Println(updated)
	// fmt.Println(err)

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		fmt.Println(err)
		// handle err
	}

	if ok == true {
		fmt.Println("通过")
		// permit alice to read data1
	} else {
		fmt.Println("未通过")
		// deny the request, show an error
	}

}
