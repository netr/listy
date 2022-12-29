package cmd

import (
	"log"

	"github.com/netr/listy/iom"
	"github.com/spf13/cobra"
)

// chunkCmd represents the chunk command
var chunkCmd = &cobra.Command{
	Use:   "chunk",
	Short: "Chunk file(s) by a given number of lines",
	Long:  `Chunk file(s) by a given number of lines`,
	Run: func(cmd *cobra.Command, args []string) {
		file := validateFlag(cmd, "file")
		chunk := getFlagInt(cmd, "by")

		out := getFlag(cmd, "out", iom.AppendSuffixToFilename(file, "-chunk"))

		log.Printf("Chunking %s to %s in %d chunks", file, out, chunk)
		n, err := iom.ChunkByLinesFile(file, out, chunk)
		if err != nil {
			log.Fatal(err)
		}

		lines, err := iom.CountFileLines(file)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Chunked %d lines into %d files", lines, n)
	},
}

func init() {
	rootCmd.AddCommand(chunkCmd)
	chunkCmd.Flags().StringP("file", "f", "", "File to dedupe")
	chunkCmd.Flags().StringP("dir", "d", "", "Directory to dedupe")
	chunkCmd.Flags().StringP("out", "o", "", "Output file")
	chunkCmd.Flags().IntP("by", "b", 2000, "Chunk by")
}
