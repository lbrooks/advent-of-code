package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type file struct {
	name string
	size int
}

type folder struct {
	name    string
	parent  *folder
	folders []*folder
	files   []file
}

func (f file) String() string {
	return fmt.Sprintf("%s (%d)", f.name, f.size)
}

func (f folder) FileSize() int {
	sum := 0
	for _, file := range f.files {
		sum += file.size
	}
	return sum
}

func (f folder) TotalSize() int {
	sum := f.FileSize()
	for _, subFolder := range f.folders {
		sum += subFolder.TotalSize()
	}
	return sum
}

func (f folder) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Name: [%s] Directory Size [%d], Files [%d]", f.name, f.TotalSize(), f.FileSize())

	if len(f.folders) > 0 {
		sb.WriteString("\nFolders: [\n")
		for _, subFolder := range f.folders {
			fmt.Fprintf(&sb, "\t%v\n", subFolder)
		}
		sb.WriteString("]")
	}
	if len(f.files) > 0 {
		sb.WriteString("\nFiles: [ ")
		for _, subFile := range f.files {
			fmt.Fprintf(&sb, " %v ", subFile)
		}
		sb.WriteString(" ]")
	}
	return sb.String()
}

func createDirectory(input []string) *folder {
	root := &folder{}

	workingDirectory := root
	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i], "$") {
			if input[i] == "$ ls" {
				continue
			}
			if input[i] == "$ cd .." {
				workingDirectory = workingDirectory.parent
				continue
			}
			if input[i] == "$ cd /" {
				workingDirectory = root
				continue
			}
			intoFolder := strings.Split(input[i], " ")[2]
			for _, f := range workingDirectory.folders {
				if f.name == intoFolder {
					workingDirectory = f
					break
				}
			}
			continue
		} else if strings.HasPrefix(input[i], "dir") {
			subFolder := &folder{
				name:   strings.Split(input[i], " ")[1],
				parent: workingDirectory,
			}

			workingDirectory.folders = append(workingDirectory.folders, subFolder)
		} else {
			parts := strings.Split(input[i], " ")
			size, _ := strconv.Atoi(parts[0])
			workingDirectory.files = append(workingDirectory.files, file{name: parts[1], size: size})
		}
	}

	return root
}

func (f folder) Part1Sum() int {
	sum := 0
	if mySize := f.TotalSize(); mySize <= 100000 {
		sum += mySize
	}
	for _, sf := range f.folders {
		sum += sf.Part1Sum()
	}
	return sum
}

func (f folder) TrueSum() int {
	sum := f.FileSize()
	for _, sf := range f.folders {
		sum += sf.TrueSum()
	}
	return sum
}

func (f folder) Part2SmallestFolderToFree(minSpace int) int {
	minSizeDelete := f.TrueSum()
	if minSizeDelete < minSpace {
		return -1
	}
	for _, sf := range f.folders {
		sfSize := sf.Part2SmallestFolderToFree(minSpace)
		if sfSize < minSpace {
			continue
		}
		if sfSize < minSizeDelete {
			minSizeDelete = sfSize
		}
	}
	return minSizeDelete
}

func main() {
	rootDirectory := createDirectory(utils.ReadPiped())

	fmt.Printf("Part 1: %d\n", rootDirectory.Part1Sum())

	totalSpace := 70000000
	updateNeeds := 30000000
	totalSpaceUsed := rootDirectory.TrueSum()
	spaceNeeded := updateNeeds - (totalSpace - totalSpaceUsed)
	fmt.Printf("Part 2: %d\n", rootDirectory.Part2SmallestFolderToFree(spaceNeeded))
}
