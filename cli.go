package main

import(
    "log"
    "os"

    "github.com/mitchellh/cli"
)

type Foo struct{}

func (f *Foo) Help() string {
    return "app foo"
    }

func (f *Foo) Run(args []string) int {
    log.Println("Foo!")
    return 0
    }

func (f *Foo) Synopsis() string {
    return "Print \"Foo\""
    }

type Bar struct{}

func (f *Bar) Help() string {
    return "app foo"
    }

func (f *Bar) Run(args []string) int {
    log.Println("Bar!")
    return 0
    }

func (f *Bar) Synopsis() string {
    return "Print \"Bar\""
    }



func main() {
    c := cli.NewCLI("whroom", "1.0.0")
    c.Args = os.Args[1:]
    c.Commands = map[string]cli.CommandFactory{
        "foo": func() (cli.Command, error) {
            return &Foo{}, nil
            },
        "bar": func() (cli.Command, error) {
            return &Bar{}, nil
            },
        }

        exitStatus, err := c.Run()
        if err != nil {
            log.Println(err)
            }

            os.Exit(exitStatus)
}
