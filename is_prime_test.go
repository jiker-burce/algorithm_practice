package algorithm_practice

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// /*
// * Complete the 'isPrime' function below.
// *
// * The function is expected to return an INTEGER.
// * The function accepts LONG_INTEGER n as parameter.
// */
func isPrime(n int64) int32 {
	// Write your code here
	if n <= 2 {
		return 1
	}
	sqrt := math.Sqrt(float64(n))
	for i := int64(2); i <= int64(sqrt); i++ {
		if n%i == 0 {
			return int32(i)
		}
	}
	return 1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := isPrime(n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
