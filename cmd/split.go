package cmd

import (
	"log"
	"strings"

	"github.com/netr/listy/iom"
	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split file(s) by a delimiter and pluck ids",
	Long:  `Split file(s) by a delimiter and pluck ids`,
	Run: func(cmd *cobra.Command, args []string) {
		delim := validateFlag(cmd, "delim")
		ids := validateFlag(cmd, "ids")
		intIds := stringToIntSlice(ids, ",")

		dir := getFlag(cmd, "dir")
		if dir != "" {
			splitDir(cmd, dir, delim, intIds)
			return
		}

		file := validateFlag(cmd, "file")
		out := getFlag(cmd, "out", iom.AppendSuffixToFilename(file, "-split"))

		log.Printf("Spltting %s to %s by %s with ids: %v", file, out, delim, ids)
		newLines, err := iom.SplitByAndPluckIDsFile(file, delim, intIds)
		if err != nil {
			log.Fatal(err)
		}

		concatLines := concatSplitLines(newLines, delim)
		if err = iom.WriteFile(out, concatLines); err != nil {
			log.Fatal(err)
		}

		log.Printf("Wrote %d lines to %s", len(concatLines), out)
	},
}

func splitDir(cmd *cobra.Command, dir string, delim string, ids []int) {
	log.Printf("Spltting directory %s by %s for ids: %v\n\n", dir, delim, ids)

	if dir != "" {
		files, err := iom.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			file = sanitizeFilename(dir + "/" + file)
			out := iom.AppendSuffixToFilename(file, "-split")

			log.Printf("Spltting %s to %s", file, out)
			newLines, err := iom.SplitByAndPluckIDsFile(file, delim, ids)
			if err != nil {
				log.Fatal(err)
			}

			concatLines := concatSplitLines(newLines, delim)
			if err = iom.WriteFile(out, concatLines); err != nil {
				log.Fatal(err)
			}

			log.Printf("Wrote %d lines to %s\n\n", len(concatLines), out)
		}
		return
	}
}

func concatSplitLines(lines [][]string, delim string) []string {
	var nl []string
	for _, v := range lines {
		nl = append(nl, strings.Join(v, delim))
	}
	return nl
}

func init() {
	rootCmd.AddCommand(splitCmd)
	splitCmd.Flags().StringP("file", "f", "", "File to split")
	splitCmd.Flags().StringP("dir", "d", "", "Directory to split")
	splitCmd.Flags().StringP("delim", "s", "", "Delimiter to split by")
	splitCmd.Flags().StringP("ids", "i", "", "IDs to split by")
	splitCmd.Flags().StringP("out", "o", "", "Output file")
}
