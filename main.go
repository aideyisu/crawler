package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jackdanger/collectlinks"
)

//Func check checks for errors. If any panic and write log.
func check(s string, e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	return n, nil
}

//init func initializes directory creation and log tracing.
func begin(URL string, SecondPath string, Year string, Month string, Name string) {
	dir, err := url.Parse(URL)
	_, err = os.Stat("files/" + SecondPath + Year + Month + "/")
	if os.IsNotExist(err) {
		errDir2 := os.MkdirAll("files/"+Year+Month+"/", 0777)
		check("Error creating files directory: ", errDir2)
	}

	_, err = os.Stat("logs/" + dir.Host + "/")
	if os.IsNotExist(err) {
		errDir2 := os.MkdirAll("logs/"+dir.Host+"/", 0777)
		check("Error creating log directory: ", errDir2)
	}

	//Log to file
	file, e := os.OpenFile("logs/"+dir.Host+"/"+"file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if e != nil {
		log.Fatalln("Failed to open log file")
	}
	log.SetOutput(file)
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

// DownloadFile 下载文件
func DownloadFile(filename string, u string, SecondPath string, Year string, Month string, Name string) {
	//Open file and append if it exist. If not create it and write.

	newFile, err := os.OpenFile("files/"+Year+Month+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check("Error opening file! ", err)
	defer newFile.Close()
	response, err := http.Get(u)

	if response.StatusCode == 404 {
		log.Println("Can't find target URL!\n Please check that inserted URL is valid.")
		os.Exit(1)
	}
	if response.StatusCode == 200 {
		web, err := url.Parse(u)
		check("Error parsing URL ", err)

		ip, err := net.LookupIP(web.Host)
		check("Error retrieving website's IP address. ", err)
		log.Printf("Connected to %v.\tIP address %v\n", web.Host, ip)
	}

	defer response.Body.Close()

	//Count bytes while downloading and copy response into saved file
	counter := &WriteCounter{}
	if _, err = io.Copy(newFile, io.TeeReader(response.Body, counter)); err != nil {
		newFile.Close()
		return
	}

	response, err = http.Get(u)
	defer response.Body.Close()
	//List all hyperlinks in the downloaded page.
	//We look for href tags only.
	links, err := goquery.NewDocumentFromReader(response.Body)
	check("Error loading links: ", err)
	// Find all links and save to file
	links.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			web, err := url.Parse(u)
			check("Error parsing URL ", err)
			// If the file doesn't exist, create it, or else append to the file
			f, err := os.OpenFile("links/"+web.Host+web.Path+"/"+"link.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
			check("Can't write log file to disk.", err)
			_, err = f.Write([]byte(href + "\n"))
			f.Close() // ignore error; Write error takes precedences
			check("Error writing bytes to file: ", err)
		}
	})
	check("", err)
	log.Println("\nCrawling finished.\t Success!\t" + u)
}

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Usage: " + os.Args[0] + " <fileurl_to_save> <target_secURL> <Year> <Month> <Name>")
		fmt.Println("Example: " + " ./main http://archive.routeviews.org bgpdata 2003 03 UPDATES")
		log.Fatalln("Not right number arguments passed!")
		os.Exit(1)
	}
	URL := os.Args[1]
	SecondPath := os.Args[2]
	Year := os.Args[3]
	Month := os.Args[4]
	Name := os.Args[5]

	FilePath := URL + "/" + SecondPath + "/" + Year + "." + Month + "/" + Name + "/"

	begin(FilePath, SecondPath, Year, Month, Name)
	// FilePath := "http://archive.routeviews.org/bgpdata/2001.10/UPDATES/"
	// ./main http://archive.routeviews.org bgpdata 2001 10 RIBS

	resp, _ := http.Get(FilePath)
	links := collectlinks.All(resp.Body)
	k := 0
	for _, j := range links {
		if strings.HasPrefix(j, strings.ToLower(Name[:len(Name)-1])) {
			log.Printf("The %d one start. %s\n", k+1, FilePath+j)
			fmt.Printf("The %d one start. %s\n", k+1, FilePath+j)
			DownloadFile(j, FilePath+j, SecondPath, Year, Month, Name)
			k = k + 1
			time.Sleep(time.Second * 35)
		}
	}
	time.Sleep(time.Minute * 7)
	fmt.Println("Finished " + SecondPath + Year + Month + Name)
}
