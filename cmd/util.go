package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func getFlag(cmd *cobra.Command, flag string, defaultValue ...string) string {
	val := cmd.Flags().Lookup(flag).Value.String()
	if val == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return val
}

func getFlagInt(cmd *cobra.Command, flag string, defaultValue ...int) int {
	val := cmd.Flags().Lookup(flag).Value.String()
	if val == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func validateFlag(cmd *cobra.Command, flag string) string {
	val := getFlag(cmd, flag)
	if val == "" {
		log.Fatalf("Please provide a value for --%s", flag)
	}

	return val
}

func stringToIntSlice(s string, delim string) []int {
	slice := strings.Split(s, delim)
	intSlice := make([]int, len(slice))
	for i, v := range slice {
		intSlice[i], _ = strconv.Atoi(v)
	}
	return intSlice
}

// sanitizeFilename replaces any repeating "/"'s with only one "/" in a filename
// and removes any leading or trailing "/"'s. If the filename is empty, it will return "".
// Example: directory////directory////file.txt -> directory/directory/file.txt
func sanitizeFilename(filename string) string {
	filename = strings.Replace(filename, "//", "/", -1)
	if filename == "/" {
		return ""
	}
	return strings.Trim(filename, "/")
}
