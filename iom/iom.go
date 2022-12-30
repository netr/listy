package iom

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// CreateFile creates a file
func CreateFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	return f.Close()
}

// FileExists checks if a file exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// GetFileExtension returns the extension of a file
func GetFileExtension(file string) string {
	return filepath.Ext(file)
}

// AppendSuffixToFilename appends a suffix to a filename before the extension
func AppendSuffixToFilename(filename, suffix string) string {
	if filename == "" {
		return ""
	}
	if !strings.Contains(filename, ".") {
		return filename + suffix
	}
	return filename[:len(filename)-len(GetFileExtension(filename))] + suffix + GetFileExtension(filename)
}

// ReadFile reads a file and returns the contents as a []string
func ReadFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// ReadFileToMap reads a file and returns the contents as a map[string]struct{}
func ReadFileToMap(file string) (map[string]struct{}, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}
	defer f.Close()

	m := make(map[string]struct{})
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		m[scanner.Text()] = struct{}{}
	}

	return m, scanner.Err()
}

// ReadDirToMap reads a directory and returns the contents as a map[string]struct{}
func ReadDirToMap(dir string) (map[string]struct{}, error) {
	files, err := ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading dir: %w", err)
	}

	m := make(map[string]struct{})
	for _, file := range files {
		lines, err := ReadFile(dir + "/" + file)
		if err != nil {
			return nil, fmt.Errorf("reading dir: %w", err)
		}
		for _, line := range lines {
			m[line] = struct{}{}
		}
	}

	return m, nil
}

// WriteFile writes a []string to a file
func WriteFile(file string, lines []string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	w := bufio.NewWriter(f)
	sb := strings.Builder{}
	for _, line := range lines {
		sb.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("write file: %w", err)
		}
	}

	_, err = w.WriteString(sb.String())
	w.Flush()
	return err
}

// AppendFile appends a []string to a file
func AppendFile(file string, lines []string) error {
	if !FileExists(file) {
		err := CreateFile(file)
		if err != nil {
			return fmt.Errorf("append file: %w", err)
		}
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("append file: %w", err)
	}

	w := bufio.NewWriter(f)
	sb := strings.Builder{}
	for _, line := range lines {
		sb.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("append file: %w", err)
		}
	}

	_, err = w.WriteString(sb.String())
	return err
}

// ReadDir reads a directory and returns the contents as a []string
func ReadDir(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	files, err := f.Readdirnames(0)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	return files, nil
}

// ReadDirExt reads a directory and returns the contents as a []string
func ReadDirExt(dir, ext string) ([]string, error) {
	var files []string
	files, err := ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir ext: %w", err)
	}

	var filtered []string
	for _, file := range files {
		if filepath.Ext(file) == ext {
			filtered = append(filtered, file)
		}
	}

	return filtered, nil
}

// ReadDirExts reads a directory and returns the contents as a []string
func ReadDirExts(dir string, exts []string) ([]string, error) {
	var files []string
	files, err := ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir ext: %w", err)
	}

	var filtered []string
	for _, file := range files {
		for _, ext := range exts {
			if filepath.Ext(file) == ext {
				filtered = append(filtered, file)
			}
		}
	}

	return filtered, nil
}

// ReadDirRecursive reads a directory recursively and returns the contents as a []string
func ReadDirRecursive(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("read dir recursive: %w", err)
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("read dir recursive: %w", err)
	}

	return files, nil
}

// ReadDirRecursiveWithExt reads a directory recursively and returns the contents as a []string
func ReadDirRecursiveWithExt(dir, ext string) ([]string, error) {
	var files []string
	files, err := ReadDirRecursive(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir recursive with ext: %w", err)
	}

	var filtered []string
	for _, file := range files {
		if filepath.Ext(file) == ext {
			filtered = append(filtered, file)
		}
	}

	return filtered, nil
}

// ReadDirRecursiveWithExts reads a directory recursively and returns the contents as a []string
func ReadDirRecursiveWithExts(dir string, exts []string) ([]string, error) {
	var files []string
	files, err := ReadDirRecursive(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir recursive with ext: %w", err)
	}

	var filtered []string
	for _, file := range files {
		for _, ext := range exts {
			if filepath.Ext(file) == ext {
				filtered = append(filtered, file)
			}
		}
	}

	return filtered, nil
}

// CopyFile copies a file
func CopyFile(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("copy file: %w", err)
	}
	defer f.Close()

	d, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("copy file: %w", err)
	}
	defer d.Close()

	_, err = io.Copy(d, f)
	return err
}

// CopyDir copies a directory
func CopyDir(src, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return fmt.Errorf("copy dir: %w", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("copy dir: %w", err)
			}
		} else {
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("copy dir: %w", err)
			}
		}
	}

	return nil
}

// MoveFile moves a file
func MoveFile(src, dst string) error {
	err := CopyFile(src, dst)
	if err != nil {
		return fmt.Errorf("move file: %w", err)
	}

	err = os.Remove(src)
	return err
}

// MoveDir moves a directory
func MoveDir(src, dst string) error {
	err := CopyDir(src, dst)
	if err != nil {
		return fmt.Errorf("move dir: %w", err)
	}

	err = os.RemoveAll(src)
	return err
}

// RemoveFile removes a file
func RemoveFile(path string) error {
	return os.Remove(path)
}

// RemoveDir removes a directory
func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

// ShuffleStrings shuffles a []string
func ShuffleStrings(s []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleFile shuffles a file
func ShuffleFile(src, dst string) error {
	lines, err := ReadFile(src)
	if err != nil {
		return fmt.Errorf("shuffle file: %w", err)
	}

	ShuffleStrings(lines)

	err = WriteFile(dst, lines)
	if err != nil {
		return fmt.Errorf("shuffle file: %w", err)
	}

	return nil
}

// RemoveDuplicates removes duplicate lines from a []string
func RemoveDuplicates(s []string) (int, []string) {
	var result []string
	seen := make(map[string]bool)

	dupes := 0
	for _, v := range s {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		} else {
			dupes++
		}
	}

	return dupes, result
}

// RemoveDuplicatesFile removes duplicate lines from a file
func RemoveDuplicatesFile(src, dst string) (int, error) {
	var n int

	lines, err := ReadFile(src)
	if err != nil {
		return n, fmt.Errorf("remove duplicates file: %w", err)
	}

	n, lines = RemoveDuplicates(lines)

	err = WriteFile(dst, lines)
	if err != nil {
		return n, fmt.Errorf("remove duplicates file: %w", err)
	}

	return n, nil
}

// SplitBy splits a []string by a delimiter and returns a [][]string
func SplitBy(s []string, delim string) [][]string {
	var result [][]string

	for _, line := range s {
		result = append(result, strings.Split(line, delim))
	}

	return result
}

// SplitByFile splits a file by a delimiter and returns a [][]string
func SplitByFile(src, delim string) ([][]string, error) {
	lines, err := ReadFile(src)
	if err != nil {
		return nil, fmt.Errorf("split by file: %w", err)
	}

	result := SplitBy(lines, delim)

	return result, nil
}

// SplitByAndPluckIDs splits a []string by a delimiter and returns a [][]string for specific indices
func SplitByAndPluckIDs(s []string, delim string, ids []int) [][]string {
	var result [][]string

	for _, line := range s {
		split := strings.Split(line, delim)
		var plucked []string
		for _, id := range ids {
			plucked = append(plucked, split[id])
		}
		result = append(result, plucked)
	}

	return result
}

// SplitByAndPluckIDsFile splits a file by a delimiter and returns a [][]string for specific indices
func SplitByAndPluckIDsFile(src, delim string, ids []int) ([][]string, error) {
	lines, err := ReadFile(src)
	if err != nil {
		return nil, fmt.Errorf("split by and pluck ids file: %w", err)
	}

	result := SplitByAndPluckIDs(lines, delim, ids)

	return result, nil
}

// CountFileLines counts the number of lines in a file
func CountFileLines(path string) (int, error) {
	lines, err := ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("count file lines: %w", err)
	}

	return len(lines), nil
}

// ChunkFile splits a file into chunks
func ChunkByLinesFile(src, dst string, chunkSize int) (int, error) {
	lines, err := ReadFile(src)
	if err != nil {
		return 0, fmt.Errorf("chunk file: %w", err)
	}
	chunks := ChunkByLines(lines, chunkSize)

	for i, chunk := range chunks {
		fname := AppendSuffixToFilename(dst, "-"+strconv.Itoa(i+1))
		err = WriteFile(fname, chunk)
		if err != nil {
			return 0, fmt.Errorf("chunk file: %w", err)
		}
	}

	return len(chunks), nil
}

// Chunk splits a []string into chunks
func ChunkByLines(s []string, chunkSize int) [][]string {
	var result [][]string

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize

		if end > len(s) {
			end = len(s)
		}

		result = append(result, s[i:end])
	}

	return result
}

func Diff(src1 map[string]struct{}, src2 []string) (int, []string) {
	var result []string
	var n int
	for _, line := range src2 {
		if _, ok := src1[line]; !ok {
			result = append(result, line)
			n++
		}
	}

	return n, result
}

// DiffFiles returns the difference between two files
func DiffFiles(src1, src2, out string) (int, error) {
	baseMap, err := ReadFileToMap(src1)
	if err != nil {
		return 0, fmt.Errorf("diff files: %w", err)
	}

	lines, err := ReadFile(src2)
	if err != nil {
		return 0, fmt.Errorf("diff files: %w", err)
	}

	n, result := Diff(baseMap, lines)

	err = WriteFile(out, result)
	if err != nil {
		return 0, fmt.Errorf("diff files: %w", err)
	}

	return n, nil
}
