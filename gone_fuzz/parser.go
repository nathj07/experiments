// +build gofuzz

package gone_fuzz

import (
	"fmt"
	"strconv"
	"strings"
)

// func main() {
// 	// read the file
// 	inFile, _ := os.Open("testdata/shapes.txt")
// 	defer inFile.Close()
// 	scanner := bufio.NewScanner(inFile)
// 	scanner.Split(bufio.ScanLines)
//
// 	for scanner.Scan() {
// 		// call Parse on each line
// 		item, err := Parse(scanner.Bytes())
// 		if err != nil {
// 			fmt.Printf("Error parsing: %v\n", err)
// 			continue
// 		}
// 		fmt.Println(item.Type())
// 	}
//
// }

type Item interface {
	Type() string
}

type Rectangle struct {
	x, y, w, h int
}

func (r Rectangle) Type() string {
	return "Rectangle"
}

type Circle struct {
	x, y, r int
}

func (c Circle) Type() string {
	return "Circle"
}

type Triangle struct {
	e1, e2, e3 int
}

func (tri Triangle) Type() string {
	return "Triangle"
}

func Parse(line []byte) (Item, error) {
	parts := strings.Split(fmt.Sprintf("%s", line), " ")
	if len(parts) == 0 {
		return nil, fmt.Errorf("No data!!!")
	}
	var item Item
	switch strings.ToLower(parts[0]) {
	case "rectangle":
		if len(parts) < 5 {
			return nil, fmt.Errorf("Unsupported length %d for type %q", len(parts), parts[0])
		}
		x, _ := strconv.Atoi(parts[1])
		y, _ := strconv.Atoi(parts[2])
		w, _ := strconv.Atoi(parts[3])
		h, _ := strconv.Atoi(parts[4])
		item = Rectangle{x, y, w, h}
	case "circle":
		if len(parts) < 4 {
			return nil, fmt.Errorf("Unsupported length %d for type %q", len(parts), parts[0])
		}
		x, _ := strconv.Atoi(parts[1])
		y, _ := strconv.Atoi(parts[2])
		r, _ := strconv.Atoi(parts[3])
		item = Circle{x, y, r}
	case "triangle":
		if len(parts) < 4 {
			return nil, fmt.Errorf("Unsupported length %d for type %q", len(parts), parts[0])
		}
		e1, _ := strconv.Atoi(parts[1])
		e2, _ := strconv.Atoi(parts[2])
		e3, _ := strconv.Atoi(parts[3])
		item = Triangle{e1, e2, e3}
	default:
		return nil, fmt.Errorf("Unknown Type: %q", parts[0])
	}
	return item, nil
}

// Fuzzing

func Fuzz(data []byte) int {
	// call Parse on each line
	_, err := Parse(data)
	if err != nil {
		return 0
	}

	return 1
}
