package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorings = 3
const delay = 5

func main() {

	showIntroduce()

	for {
		showMenu()
		readSitesInFile()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Command don't found!")
			os.Exit(1)
		}
	}

}

// sub-functions

// menu

func showIntroduce() {
	fmt.Println("Hello! You're welcome to website tester!")
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit program")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("")

	return commandRead
}

//options 

// option 1
func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readSitesInFile()

	for i := 0; i < monitorings; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func readSitesInFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
			fmt.Println(err)
			return sites
	}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			sites = append(sites, line)
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return sites
	}
	
	file.Close()
	
	return sites
}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "has been loaded successfully!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "with problems. Status code:", res.StatusCode)
		registerLog(site, false)
	}
}

// option 2
func registerLog(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

//option 3
func printLogs() {
	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}