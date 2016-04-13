package main

import (
  "./routes"
)

func main() {
  n  := routes.GetRouter()
  n.Run(":3000")
}
