package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  fmt.Println("eludedGuards has a value of", eludedGuards) 
  
  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
} else {
    fmt.Println("Plan a better disguise next time?")
}
 
  fmt.Println("isHeistOn is currently:", isHeistOn)

  openedVault := rand.Intn(100)

  if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
} else if isHeistOn || openedVault <= 70 {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
}

  leftSafely := rand.Intn(5)

  if isHeistOn == true {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("Looks like you tripped an alarm… run?")
      case 1:
        isHeistOn = false
        fmt.Println("Turns out vault doors don’t open from the inside…")
      case 2:
        isHeistOn = false
        fmt.Println("The alarm went off…")
      case 3:
        isHeistOn = false
        fmt.Println("Cops are coming…")
      default:
        fmt.Println("Start the getaway car!")
}

  if isHeistOn == true {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("$", amtStolen, "not bad!")
}
}
  fmt.Println("isHeistOn is currently:", isHeistOn)
}