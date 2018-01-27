package main

import "fmt"


type IntNode struct{
  value int
  next *IntNode
}

type IntLinkedList struct{
  first *IntNode
}

func (n IntNode) hasnext() bool{
  if n.next == nil {
    return true
  } else {
    return false
  }
}

func (l IntLinkedList) length() int{
  var counter int
  nodepointer := l.first
  for nodepointer != nil{
    counter += 1
    nodepointer = (*nodepointer).next
  }
  return counter
}

func (l IntLinkedList) sum() int {
  var counter int
  nodepointer := l.first
  for nodepointer != nil{
    curNode := *nodepointer
    counter += curNode.value
    nodepointer = curNode.next
  }
  return counter
}

type SandNode struct{
  value int
  n, s, e, w *SandNode
}

func (n SandNode) dirMap(s string) *SandNode{
  var nodeptr *SandNode
  switch s{
  case "n":
    nodeptr = n.n
  case "s":
    nodeptr = n.s
  case "e":
    nodeptr = n.e
  case "w":
    nodeptr = n.w
  }
  return nodeptr
}

func (n SandNode) openDirs() [4]string {
  var open [4]string
  var dirs = [4]string{"n", "s", "e", "w"}
  for i, s := range dirs {
    if n.dirMap(s) == nil {
      open[i] = s
    }
  }
  return open
}

func addnode(old SandNode, new SandNode, direction string){
  var nodeptr = &new
  switch direction{
  case "n":
    old.n = nodeptr
  case "s":
    old.s = nodeptr
  case "e":
    old.e = nodeptr
  case "w":
    old.w = nodeptr
  }
}



func main(){
  var l = IntLinkedList{first : &IntNode{value:1, next: &IntNode{value: 2, next: nil}}}
  fmt.Println(l.length())
  fmt.Println(l.sum())
  fmt.Println("Hello World")
}
