package create

import (
	"flag"
	"fmt"

	"github.com/mitchellh/cli"

	"github.com/hiromaily/go-bitcoin/pkg/wallet"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/key"
)

// key subcommand
type KeyCommand struct {
	name     string
	synopsis string
	ui       cli.Ui
	wallet   wallet.Keygener
}

func (c *KeyCommand) Synopsis() string {
	return c.synopsis
}

func (c *KeyCommand) Help() string {
	return `Usage: keygen key create key
`
}

func (c *KeyCommand) Run(args []string) int {
	c.ui.Output(c.Synopsis())

	flags := flag.NewFlagSet(c.name, flag.ContinueOnError)
	if err := flags.Parse(args); err != nil {
		return 1
	}

	// create one key for debug use
	wif, pubAddress, err := key.GenerateKey(c.wallet.GetBTC().GetChainConf())
	if err != nil {
		c.ui.Error(fmt.Sprintf("fail to call key.GenerateKey() %+v", err))
	}
	c.ui.Info(fmt.Sprintf("[WIF] %s - [Pub Address] %s", wif.String(), pubAddress))

	return 0
}
