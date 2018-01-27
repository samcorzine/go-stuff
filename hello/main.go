package main

import "fmt"

type SandNode struct{
  x, y int
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
}
