package main

// import "fmt"

type SandNode struct{
  x, y int
  value int
  n, s, e, w *SandNode
}

type SandPile struct{
  root *SandNode
}

func addPile(p1 SandPile, p2 SandPile) SandPile {
  var newroot = addNode(*(p1.root), *(p2.root))
  return SandPile{root: &newroot}
}


func addNode(n1 SandNode, n2 SandNode) SandNode {
  var newnorth = addNode(*n1.n, *n2.n)
  var newsouth = addNode(*n1.s, *n2.s)
  var neweast = addNode(*n1.e, *n2.e)
  var newwest = addNode(*n1.w, *n2.w)
  return SandNode{x: n1.x, y: n1.y, value: n1.value + n2.value, n: &newnorth,s: &newsouth,e: &neweast,w: &newwest }
}

func (n SandNode) isValid() bool {
  if (n.value >= 0) {
    return true
  } else {
    return false
  }
}

func (n SandNode) isStable() bool{
  if (n.value < 4) {
    return true
  } else {
    return false
  }
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
