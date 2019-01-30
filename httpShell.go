package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"unicode"
)

type env struct {
	url *url.URL
}

func main() {
	fmt.Println("Not Yet Implemented")
	os.Exit(1)
}

func newEnv(rawurl string) (env, error) {
	e := env{}
	u, err := url.Parse(rawurl)
	if err != nil {
		return e, err
	}

	e.url = u

	return e, nil
}

func (e *env) shell() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("httpShell %s > ", e.url.Host)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Print(err)
			continue
		}

		e.parseLine(bytes.Trim(line, " "))
	}
}

func (e *env) parseLine(bline []byte) error {
	// Tokenize
	tokens := e.bytesToTok(bline)
	// Execute
	e.execute(tokens)
	return nil
}

func (e *env) bytesToTok(bline []byte) []string {
	token := []byte{}
	tokens := []string{}

	for i, b := range bline {
		if !unicode.IsSpace(rune(b)) {
			if b == '[' || b == '{' { // JSON data
				token = append(token, bline[i:]...)
				tokens = append(tokens, string(token))
				break
			}
			token = append(token, b)
		} else {
			tokens = append(tokens, string(token))
			token = []byte{}
		}
	}

	return tokens
}

func (e *env) execute(tok []string) error {
	// tok should be { VERB, [URL], [PATH], [DATA] }

	switch strings.ToUpper(tok[0]) {
	case "CONNECT":
	case "GET":
	case "POST":
	case "PUT":
	case "PATCH":
	case "DELETE":
	case "DISCONNECT":
	default:
		return errors.New("Unknown operation")
	}

	return nil
}
