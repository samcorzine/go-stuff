package main

import (
  "net/http"
  "strings"
)

type Board struct{
  spaces [9][9] int 
}
