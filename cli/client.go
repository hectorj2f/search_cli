package main

import (
  "fmt"
  "flag"
  "os"
  "strconv"
  "text/tabwriter"

  "github.com/hectorj2f/search_cli/cli/ascii_art"
  "github.com/hectorj2f/search_cli/resources"
)

const (
  cliName        = "swarmsearch"
  cliDescription = "swarmsearch, the application to search swarm users"
)

var (
  commands      []*Command // Commands should register themselves by appending
  globalFlagset = flag.NewFlagSet(cliName, flag.ExitOnError)
  globalFlags   = struct {
    Debug              bool
    Help               bool
    UseTls             bool
    ServerAddr         string
    ServerPort         int
  }{}
  out           *tabwriter.Writer
)

type Command struct {
  Name        string
  Description string
  Flags       flag.FlagSet
  Summary     string
  Usage       string

  Run func(args []string) int // Run a command with the given arguments, return exit status
}

func init() {
  globalFlagset.BoolVar(&globalFlags.Help, "help", false, "Print usage information and exit")
  globalFlagset.BoolVar(&globalFlags.Debug, "debug", false, "Print out more debug information to stderr")
  globalFlagset.BoolVar(&globalFlags.UseTls, "use-tls", false, "Use TLS in the communication layer")
  globalFlagset.StringVar(&globalFlags.ServerAddr, "server-addr", "", "Set the swarm server to be connected")

  if os.Getenv(resources.PORT_FLAG) != "" {
    port, _ := strconv.Atoi(os.Getenv(resources.PORT_FLAG))
    globalFlagset.IntVar(&globalFlags.ServerPort, "server-port", port, "Set the port server to be connected")
  } else {
    globalFlagset.IntVar(&globalFlags.ServerPort, "server-port", resources.SERVER_PORT, "Set the port server to be connected")
  }
}

func init() {
  out = new(tabwriter.Writer)
  out.Init(os.Stdout, 0, 4, 1, '\t', 0)
}

func main(){
  globalFlagset.Parse(os.Args[1:])
  args := globalFlagset.Args()
  if len(args) < 1 || globalFlags.Help {
    args = []string{"help"}
  }
  var cmd *Command
  for _, command := range commands {
    if command.Name == args[0] {
      cmd = command
      switch cmd.Name {
        case "filter":
          if len(os.Args[2:]) > 2 {
            if err := cmd.Flags.Parse(os.Args[3:]); err != nil {
              fmt.Fprintf(os.Stderr, "%v\n", err)
              os.Exit(2)
            }
          }
        case "list":
          if err := cmd.Flags.Parse(args[1:]); err != nil {
            fmt.Fprintf(os.Stderr, "%v\n", err)
            os.Exit(2)
          }
      }
      break
    }
  }

  if cmd == nil {
    fmt.Fprintf(os.Stderr, "%v: unknown subcommand: %q\n", cliName, args[0])
    fmt.Fprintf(os.Stderr, "Run '%v help' for usage.\n", cliName)
    os.Exit(2)
  }
  // Print the GiantSwarm logo
  ascii_art.GiantSwarmfy()

  os.Exit(cmd.Run(cmd.Flags.Args()))
}

func getAllFlags() (flags []*flag.Flag) {
  return getFlags(globalFlagset)
}

func getFlags(flagset *flag.FlagSet) (flags []*flag.Flag) {
  flags = make([]*flag.Flag, 0)
  flagset.VisitAll(func(f *flag.Flag) {
    flags = append(flags, f)
  })
  return
}
