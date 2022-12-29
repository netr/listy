package cmd

import (
	"log"

	"github.com/netr/listy/iom"
	"github.com/spf13/cobra"
)

// concatCmd represents the concat command
var concatCmd = &cobra.Command{
	Use:   "concat",
	Short: "Concatenate files in a directory into a single file. Output default `{dir}/all.txt`",
	Long:  `Concatenate files in a directory into a single file`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := validateFlag(cmd, "dir")
		out := getFlag(cmd, "out")
		if out == "" {
			out = sanitizeFilename(dir + "/all.txt")
		}

		files, err := iom.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		var result []string
		for _, file := range files {
			file = sanitizeFilename(dir + "/" + file)
			log.Printf("Concatenating %s to %s ... ", file, out)
			lines, err := iom.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}

			result = append(result, lines...)
		}

		err = iom.WriteFile(out, result)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Done! %d lines written to %s", len(result), out)
	},
}

func init() {
	rootCmd.AddCommand(concatCmd)
	concatCmd.Flags().StringP("dir", "d", "", "Directory to concat into a single file")
	concatCmd.Flags().StringP("out", "o", "", "Output file")
}
