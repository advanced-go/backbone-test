package main

import (
	"bytes"
	"fmt"
	"github.com/advanced-go/stdlib/httpx"
	io2 "github.com/advanced-go/stdlib/io"
	"net/http"
)

const (
	GuidanceUrl    = "http://localhost:8081/github/advanced-go/guidance:resiliency"
	ObservationUrl = "http://localhost:8081/github/advanced-go/observation:timeseries"

	GuidanceResource    = "file://[cwd]/pkg/resource/documents-v1.json"
	ObservationResource = "file://[cwd]/pkg/resource/access-v1.json"
)

func main() {
	initialLoad()

}

func initialLoad() {
	//Put(GuidanceResource, GuidanceUrl, "")
	Put(ObservationResource, ObservationUrl, "")
}

func Put(file, uri, variant string) bool {
	buf, status := io2.ReadFile(file)
	if !status.OK() {
		fmt.Printf("read file err: %v\n", status.Err)
		return false
	}
	reader := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, uri, reader)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return false
	}
	resp, status1 := httpx.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
	fmt.Printf("Put() [status:%v]\n", status1)
	return true
}
