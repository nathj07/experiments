package main

import (
	"fmt"
	"os"
)

func main() {
	lp := []string{"testdata/pdf1.pdf", "testdata/pdf2.pdf", "testdata/subdir"}
	for _, l := range lp {
		// document, err := ioutil.ReadFile(l)
		// if err != nil {
		// 	fmt.Printf("Unable to read local file %q: %v\n", l, err)
		// 	os.Exit(1)
		// }
		// s := int64(simhash.Simhash(simhash.NewWordFeatureSet(document)))
		// fnv := fnv.New64()
		// fnv.Write(document)
		// fp := int64(fnv.Sum64())
		// fmt.Printf("%s simhash = %v; fnv = %v\n", l, s, fp)

		fd, _ := os.Open(l)
		fmt.Println(fd.Name())
		fi, _ := fd.Stat()
		fmt.Println(fi.Name())
		fd.Close()

		_, err := os.Stat("testdata/stat")
		fmt.Println(err)
	}
}
