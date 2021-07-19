package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jmrtt/base62/encoding"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encodeCmd)
}

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode base62 data",
	Long:  `Receives string as a parameter and encodes to base62`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()
			if scanner.Err() != nil {
				fmt.Println("Error: ", scanner.Err())
			}
			data := encoding.Encode([]byte(text))
			fmt.Println(string(data))
			return
		}

		data := encoding.Encode([]byte(strings.Join(args, " ")))
		fmt.Println(string(data))
	},
}
