package main

import (
	"github.com/morpheyesh/libgo/cmd"
	"launchpad.net/gnuflag"

)

type DDStart struct {
	fs *gnuflag.FlagSet
	dry bool
}

func (g *DDStart) Info() *cmd.Info {
	desc := `starts the dd API Server.`
	return &cmd.Info{
		Name:    "start",
		Usage:   `start`,
		Desc:    desc,
		MinArgs: 0,
	}
}

func (c *DDStart) Run(context *cmd.Context, client *cmd.Client) error {

	serverRun(c.dry)
	return nil
}

func (c *DDStart) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("sg-scramjet", gnuflag.ExitOnError)
		c.fs.BoolVar(&c.dry, "config", false, "config: the configuration file to use")
		c.fs.BoolVar(&c.dry, "c", false, "config: the configuration file to use")
		c.fs.BoolVar(&c.dry, "dry", false, "dry-run: does not start the sg-scramjet (for testing purpose)")
		c.fs.BoolVar(&c.dry, "d", false, "dry-run: does not start the sg-scramjet (for testing purpose)")


	}
	return c.fs
}
