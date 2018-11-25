package main

import (
	"fmt"
	"io"
	"path/filepath"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	newLine    = "\n"
	emptySpace = ""
	item       = "├───"
	start      = "│	"
	space      = "	"
	lastItem   = "└───"
)

func dirTree(out io.Writer, dirName string, full bool) error {
	return readDir(out, dirName, full, emptySpace)
}

type ByName []os.FileInfo

func (fileInfo ByName) Len() int      { return len(fileInfo) }
func (fileInfo ByName) Swap(i, j int) { fileInfo[i], fileInfo[j] = fileInfo[j], fileInfo[i] }
func (fileInfo ByName) Less(i, j int) bool {
	return (strings.Compare(fileInfo[i].Name(), fileInfo[j].Name()) < 0)
}


func readDir(out io.Writer, fileName string, full bool, preText string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	// получаю ниформацию о дочерних элементах f
	filesInfo, err := f.Readdir(0)
	if err != nil {
		return err
	}

	if !full {
		temp := filesInfo[:0]
		// удаляю файлы
		for _, info := range filesInfo {
			if info.IsDir() {
				temp = append(temp, info)
			}
		}
		filesInfo = temp
	}

	// сортирую элементы
	sort.Sort(ByName(filesInfo))

	for i, val := range filesInfo {
		var curPreText string
		var nextPreText string

		// проверка является ли элемент последним в списке
		if i+1 != len(filesInfo) {
			curPreText = preText + item
			nextPreText = preText + start
		} else {
			curPreText = preText + lastItem
			nextPreText = preText + space
		}

		filePath := filepath.Join(fileName, val.Name())

		if val.IsDir() {
			fmt.Fprintf(out, curPreText)
			fmt.Fprintf(out, val.Name())
			fmt.Fprintf(out, newLine)
			readDir(out, filePath, full, nextPreText)
		} else {
			if full {
				if val.Size() == 0 {
					fmt.Fprintf(out, curPreText+val.Name()+" (empty)"+newLine)
				} else {
					fmt.Fprintf(out, curPreText+val.Name()+" ("+strconv.FormatInt(val.Size(), 10)+"b)"+newLine)
				}
			}
		}

	}

	return nil
}
