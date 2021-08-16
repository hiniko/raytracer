package rt

import (
	"errors"
	"fmt"
	"io/fs"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const SMALL_NUMBER_F64 float64 = 0.000001

func Equal(a float64, b float64) bool {
	return math.Abs(a-b) < SMALL_NUMBER_F64
}

// Convert float 64 To int, clamping between 0 and 255 and rounding up
func F64ToStr_RGB255(f float64) string {
	return strconv.Itoa(int(Clamp((math.Ceil(f * 255)), 0, 255)))
}

func Clamp(v, lo, hi float64) float64 {
	switch {
	case (v < lo):
		return lo
	case (v > hi):
		return hi
	default:
		return v
	}

}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info != nil && !info.IsDir()
}

func WriteFile(filename string, contents string) error {

	// Check the dir exists
	dir := path.Dir(filename)

	// Check for filename collision, append name if needed
	_, err := os.Stat(dir)

	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, fs.ModeDir)
		if err != nil {
			fmt.Printf("failed to create directory %s: %s", dir, err.Error())
		}
	}

	// WHY IS THIS BROKE?
	fn := path.Base(filename)

	if FileExists(filename) {
		ts := strconv.FormatInt(time.Now().Unix(), 10)
		parts := strings.Split(path.Base(filename), ".")

		if len(parts) > 1 {
			fn = parts[0] + "_" + ts + "." + parts[1]
		}
	}

	// create file and open handle
	f, err := os.Create(path.Join(dir, fn))

	if err != nil {
		fmt.Printf("failed to open `%s`! %s", fn, err.Error())
		return err
	} else {
		defer f.Close()
	}

	// Write contents
	n, err := f.Write([]byte(contents))

	// Check for errors
	if err != nil {
		fmt.Printf("Failed to write data to file %s: %s", fn, err.Error())
		return err
	} else {
		fmt.Printf("Wrote %d bytes to %s \n", n, fn)
		return nil
	}
}
