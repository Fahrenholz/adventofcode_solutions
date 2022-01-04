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
	"runtime"
	"strconv"
	"time"
)

const instructionsURLTemplate = "https://adventofcode.com/%d/day/%d"
const inputsURLTemplate = "https://adventofcode.com/%d/day/%d/input"
const folderNameTmpl = "%d/day_%02d"

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Wrong Parameter count. Usage: 'go run . <yearnum> <daynum>'")
	}

	yearnum, err := strconv.Atoi(os.Args[1])

	daynum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("could not parse daynum-parameter")
	}

	openBrowser(fmt.Sprintf(instructionsURLTemplate, yearnum, daynum))
	createFolder(yearnum, daynum)
	time.Sleep(2 * time.Second)
	getInputs(yearnum, daynum)

}

func createFolder(yearnum, daynum int) {
	_ = os.Mkdir(fmt.Sprintf("%d", yearnum), 0775)
	_ = os.Mkdir(fmt.Sprintf(folderNameTmpl, yearnum, daynum), 0775)
	_, _ = copyFile("main_template.txt", fmt.Sprintf(folderNameTmpl+"/%s", yearnum, daynum, "main.go"))
}

func getInputs(yearnum, daynum int) {
	cl, err := getClientWithCookieJar()

	rq, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf(inputsURLTemplate, yearnum, daynum), nil)
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

	err = os.WriteFile(fmt.Sprintf(folderNameTmpl+"/%s", yearnum, daynum, "inputs.txt"), bod, 0664)
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

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
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
