package main

import(
	"fmt"
	"./tree"
)

func auxrule() bool{
	fmt.Println("root node, (aux rule function)")
	return true
}

func main(){
	var testtree tree.Node
	testtree.Insert(auxrule)
	testtree.Insert(auxrule)
}