package cmd

import (
	"log"

	"github.com/netr/listy/iom"
	"github.com/spf13/cobra"
)

// dedupeCmd represents the dedupe command
var dedupeCmd = &cobra.Command{
	Use:   "dedupe",
	Short: "Dedupe file(s)",
	Long:  `Dedupe file(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := getFlag(cmd, "dir")
		if dir != "" {
			dedupeDir(cmd, dir)
			return
		}

		file := validateFlag(cmd, "file")
		out := getFlag(cmd, "out", iom.AppendSuffixToFilename(file, "-deduped"))

		log.Printf("Deduping %s to %s", file, out)
		n, err := iom.RemoveDuplicatesFile(file, out)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Removed %d duplicates", n)
	},
}

func dedupeDir(cmd *cobra.Command, dir string) {
	log.Printf("Deduping directory %s\n\n", dir)

	if dir != "" {
		files, err := iom.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			file = sanitizeFilename(dir + "/" + file)
			out := iom.AppendSuffixToFilename(file, "-deduped")
			log.Printf("Deduping %s to %s", file, out)
			n, err := iom.RemoveDuplicatesFile(file, out)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Removed %d duplicates\n\n", n)
		}
		return
	}
}

func init() {
	rootCmd.AddCommand(dedupeCmd)
	dedupeCmd.Flags().StringP("file", "f", "", "File to dedupe")
	dedupeCmd.Flags().StringP("dir", "d", "", "Directory to dedupe")
	dedupeCmd.Flags().StringP("out", "o", "", "Output file")
}
