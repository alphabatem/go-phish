package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

/**
 * Rebuild the blocklist from the programs logs
 * Each block_request log entry increments the count for a site URI - If its above the given threshold its added to the blocklist
 */

type LogReBuilder struct {
	threshold    int64
	file         *csv.Writer
	fileLocation string
	uriCache     map[string]int64
}

func main() {
	fileLoc := "./blocklist.csv"
	w, err := os.Create(fileLoc)

	b := LogReBuilder{
		threshold:    5,
		file:         csv.NewWriter(w),
		fileLocation: fileLoc,
		uriCache:     map[string]int64{},
	}

	//TODO Get logs
	logs := []string{}

	for _, l := range logs {
		b.onLog(l)
	}

	//Add any remaining files not flushed in based on their count
	err = b.buildFile()
	if err != nil {
		log.Fatalf("Unable to build file: %s", err)
		return
	}

	log.Printf("File Rebuilt: %s", b.fileLocation)
	log.Printf("")
}

func (b *LogReBuilder) onLog(log string) {
	uri := "" //TODO Get uri from log

	if _, ok := b.uriCache[uri]; !ok {
		b.uriCache[uri] = 1
		return
	}

	//Existing uri
	b.uriCache[uri]++

	//if b.uriCache[uri] >= b.threshold {
	//	b.write(uri, b.uriCache[uri])
	//}
}

func (b *LogReBuilder) buildFile() error {
	for uri, count := range b.uriCache {
		err := b.write(uri, count)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *LogReBuilder) write(uri string, count int64) error {
	err := b.file.Write([]string{uri, fmt.Sprintf("%v", count)})
	if err != nil {
		return err
	}

	delete(b.uriCache, uri) //Keep nice and lean
	return nil
}
