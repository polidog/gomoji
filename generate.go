package main

import (
  "flag"
  "os"
  "encoding/base64"
  gomoji "./lib"
)

func generateCmd() command {
  fs := flag.NewFlagSet("gomoji generate", flag.ExitOnError)

  opts := &generateOpts {
    fontType: 1,
  }

  fs.StringVar(&opts.text, "text", "", "generate text")
  fs.StringVar(&opts.output, "output", "", "Output file name")


  return command{fs, func(args []string) error {
    fs.Parse(args)
    return generate(*opts)
  }}
}

type generateOpts struct {
  text string
  fontType int
  output string
}


func generate(opts generateOpts) error {
  dataStr, err := gomoji.Generate(opts.text, opts.fontType)
  if err != nil {
      return err
  }
  return createFile(dataStr, opts.output)
}

func createFile(dataStr string, output string) error {
  data, _ := base64.StdEncoding.DecodeString(dataStr)
  file, err := os.Create(output)
  if err != nil {
    return err
  }
  defer file.Close()
  file.Write(data)
  return err
}
