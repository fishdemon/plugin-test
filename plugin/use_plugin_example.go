package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

type GoodDoctor interface {
	HealthCheck() error
}

func init()  {
	log.Println("main package init function called")
}

func main() {
	log.Println("main function stared")
	// load module 插件您也可以使用go http.Request从远程下载到本地,在加载做到动态的执行不同的功能
	// 1. open the so file to load the symbols

	plug3, err := plugin.Open("./plugin_db.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("plugin db opened")

	function, err := plug3.Lookup("TestCRUD")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	callFunc, ok := function.(func())
	if !ok {
		panic("Plugin has no 'TestCRUD()' function")
	}

	// 调用函数
	callFunc()

	//plug2, err := plugin.Open("./plugin_govaluate2.so")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//log.Println("plugin govaluate2 opened")
	//
	//valuate, err := plug2.Lookup("Valuate")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//valuateFunc, ok := valuate.(func(int32) bool)
	//if !ok {
	//	panic("Plugin has no 'Valuate(int32)bool' function")
	//}
	//
	//// 调用函数
	//result := valuateFunc(4)
	//fmt.Printf("valuate(%v)\n", result)

	//plug1, err := plugin.Open("./plugin_govaluate3.so")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//log.Println("plugin govaluate3 opened")
	//
	//valuate, err = plug1.Lookup("Valuate")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//valuateFunc, ok = valuate.(func(int32) bool)
	//if !ok {
	//	panic("Plugin has no 'Valuate(int32)bool' function")
	//}
	//
	//// 调用函数
	//result = valuateFunc(4)
	//fmt.Printf("valuate(%v)\n", result)


	//plug, err := plugin.Open("./plugin_doctor.so")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//log.Println("plugin opened")
	//
	//// 2. look up a symbol (an exported function or variable)
	//// in this case, variable Greeter
	//doc, err := plug.Lookup("Doctor")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//// 3. Assert that loaded symbol is of a desired type
	//// in this case interface type GoodDoctor (defined above)
	//doctor, ok := doc.(GoodDoctor)
	//if !ok {
	//	fmt.Println("unexpected type from module symbol")
	//	os.Exit(1)
	//}
	//
	//// 4. use the module
	//if err := doctor.HealthCheck(); err != nil {
	//	log.Println("use plugin doctor failed, ", err)
	//}
}