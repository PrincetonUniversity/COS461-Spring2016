package main

import (
	//"compress/gzip" // Uncomment to use the gzip package.
	"flag"
	"fmt"
	"log"
	//"net"
	//"net/http"
	//"golang.org/x/net/html" // Uncommen to use the html package.
	"os"
	"time"
)

type TimeLog struct {
	start	time.Time
	end	time.Time
	label	int
}

// The variables below are used to create labels for the time log
var (
	dnsLabel	int	// Sequential number to label DNS map entries
	tcpLabel	int	// Sequential number to label TCP-connect map entries
	loadLabel	int	// Sequential number to label load-time map entries.
)

var dnsMap	map[string]string
var dnsTimeMap	map[string]TimeLog
var tcpTimeMap	map[string]TimeLog
var loadTimeMap	map[string]TimeLog

var (
	targetURL	string	// URL to fetch
	waterfall	bool	// If true, generates data for the waterfall chart
	browser		bool	// If true, generates data for the proxy measurement
	outputFile	string	// Output file name
	numRequests	uint	// Maximum number of requests for linked pages
	initTime	uint	// Sleep time before starting 
	sleepTime	uint	// Sleep time between page fetches
	proxy		string	// IP address and port # of the proxy (ex: 192.168.0.2:8080)
)


func printWaterFall(basename string, start time.Time) {
	f, err := os.Create(basename + ".log")
	if err != nil {
		fmt.Println("Could not create file", basename + ".log")
		return
	}
	defer f.Close()
	
	g, err := os.Create(basename)
	if err != nil {
		fmt.Println("Could not create file", basename)
		return
	}
	defer g.Close()

	f.WriteString("DNS mapping:\n")
	for k, v := range dnsMap {
		str := fmt.Sprintf("%s %s\n", k, v);
		f.WriteString(str)
	}
	f.WriteString("\nTimes\n")
	for k, v := range dnsTimeMap {
		delta := v.end.Sub(v.start)
		begin := v.start.Sub(start).Seconds()*1000
		end   := v.end.Sub(start).Seconds()*1000
		str := fmt.Sprintf("%f - DNS(%s) = %f ms\n", begin, k, delta.Seconds()*1000)
		f.WriteString(str)
		str = fmt.Sprintf("D%02d\t%f\t%f\n", v.label, begin, end)
		g.WriteString(str)
	}
	for k, v := range tcpTimeMap {
		delta := v.end.Sub(v.start)
		begin := v.start.Sub(start).Seconds()*1000
		end   := v.end.Sub(start).Seconds()*1000
		str := fmt.Sprintf("%f - TCP(%s) = %f ms\n", begin, k, delta.Seconds()*1000)
		f.WriteString(str)
		str = fmt.Sprintf("T%02d\t%f\t%f\n", v.label, begin, end)
		g.WriteString(str)
	}
	for k, v := range loadTimeMap {
		delta := v.end.Sub(v.start)
		begin := v.start.Sub(start).Seconds()*1000
		end   := v.end.Sub(start).Seconds()*1000
		str := fmt.Sprintf("%f - URL(%s) = %f ms\n", begin, k, delta.Seconds()*1000)
		f.WriteString(str)
		str = fmt.Sprintf("L%02d\t%f\t%f\n", v.label, begin, end)
		g.WriteString(str)
	}
}

func doWaterfall(url string) {
	start := time.Now()

	// TOOD: Implement the rest of the function
	printWaterFall(outputFile, start)
}

func doBrowser(url string) {
	// TODO: Implement the function
}


func initFlags() {
	flag.StringVar(&targetURL, "url", "", "Target URL")
	flag.BoolVar(&waterfall, "waterfall", false, "Generate waterfall chart")
	flag.BoolVar(&browser, "browser", false, "Simulate human browsing")
	flag.StringVar(&outputFile, "file", "log.txt", "Base name for output files")
	flag.UintVar(&numRequests, "num_req", 20, "Number of requests for linked pages")
	flag.UintVar(&initTime, "init_time", 30, "Sleep time before starting.")
	flag.UintVar(&sleepTime, "sleep_time", 5, "Sleep time between page fetches")
	flag.StringVar(&proxy, "proxy", "", "Address of the proxy server")
	flag.Parse()
	if targetURL == "" {
		log.Fatal("You must specify the target URL")
	}
}

func main() {
	initFlags()
	dnsMap      = make(map[string]string)
	dnsTimeMap  = make(map[string]TimeLog)
	tcpTimeMap  = make(map[string]TimeLog)
	loadTimeMap = make(map[string]TimeLog)

	if (!waterfall && !browser) || (waterfall && browser) {
		flag.PrintDefaults()
		log.Fatal("You must specify either waterfall or browser")
	}
	if waterfall {
		doWaterfall(targetURL)
	} else {
		doBrowser(targetURL)
	}
}
