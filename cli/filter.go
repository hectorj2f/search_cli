package main

import (
  "fmt"

  "github.com/hectorj2f/search_cli/resources"
  "github.com/hectorj2f/search_networking/networking"
  )

const (
  cmdFilterName = "filter"
  )

var (
  cmdFilter = &Command{
    Name:    "filter",
    Summary: "Filter users by a field",
    Usage:   "--id [USERID] --role [ROLE] --organization [ORGANIZATION] --username [USERNAME]",
    Run:     runFilter,
  }
  flagRole  string
  flagOrganization  string
  flagUsername string
  flagUserId  string
)

func init() {
  commands = append(commands, cmdFilter)
  cmdFilter.Flags.StringVar(&flagUserId, "id", "", "id of the user to search for")
  cmdFilter.Flags.StringVar(&flagRole, "role", "", "role of the user to search for")
  cmdFilter.Flags.StringVar(&flagOrganization, "organization", "", "name of the organization to search for")
  cmdFilter.Flags.StringVar(&flagUsername, "username", "", "name of the user to search for")
}

func runFilter(args []string) (exit int) {
  if len(args) > 0 {
    printCommandUsageByName(cmdFilterName)
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

  search_query := processArgs()
  result, err := networking.Query(search_query, server_addr, use_tls, globalFlags.ServerPort)
  if err != nil {
    panic(err)
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

func processArgs() (map[string]interface{}){
  query := make(map[string]interface{})
  if flagOrganization != "" {
    query["organization"] = flagOrganization
  }
  if flagRole != "" {
    query["role"] = flagRole
  }
  if flagUserId != "" {
    query["id"] = flagUserId
  }
  if flagUsername != "" {
    query["username"] = flagUsername
  }
  return query
}
