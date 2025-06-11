package main

import (
	"fmt"
	parseCurl "github.com/thesixonenine/parse-curl"
)

func main() {
	request, _ := parseCurl.Parse("curl 'https://example.com' \\\n  -H 'Accept: text/plain' \\\n  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36' \\\n  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' \\\n  -H 'Connection: keep-alive' \\\n  --compressed\n")
	fmt.Println(request.ToJson(true))
}
