package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
)

func excuteTakeJSON(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		result := ""
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
				result = string(b)
			}
		}
		if flags.fallbackPrint && len(s) == 0 {
			result = txt
		}
		if result != "" {
			result += "\n"
		}
		_, e := fmt.Fprint(w, result)
		if e != nil {
			return e
		}
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
