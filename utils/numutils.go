package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	s2 := strings.TrimSpace(s)
	num, err2 := strconv.ParseFloat(s2, 64)
	if err2 != nil {
		log.Fatal(err2)
	}

	return num, nil
}
