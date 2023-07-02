package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TestResponse struct {
	Follow string `json:"follow"`
	Message string `json:"message"`
}

func main() {
    fmt.Println("Hello, World!")
    client := &http.Client{}
    req, _ := http.NewRequest("GET", "https://www.letsrevolutionizetesting.com/challenge", nil)
    req.Header.Set("Accept", "application/json")
    res, _ := client.Do(req)
    fmt.Println(res)
    body, _ := ioutil.ReadAll(res.Body)
    var tr TestResponse
    json.Unmarshal(body, &tr)

    for {
	    req, _ = http.NewRequest("GET", tr.Follow, nil)
	    req.Header.Set("Accept", "application/json")
            res, _ = client.Do(req)
	    body, _ = ioutil.ReadAll(res.Body)
	    json.Unmarshal(body, &tr)
	    if (tr.Message != "This is not the end") {
		    fmt.Println(tr.Message)
		    fmt.Println("Done")
		    return
	    }
    }
}
