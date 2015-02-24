package main

import (
  "fmt"
  "log"

  "github.com/hectorj2f/search_cli/resources"
  "github.com/hectorj2f/search_networking/networking"
  )

const (
  cmdListName = "list"
  )

var (
  cmdList = &Command{
    Name:    "list",
    Summary: "List all the users",
    Run:     runList,
  }
)

func init() {
  commands = append(commands, cmdList)
}

func runList(args []string) (exit int) {
  if len(args) > 0 {
    printCommandUsageByName(cmdListName)
    return 1
  }

  var use_tls = false
  if globalFlags.UseTls {
    use_tls = true
  }
  server_addr := globalFlags.ServerAddr
  if globalFlags.ServerAddr == "" {
    server_addr = resources.SERVER_ADDR
  }
  search_query := map[string]interface{}{}
  result, err := networking.Query(search_query, server_addr, use_tls, globalFlags.ServerPort)
  if err != nil {
    log.Print(err)
    return 1
  }

  for _, user := range result {
    fmt.Printf("ID: %d CREATED: %s USERNAME: %s ROLE: %s ORGANIZATION: %s\n",
                 user["id"].(int64),
                 user["created"].(string),
                 user["username"].(string),
                 user["role"].(string),
                 user["organization"].(string))
  }
  if len(result) == 0 {
    fmt.Println("=> No results found with this criteria.")
  }
  return
}
