package cmd

import (
	"log"

	"github.com/netr/listy/iom"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Filter differences between file(s)",
	Long:  `Filter lines of file(s) and output only those that are different from the other file(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		base := validateFlag(cmd, "base")
		dir := getFlag(cmd, "dir")
		if dir != "" {
			diffDir(cmd, base, dir)
			return
		}

		file := validateFlag(cmd, "file")
		out := getFlag(cmd, "out", iom.AppendSuffixToFilename(file, "-diff"))

		n, err := iom.DiffFiles(base, file, out)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Found %d different lines", n)
	},
}

func diffDir(cmd *cobra.Command, base string, dir string) {
	log.Printf("Getting differences from directory %s\n\n", dir)

	if dir != "" {
		files, err := iom.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		var result []string
		baseMap, err := iom.ReadFileToMap(base)
		if err != nil {
			log.Fatal(err)
		}

		var totalN int
		for _, file := range files {
			file = sanitizeFilename(dir + "/" + file)
			if file == base {
				continue
			}

			lines, err := iom.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}

			n, diff := iom.Diff(baseMap, lines)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, diff...)
			totalN += n
			log.Printf("Found %d differences from %s\n", n, file)
		}

		iom.WriteFile(iom.AppendSuffixToFilename(base, "-diff"), result)
		return
	}
}

func init() {
	rootCmd.AddCommand(diffCmd)
	diffCmd.Flags().StringP("base", "b", "", "Base file to compare")
	diffCmd.Flags().StringP("file", "f", "", "File to check against base")
	diffCmd.Flags().StringP("dir", "d", "", "Directory of files to check against base")
	diffCmd.Flags().StringP("out", "o", "", "Output file")
}
