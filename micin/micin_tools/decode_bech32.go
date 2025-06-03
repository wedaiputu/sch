package main

import (
  "fmt"
  "github.com/btcsuite/btcutil/bech32"
)

func main() {
  hrp, data, err := bech32.Decode("cosmos1f20rsfxujw8v8xmzcl5a0608ahvdwt0l2st8at")
  if err != nil {
    panic(err)
  }
  decoded, err := bech32.ConvertBits(data, 5, 8, false)
  if err != nil {
    panic(err)
  }
  fmt.Printf("HRP: %s\nBytes (hex): %x\n", hrp, decoded)
}

