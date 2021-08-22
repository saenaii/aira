package cyclomatic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"aira/config"
)

const (
	httpTimeout = time.Second * 5
)

var handler = map[string]func(map[string][]Cyclo) error{
	config.OutputStdout: handleStdout,
	config.OutputHTTP:   handleHTTP,
}

var handleStdout = func(input map[string][]Cyclo) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

var handleHTTP = func(input map[string][]Cyclo) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: httpTimeout}
	req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("http code is not 200" + strconv.Itoa(resp.StatusCode))
	}
	return nil
}
