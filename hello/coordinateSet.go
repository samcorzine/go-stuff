package main

type Coordinate struct {
  x,y int
}

type CoordinateSet struct {
  coordinateMap map[Coordinate]bool
}

func (s CoordinateSet) put(c Coordinate) {
  s.coordinateMap[c] = true
}

func (s CoordinateSet) get(c Coordinate) bool {
  return s.coordinateMap[c]
}
