// challenge: https://www.hackerrank.com/challenges/flipping-bits/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func flippingBits(n int64) int64 {
	return 4294967295 - n
}

//Todo dado é armazenado em um binario
//1 bit pode armazenar 0 ou 1
//2 bits podem armazenar 00, 01, 10, 11
//entao 32 bits podem armazenar?? 2^32, de 0 a 4294967295
// um unsigned integer é um numero inteiro e positivo
// um dos valores é zero, então o valor maximo que pode ser armazenado é 2^32 -1
// 4294967295 (base10) = 11111111111111111111111111111111
// qualquer valor que você subtrair do de cima, terá o resultado dos bits invertidos

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := flippingBits(n)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
