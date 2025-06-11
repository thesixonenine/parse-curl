# parse-curl
Parse curl commands, returning an object representing the request.

## Example
```shell
curl 'https://example.com' \
  -H 'Accept: text/plain' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36' \
  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' \
  -H 'Connection: keep-alive' \
  --compressed
```
out
```json
{
  "method": "GET",
  "url": "https://example.com",
  "header": {
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
    "Connection": "keep-alive",
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36"
  },
  "body": ""
}
```

## Reference
* https://github.com/tj/parse-curl.js
