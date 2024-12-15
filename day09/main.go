package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	isEmpty bool
	value   int
	size    int
}

type Disk struct {
	files []int
}

type ProblemInput struct {
	disk Disk
}

func main() {
	inputs := loadInputs("inputs.txt")

	fmt.Println("First part solution: ", firstPart(inputs))
	fmt.Println("Second part solution: ", secondPart(inputs))
}

func firstPart(inputs ProblemInput) int {
	return inputs.disk.getChecksum()
}

func secondPart(inputs ProblemInput) int {
	return inputs.disk.getWholeFileChecksum()
}

func loadInputs(filename string) (inputs ProblemInput) {
	input, _ := os.ReadFile(filename)

	files := []int{}

	// Parse single line of chars to int
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for _, c := range line {
			file, _ := strconv.Atoi(string(c))
			files = append(files, file)
		}
	}

	inputs.disk = Disk{files: files}

	return inputs
}

func getHeighestKey(m map[int]int) int {
	// Get the highest key in the map
	highestKey := 0
	for key := range m {
		if key > highestKey {
			highestKey = key
		}
	}

	return highestKey
}

// Disk methods
func (d Disk) getChecksum() int {
	total := 0

	for i, file := range d.compact() {
		total += i * file
	}

	return total
}

func (d Disk) getWholeFileChecksum() int {
	total := 0

	for i, file := range d.compactWithWholeFile() {
		total += i * file
	}

	return total
}

func (d Disk) compact() []int {
	// Contains the compacted files, map of int -> int
	compacted := map[int]int{}
	// Contains the indexes of the free spaces
	freeSpaces := []int{}

	idNumber := 0
	k := 0
	// Fill the collections
	for i, file := range d.files {
		// Check if the file is empty : 0 or even
		isEmpty := i%2 != 0

		// for 0 to file
		for j := 0; j < file; j++ {
			// if the file is empty
			if isEmpty {
				// add the index to the free spaces
				freeSpaces = append(freeSpaces, k)
			} else {
				// add the file to the compacted files
				compacted[k] = idNumber
			}
			k++
		}

		if !isEmpty {
			idNumber++
		}
	}

	// Use the free spaces to fill the compacted files
	// for each free space, fill it with right most value of compacted files
	for _, freeSpace := range freeSpaces {
		// get highest key in compacted files
		rightMost := getHeighestKey(compacted)
		// get the value of the right most key
		rightMostValue := compacted[rightMost]
		// add the right most value to the free space
		compacted[freeSpace] = rightMostValue
		// remove the right most value from the compacted files
		delete(compacted, rightMost)
	}

	// Sort the keys and create a slice of all values
	keys := slices.Collect(maps.Keys(compacted))
	sort.Ints(keys)

	allValues := []int{}
	for _, key := range keys {
		allValues = append(allValues, compacted[key])
	}

	return allValues
}

func (d Disk) compactWithWholeFile() []int {
	// Contains the compacted files, map of int -> int
	files := []File{}

	idNumber := 0
	for i, file := range d.files {
		// Check if the file is empty : 0 or even
		isEmpty := i%2 != 0
		size := file

		value := 0
		if !isEmpty {
			value = idNumber
		}

		files = append(files, File{isEmpty: isEmpty, value: value, size: size})

		if !isEmpty {
			idNumber++
		}
	}

	compacted := make([]File, len(files))

	for i := 0; i < len(files); i++ {
		file := files[i]
		if !file.isEmpty {
			compacted = append(compacted, file)
			continue
		}

		for j := len(files) - 1; j > i; j-- {
			rightFile := files[j]
			if rightFile.isEmpty {
				continue
			}

			if rightFile.size > file.size {
				continue
			}

			rightFile.isEmpty = true

			if rightFile.size == file.size {
				file.value = rightFile.value
				file.isEmpty = false
				files[j].isEmpty = true
				break
			}

			// if less, change the empty file size and add the right file before the empty file

			file.size = file.size - rightFile.size
			files[j].isEmpty = true

			// insert the right file before the empty file
			newFile := File{isEmpty: false, value: rightFile.value, size: rightFile.size}
			compacted = append(compacted, newFile)
		}

		compacted = append(compacted, file)
	}

	// Create a slice of all values
	allValues := []int{}
	for _, file := range compacted {
		value := 0
		if !file.isEmpty {
			value = file.value
		}

		for i := 0; i < file.size; i++ {
			allValues = append(allValues, value)
		}
	}

	return allValues
}
