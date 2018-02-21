package main

import "fmt"
import "practice/hello/coordinateSet"

type SandNode struct{
  x, y int
  value int
  n, s, e, w *SandNode
}

type SandPile struct{
  root *SandNode
}

func newRootNode(value int) SandNode {
  return SandNode{x: 0, y: 0, value: value, n: nil, s: nil, e: nil, w:nil}
}

func addPiles(p1 SandPile, p2 SandPile) SandPile {
  var newroot = addNodes(p1.root, p2.root)
  return SandPile{root: newroot}
}


func addNodes(n1 *SandNode, n2 *SandNode) *SandNode {
  if (n1 == nil) || (n2 == nil) {
    return nil
  }
  var newnorth = addNodes(n1.n, n2.n)
  var newsouth = addNodes(n1.s, n2.s)
  var neweast = addNodes(n1.e, n2.e)
  var newwest = addNodes(n1.w, n2.w)
  var n1Node = *n1
  var n2Node = *n2
  var newNode = SandNode{x: n1Node.x, y: n1Node.y, value: n1Node.value + n2Node.value, n: newnorth, s: newsouth, e: neweast, w: newwest }
  return &newNode
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

func (old SandNode) addNode( newNodeVal int, direction string){
  var newx int
  var newy int
  switch direction{
    case "n":
      newx = old.x
      newy = old.y + 1
    case "s":
      newx = old.x
      newy = old.y - 1
    case "e":
      newx = old.x + 1
      newy = old.y
    case "w":
      newx = old.x - 1
      newy = old.y
  }
  var oldnodeptr = &old
  switch direction{
    case "n":
      var newnode = SandNode{x: newx, y: newy, value: newNodeVal, n: nil, s: oldnodeptr, e: nil, w: nil}
      var nodeptr = &newnode
      old.n = nodeptr
    case "s":
      var newnode = SandNode{x: newx, y: newy, value: newNodeVal, n: oldnodeptr, s: nil, e: nil, w: nil}
      var nodeptr = &newnode
      old.s = nodeptr
    case "e":
      var newnode = SandNode{x: newx, y: newy, value: newNodeVal, n: nil, s: nil, e: nil, w: oldnodeptr}
      var nodeptr = &newnode
      old.e = nodeptr
    case "w":
      var newnode = SandNode{x: newx, y: newy, value: newNodeVal, n: nil, s: nil, e: oldnodeptr, w: nil}
      var nodeptr = &newnode
      old.w = nodeptr
  }
}

func main(){
  var root1 = newRootNode(3)
  root1.addNode(3, "n")
  root1.addNode(5, "w")
  var root2 = newRootNode(2)
  root2.addNode(6, "n")
  root2.addNode(3, "e")
  var pile1 = SandPile{ root: &root1}
  var pile2 = SandPile{ root: &root2}
  var pile3 = addPiles(pile1, pile2)
  fmt.Println(pile3)

}
