package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/gorilla/mux"
)

const (
	red    = "\u001b[91m"
	green  = "\u001b[32m"
	blue   = "\u001b[94m"
	yellow = "\u001b[33m"
	end    = "\u001b[0m"
	bold   = "\u001b[1m"
	underL = "\u001b[4m"
)

var args struct {
	Files             []string `arg:"positional" help:"files to host by nickname"`
	UseRandomFileName bool     `arg:"-r,--random" help:"generate random file name for hosted file"`
	Port              int      `arg:"-p,--port" help:"the port to listen on" default:"7000"`
}

type Config struct {
	Files []struct {
		Name string
		Path string
	}
}

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyz")
	files   = make(map[string]string)
	config  *Config
)

func readConfig() {
	usr, _ := user.Current()
	configFile := fmt.Sprintf("%s/.htf.json", usr.HomeDir)

	f, err := os.Open(configFile)
	if err != nil {
		fmt.Println("> failed to open configuration file: ~/.htf.yaml")
	}
	fileBytes, _ := ioutil.ReadAll(f)

	json.Unmarshal(fileBytes, &config)
}

func getFilePathFromNickname(nickname string) string {
	//var filePath string
	for _, entry := range config.Files {
		if entry.Name == nickname {
			return entry.Path
		}
	}
	return ""
}

func getRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func handleFileDownload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["file"]

	d, err := ioutil.ReadFile(files[f])
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	http.ServeContent(w, r, f, time.Now(), bytes.NewReader(d))
}

func main() {
	arg.MustParse(&args)

	readConfig()

	if len(args.Files) == 0 {
		fmt.Println("- must provide files to host from ~/.htf.yaml")
		return
	}

	for _, file := range args.Files {
		filePath := getFilePathFromNickname(file)
		if filePath == "" {
			fmt.Printf(
				"- could not find file path for %s ...\n- make sure it has an entry in ~/.htf.json ...\n",
				file,
			)
			return
		}
		if args.UseRandomFileName == true {
			r := getRandomString(6)
			files[r] = filePath
		} else {
			files[file] = filePath
		}
	}

	for file, path := range files {
		fmt.Println(
			fmt.Sprintf("Hosting %s%s%s%s %sat%s %shttp://IP:%d/f/%s%s",
				red, underL, filepath.Base(path), end, yellow, end, blue, args.Port, file, end),
		)
	}

	r := mux.NewRouter()
	r.HandleFunc("/f/{file}", handleFileDownload).Methods("GET")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", args.Port), r))
}
