package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func excuteTakeJSON(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	results := []string{}
	for scanner.Scan() {
		txt := scanner.Text()
		if flags.decodeURI {
			q, err := url.QueryUnescape(txt)
			if err != nil {
				fmt.Println("error:", err)
			}
			txt = q
		}
		r := regexp.MustCompile(`{.*}`)
		s := r.FindAllString(txt, -1)
		for _, v := range s {
			b, err := json.MarshalIndent(json.RawMessage(v), "", "\t")
			if err != nil {
				fmt.Println("error:", err)
			}
			if len(b) > 0 {
				results = append(results, string(b))
			}
		}
		if flags.fallbackPrint && len(results) == 0 {
			results = append(results, txt)
		}
	}
	_, e := fmt.Fprintln(w, strings.Join(results, ",\n"))
	if e != nil {
		return e
	}
	return nil
}

func (r *takeJSON) excute() error {
	if isInputFromPipe() {
		return excuteTakeJSON(os.Stdin, os.Stdout)
	}
	file, e := getFile()
	if e != nil {
		return e
	}
	defer file.Close()
	return excuteTakeJSON(file, os.Stdout)
}
