package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute(version string) {
	rootCmd := &cobra.Command{
		Use:   "kraw",
		Short: "A tool for bypass firewall",
		Long: `Kraw is a proxy tool.
        It will help you bypass firewall.`,
		Version: version,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
