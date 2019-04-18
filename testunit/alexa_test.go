package main

import (
  "os"
  "io/ioutil"
//  "bytes"
  "fmt"
  "testing"
  "net"
//  "net/rpc"
  "net/rpc/jsonrpc"
  "time"
  
  "github.com/webability-go/alexa"
)

func TestAlexa(t *testing.T) {
  go alexa.Start()
  
  time.Sleep(time.Second)
  
  jsonFile, _ := os.Open("launch.json")
  jsondata, _ := ioutil.ReadAll(jsonFile)
  jsonFile.Close()
  
  client, err := net.Dial("tcp", "127.0.0.1:1443")
  if err != nil {
    fmt.Println("Error Dialing:", err)
    return
  }
  
  c := jsonrpc.NewClient(client)

  var reply interface{}

  err = c.Call("Launch", string(jsondata), &reply)
  if err != nil {
    fmt.Println("ERROR 1", err)
    return
  }
  fmt.Println(reply)
}



