
import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func container() {

	intSlice := []int{}
	sliceOfAnything := []time.Time{}

	// String to byte slice and vice versa.
	bytes := []byte("foo")
	msg := string(bytes)

	// Copy a slice
	nCopied := copy(dst, source)

	// Iterating over characters in a string or byte slice

	runes := []rune("foo")
	for i, rune := range runes {
	}
	fmt.Printf("%c", rune)

	// Binary conversion
	stringRepr := strconv.FormatInt(myInt64, 2)
	stringRepr := fmt.Sprintf("%b", myInt64)

	// Bitwise, shifting and masking
	// & | ^ << >>

	// Sort hard coded type (in place)
	sort.Ints(intSlice)

	// Sort generic slice using predicate
	sort.SliceStable(sliceOfAnything, func(i, j int) bool {
		return sliceOfAnything[i].Name < sliceOfAnything[j].Name
	})

	// Binary search
	found := sort.Search(len(intSlice), func(i int) bool {
		// Return true if this (and later i) satisfies.
		return seq[i] != i+1
	})

	// Reg ex match and groups
	r, _ := regexp.Compile("p([a-z]+)ch")
	it := r.FindString(str)
	them := r.FindAllString(str)
	groups := r.FindStringSubmatch(str)
	start, end := r.FindStringIndex(str)
	isPresent := r.MatchString(str)
	new := r.ReplaceAllSring(str, replacement)
	segs := r.Split(str, -1) // -1 means no limit to n returned

	// Random
	rand.Seed(int64)
	f := rand.Float() // typ
	n := rand.Intn(99)
}
