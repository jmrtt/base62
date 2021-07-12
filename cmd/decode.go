package cmd

import (
	"base62/encoding"
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decodeCmd)
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode base62 data",
	Long:  `Receives data as a parameter and decodes to a string`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()
			if scanner.Err() != nil {
				fmt.Println("Error: ", scanner.Err())
			}
			data := encoding.Decode([]byte(text))
			fmt.Println(string(data))
			return
		}

		data := encoding.Decode([]byte(args[0]))
		fmt.Println(string(data))
	},
}
