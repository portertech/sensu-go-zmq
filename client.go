package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "os/exec"
)

import zmq "github.com/alecthomas/gozmq"

type SensuConfig struct {
  Servers []SensuServers
}

type SensuServers struct {
  Host string
  Port int
}

func main() {
  file, err := ioutil.ReadFile("./config.json")
  if err != nil {
    fmt.Printf("File error: %v\n", err)
    os.Exit(1)
  }
  fmt.Printf("%s\n", string(file))

  config := SensuConfig{}
  json.Unmarshal(file, &config)
  fmt.Printf("Results: %v\n", config)

  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REQ)

  for _, server := range config.Servers {
    uri := fmt.Sprintf("tcp://%s:%d", server.Host, server.Port)
    fmt.Printf("Connecting to server: %v\n", uri)
    socket.Connect(uri)
  }

  cmd := exec.Command("echo", "-n", "foo")
  output, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Printf("Command error: %v\n", err)
  }
  cmd.Wait()
  fmt.Printf("Command output: %s\n", output)
}
