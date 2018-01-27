package "intlinkedist"

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
