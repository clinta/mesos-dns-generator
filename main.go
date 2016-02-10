package main

import (
        "crypto/sha1"
        "github.com/tv42/zbase32"
        "flag"
        "fmt"
        "strings"
)

var taskid = flag.String("taskid", "", "The task ID")
var slaveid = flag.String("slaveid", "", "The Slave ID")

func main () {
        flag.Parse()
        fmt.Println(hashString(*taskid) + "-" + slaveIDTail(*slaveid))
}

func hashString(s string) string {
	hash := sha1.Sum([]byte(s))
	return zbase32.EncodeToString(hash[:])[:5]
}

func slaveIDTail(slaveID string) string {
        fields := strings.Split(slaveID, "-")
        return strings.ToLower(fields[len(fields)-1])
}
