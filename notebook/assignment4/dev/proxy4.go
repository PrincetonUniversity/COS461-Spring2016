package main

import (
	//"compress/gzip" // uncomment this line to use the gzip package
	"flag"
	"fmt"
	//"golang.org/x/net/html" // uncomment this line to use the html package
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

// Global constants
const (
	SERVERROR	= 500
	BADREQ		= 400
	MAX_OBJ_SIZE	= 500*1024
	MAX_CACHE_SIZE	= 10*1024*1024
)

// Command line parameters
var (
	listeningPort	uint
	dnsPrefetching	bool
	caching		bool
	cacheTimeout	uint
	maxCacheSize	uint
	maxObjSize	uint
	linkPrefetching	bool
	maxConcurrency	uint
	outputFile	string
)

// Channel to synchronize number of prefetch threads
var semConc chan bool

// Stat variables
var (
	clientRequests	int	// HTTP requests
	cacheHits	int	// Cache Hits
	cacheMisses	int	// Cache misses
	cacheEvictions	int	// Cache evictions
	trafficSent	int	// Bytes sent to clients
	volumeFromCache	int	// Bytes sent from the cache
	downloadVolume	int	// Bytes downloaded from servers
)

// RW lock for the stat variables. 
// You need to lock the stat variables when updating them.
var statLock	sync.RWMutex

func saveStatistics() {
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Error creating output file", outputFile)
	}
	start := time.Now()
	str := "#Time(s)\tclientRequests\tcacheHits\tcacheMisses\tcacheEvictions" +
		"\ttrafficSent\tvolumeFromCache\tdownloadVolume\ttrafficWastage\tcacheEfficiency";
	f.WriteString(str)
	for {
		var cacheEfficiency	float64
		var trafficWastage	int
		
		currentTime := time.Now().Sub(start)
		statLock.RLock()
		if trafficSent > 0 {
			cacheEfficiency = float64(volumeFromCache) / float64(trafficSent)
		} else {
			cacheEfficiency = 0.0
		}
		if downloadVolume > trafficSent {
			trafficWastage = downloadVolume-trafficSent
		} else {
			trafficWastage = 0;
		}
		
		str := fmt.Sprintf("\n%d\t\t%d\t\t%d\t\t%d\t\t%d\t\t%d\t\t%d\t\t%d\t\t%d\t\t%f",
			int(currentTime.Seconds()),	clientRequests,
			cacheHits, cacheMisses, cacheEvictions,
			trafficSent, volumeFromCache, downloadVolume,
			trafficWastage, cacheEfficiency)
		statLock.RUnlock()
		f.WriteString(str)
		f.Sync()
		time.Sleep(time.Second * 10)		
	}
}

func doLinkPrefetching(/* Declare the parameters you need. */) {
}

func doDnsPrefetching(/* Declare the parameters you need. */ ) {
}

func handleRequest(w net.Conn) {
	defer w.Close()
	var resp http.Response

	// TODO: Handle http request
	if resp.StatusCode == 200 {
		if linkPrefetching {
			go doLinkPrefetching(/* Pass the parameters you need. */)
		} else if dnsPrefetching {
			go doDnsPrefetching(/* Pass the parameters you need. */) 
		}
	}
}

func initFlags() {
	flag.UintVar(&listeningPort, "port", 8080, "Proxy listening port")
	flag.BoolVar(&dnsPrefetching, "dns", false, "Enable DNS prefetching")
	flag.BoolVar(&caching, "cache", false, "Enable object caching")
	flag.UintVar(&cacheTimeout, "timeout", 120, "Cache timeout in seconds")
	flag.UintVar(&maxCacheSize, "max_cache", MAX_CACHE_SIZE, "Maximum cache size")
	flag.UintVar(&maxObjSize, "max_obj", MAX_OBJ_SIZE, "Maximum object size")
	flag.BoolVar(&linkPrefetching, "link", false, "Enable link prefetching")
	flag.UintVar(&maxConcurrency, "max_conc", 10, "Number of threads for link prefetching")
	flag.StringVar(&outputFile, "file", "proxy.log", "Output file name")
	flag.Parse()
}

func main() {
	initFlags()

	go saveStatistics()

	// TODO: Other initializations
	
	for {
		// TODO: Main loop for accepting connections.
		//go handleRequest(conn)
	}
}
