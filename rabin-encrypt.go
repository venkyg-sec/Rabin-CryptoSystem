// Copyright 2017 Venkatesh Gopal(vgopal3@jhu.edu), All rights reserved
package main

import (
  "fmt"
  "io/ioutil"
  "math/big"
  "os"
)

func main() {

  if len(os.Args) != 3 {
    fmt.Println(" \n Follow command line specification \n ./rabin-encrypt" +
      "<publickey-file-name> <message to be encrypted, in decimal>\n")

  } else {

  file_name := os.Args[1]
  MessageInString := os.Args[2]
  Message := ConvertMessageToBigInt(MessageInString)

  N := ExtractDetailsFromPublicKeyFile(file_name)

  Ciphertext := Encrypt(Message, N)
  fmt.Println("Ciphertext is ", Ciphertext)

  }
}

func Encrypt(Message *big.Int, N *big.Int) (*big.Int) {

  exponentationComponent := big.NewInt(2)
  Ciphertext := squareAndMultiple(Message, exponentationComponent, N)
  return Ciphertext

}

func ExtractDetailsFromPublicKeyFile(file_name string) (*big.Int) {

  // In Rabin's crypto-system, N is the public key
  FileContent, err := ioutil.ReadFile(file_name)
  N := big.NewInt(0)

  if err != nil {
    fmt.Println(" Error readng data from the file")
  } else {

  NinString := string(FileContent)

  boolError := false
  N, boolError = N.SetString(NinString,10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  }
  return N
}

func ConvertMessageToBigInt(MessageInString string) (*big.Int) {

  boolError := false
  Message := big.NewInt(0)

  Message, boolError = Message.SetString(MessageInString, 10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  return Message
}


func squareAndMultiple(a *big.Int, b *big.Int, c *big.Int) (*big.Int) {

  // FormatInt will provide the binary representation of a number
  binExp := fmt.Sprintf("%b", b)
  binExpLength := len(binExp)

  initialValue := big.NewInt(0)
  initialValue = initialValue.Mod(a,c)

  // Hold the initial value in result
  result := big.NewInt(0)
  result = result.Set(initialValue)

  // Using the square and multipy algorithm to perform modular exponentation
  for i := 1; i < binExpLength; i++ {

    // 49 is the ASCII representation of 1 and 48 is the ASCII representation
    // of 0
    interMediateResult := big.NewInt(0)
    interMediateResult = interMediateResult.Mul(result,result)
    result = result.Mod(interMediateResult, c)

    if byte(binExp[i]) == byte(49) {
      interResult := big.NewInt(0)
      interResult = interResult.Mul(result,initialValue)
      result = result.Mod(interResult, c)
    }
  }
  return result

}
