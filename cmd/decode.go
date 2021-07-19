package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jmrtt/base62/encoding"
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
			data, err := encoding.Decode([]byte(text))
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
			fmt.Println(string(data))
			return
		}

		data, err := encoding.Decode([]byte(args[0]))
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		fmt.Println(string(data))
	},
}
