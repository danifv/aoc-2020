package readstringlines

import (
	"bufio"
	"os"
)

type scanUnitParser func(scanUnit string) interface{}

// ReadStringLines reads file from path, splits it using a parametrizable SplitFunc and parses the splits with a 'scanUnitParser func(scanUnit string) interface{}'
// the interface{} results of this parse are appended to a slice and finally returned by ReadStringLines
func ReadStringLines(path string, splitFunc bufio.SplitFunc, scanUnitParser scanUnitParser) ([]interface{}, error) {
	memoryObjects := make([]interface{}, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		scanUnit := scanner.Text()
		memoryObject := scanUnitParser(scanUnit)
		if err != nil {
			return nil, err
		}
		memoryObjects = append(memoryObjects, memoryObject)
	}
	return memoryObjects, nil
}
