package handlers

import (
	defs "goroutine-pool/defs"
	//	"fmt"
	"io/ioutil"
	"net/http"
)

func Crawl() error {

	r, err := http.Get(defs.URL)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(r.Body)
	//	fmt.Println(string(h))
	if err != nil {
		return err
	}
	return nil
}
