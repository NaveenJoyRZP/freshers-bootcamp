package main

import "fmt"

type tree struct {
  value string
  leftBranch *tree
  rightBranch *tree
}

func generateTree() tree {
  b1 := tree{"b", nil, nil}
  b2 := tree{"c", nil, nil}
  b3 := tree{"_", &b1, &b2}
  b4 := tree{"a", nil, nil}
  finalTree := tree{"+", &b4, &b3}
  return finalTree
}

func (btree *tree) printTreeInOrder() {
  if btree == nil {
    return
  }
  btree.leftBranch.printTreeInOrder()
  fmt.Print(btree.value + " ")
  btree.rightBranch.printTreeInOrder()
}

func (btree *tree) printTreePreOrder() {
  if btree == nil {
    return
  }
  fmt.Print(btree.value + " ")
  btree.leftBranch.printTreePreOrder()
  btree.rightBranch.printTreePreOrder()
}

func (btree *tree) printTreePostOrder() {
  if btree == nil {
    return
  }
  btree.leftBranch.printTreePostOrder()
  btree.rightBranch.printTreePostOrder()
  fmt.Print(btree.value + " ")
}

func main() {
  branchedTree := generateTree()

  fmt.Println("In Order traversal")
  branchedTree.printTreeInOrder()
  fmt.Println("\nPre Order Traversal")
  branchedTree.printTreePreOrder()
  fmt.Println("\nPost Order Traversal")
  branchedTree.printTreePostOrder()
  fmt.Println("")
}
