
# httphash
httphash makes parallel requests for a given list of hosts. The tool requested 
address along with the MD5 hash of the response.

- httphash can do up to 10 parallel requests by default.

- you can specify number of parallel requests by using `-parallel` flag. 
**Example:**
```
./httphash -parallel 2 google.com facebook.com www.sapo.pt http://yahoo.com
```

# Requirements
  Before you begin you must have Go installed and configured properly for your
  computer. Please see https://golang.org/doc/install

# Usage
1) Clone this repository or use `go get github.com/jmcalves275/httphash`
2) Go into project folder: `cd /Users/{username}/go/src/github.com/jmcalves275/httphash`
3) Build everything: `go build`
4) Execute the app: `./httphash google.com`
