package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func init() {
	fmt.Println("init plugin govaluate3...")
}

func Valuate(a int32) bool {
	fmt.Println("call plugin method...")
	expression, _ := govaluate.NewEvaluableExpression("a > 0")

	parameters := make(map[string]interface{}, 8)
	parameters["a"] = a

	result, _ := expression.Evaluate(parameters)
	fmt.Println("finish plugin method...")
	return result.(bool)
}

