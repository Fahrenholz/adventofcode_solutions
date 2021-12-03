package main

import (
	"context"
	"fmt"
	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/firefox"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const instructionsURLTemplate = "https://adventofcode.com/2021/day/%d"
const inputsURLTemplate = "https://adventofcode.com/2021/day/%d/input"
const folderNameTmpl = "day%s"

var days = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twentyone", "twentytwo", "twentythree", "twentyfour", "twentyfive"}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Wrong Parameter count. Usage: 'go run . <daynum>'")
	}

	daynum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("could not parse daynum-parameter")
	}

	openBrowser(fmt.Sprintf(instructionsURLTemplate, daynum))
	createFolder(daynum)
	time.Sleep(2 * time.Second)
	getInputs(daynum)

}

func createFolder(daynum int) {
	daystr := days[daynum]
	_ = os.Mkdir(fmt.Sprintf(folderNameTmpl, daystr), 0775)
	_, _ = copyFile("main_template.txt", fmt.Sprintf(folderNameTmpl+"/%s", daystr, "main.go"))
}

func getInputs(daynum int) {
	cl, err := getClientWithCookieJar()

	rq, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf(inputsURLTemplate, daynum), nil)
	if err != nil {
		log.Fatal(err)
	}

	rp, err := cl.Do(rq)
	if err != nil {
		log.Fatal(err)
	}

	defer rp.Body.Close()

	bod, err := ioutil.ReadAll(rp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fmt.Sprintf(folderNameTmpl+"/%s", days[daynum], "inputs.txt"), bod, 0664)
	if err != nil {
		log.Fatal(err)
	}
}

func getClientWithCookieJar() (http.Client, error) {
	cookies := kooky.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`adventofcode.com`), kooky.Name("session"))
	if len(cookies) != 1 {
		log.Fatal("no cookie found")
	}

	jar, _ := cookiejar.New(nil)
	cookieUrl, err := url.Parse("https://adventofcode.com")
	if err != nil {
		log.Fatal(err)
	}

	cook := cookies[0].HTTPCookie()
	jar.SetCookies(cookieUrl, []*http.Cookie{&cook})
	cl := http.Client{Jar: jar}

	return cl, err
}

func openBrowser(url string) {
	var err error
	err = exec.Command("xdg-open", url).Start()

	if err != nil {
		log.Fatal(err)
	}
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
