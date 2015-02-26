package ascii_art

import (
  "fmt"

  "github.com/hectorj2f/search_cli/version"
  )

func GiantSwarmfy() {
  fmt.Printf("\x1b[31;1mGiant swarm user search-engine, version: %s\x1b[0m \n", version.Version)
  fmt.Println("   ////*       *\\\\\\\\ ")
  fmt.Println(" ,//  ///////////  \\\\, ")
  fmt.Println("./   /////////////    \\.")
  fmt.Println("   /////////////////.    ")
  fmt.Println(" ./.    ///////    ./.   ")
  fmt.Println(" //   >  )///( --  //    ")
  fmt.Println("  ,/*    )///(    ,/*    ")
  fmt.Println("    ,/////////////,      ")
  fmt.Println("       //////////        ")
  fmt.Println("       ((.   .))         ")
}
