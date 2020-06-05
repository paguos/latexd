package main

import (
  "fmt"
  "os"
  
  log "github.com/sirupsen/logrus"
  "github.com/urfave/cli/v2"
)

func main() {
  version := "v0.1.0-rc1"
  app := &cli.App{
    Name: "latexd",
    Usage: "dockerized tooling for latex",
    Flags: []cli.Flag{
      &cli.StringFlag{
        Name:    "image",
        Value: "paguos/latexd:" + version,
        Aliases: []string{"i"},
        Usage:   "Docker IMAGE to be used",
      },
    },
    Commands: []*cli.Command{
      {
        Name:        "run",
        Aliases:     []string{"r"},
        Usage:       "Generate latex document",
        Action:  func(c *cli.Context) error {
          if c.NArg() > 0 {
            filePath := c.Args().First()
            imageName := c.String("image")
            runContainer(pdfCommand(filePath), imageName)
          } else {
            log.Fatal("FILE PATH was not provided")
          }
          return nil
        },
      },
      {
        Name:        "shell",
        Aliases:     []string{"s"},
        Usage:       "Open a shell inside the latexd container (not implemented)",
        Action:  func(c *cli.Context) error {
          log.Warning("Command not implemented!")
          return nil
        },
      },
      {
        Name:        "version",
        Aliases:     []string{"v"},
        Usage:       "Print version information",
        Action:  func(c *cli.Context) error {
          fmt.Println("Version: ", version)
          return nil
        },
      },
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}