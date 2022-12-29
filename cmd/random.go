package cmd

import (
	"log"

	"github.com/netr/listy/iom"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Randomize lines of file(s)",
	Long:  `Randomize lines of file(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := getFlag(cmd, "dir")
		if dir != "" {
			randomizeDir(cmd, dir)
			return
		}

		file := validateFlag(cmd, "file")
		out := getFlag(cmd, "out", iom.AppendSuffixToFilename(file, "-rnd"))

		log.Printf("Randomizing %s to %s", file, out)
		err := iom.ShuffleFile(file, out)
		if err != nil {
			log.Fatal(err)
		}

		count, err := iom.CountFileLines(file)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Randomized %d lines", count)
	},
}

func randomizeDir(cmd *cobra.Command, dir string) {
	log.Printf("Randomizing directory %s\n\n", dir)

	if dir != "" {
		files, err := iom.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			file = sanitizeFilename(dir + "/" + file)
			out := iom.AppendSuffixToFilename(file, "-rnd")

			log.Printf("Randomizing %s to %s", file, out)
			err := iom.ShuffleFile(file, out)
			if err != nil {
				log.Fatal(err)
			}

			count, err := iom.CountFileLines(file)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Randomized %d lines\n\n", count)
		}
		return
	}
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().StringP("file", "f", "", "File to randomize")
	randomCmd.Flags().StringP("dir", "d", "", "Directory to randomize")
	randomCmd.Flags().StringP("out", "o", "", "Output file")
}
