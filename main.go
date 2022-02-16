package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func addToFile(name string, content string) {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(content + "\n"); err != nil {
		log.Println(err)
	}
}

func createFile(name string, content string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(content)
}
func createMakeFile(repoName string) {
	content := `##Create your Make file
	deps: # install dependencies
	data: # runs go generate and/or other things like protoc
	build: # builds the Go binary
	image: # builds Docker image
	lint: # runs go vet or other similar tools
	test: # runs tests
	package: # packages the service/app
	publish: # publishes the service/app`

	createFile(repoName+"/Makefile", content)
}

func createMainFile(repoName string) {
	content := `
				package main
	            import "fmt"
	
	            func main() {
	            	fmt.Println("goinit is awesome")
	            }`

	createFile(repoName+"/main.go", content)
}

func main() {

	println(color.Colorize(color.Blue, `
     _____           _____           _   _   
    / ____|         |_   _|         (_) | |  
   | |  __    ___     | |    _ __    _  | |_ 
   | | |_ |  / _ \    | |   | '_ \  | | | __|
   | |__| | | (_) |  _| |_  | | | | | | | |_ 
    \_____|  \___/  |_____| |_| |_| |_|  \__| goinitVTAG\n
	`))

	goinitV := flag.Bool("v", false, "goinit version")
	giName := flag.String("name", "", "repository name")
	giDesc := flag.String("desc", "", "repository description")
	giPublic := flag.Bool("public", false, "repository access (default to private)")

	flag.Parse()
	if *goinitV {
		return
	}

	token := goDotEnvVariable("GH_TOKEN")
	org := goDotEnvVariable("GH_ORG")
	folders := goDotEnvVariable("GO_FOLDERS")

	description := "New repository created by goinit"
	var repoName string
	var privateRepo string

	if *giName == "" {
		println(color.Colorize(color.Red, "Error creating a new repository: missing argument Repo Name"))
		for repoName == "" {
			println(color.Colorize(color.Blue, "Please Enter Repository Name: "))
			fmt.Scanln(&repoName)
		}
	} else {
		repoName = *giName
	}

	if *giDesc != "" {
		description = *giDesc
	}

	if *giPublic {
		privateRepo = "false"
	} else {
		privateRepo = "true"
	}

	println(color.Colorize(color.Yellow, "About to create the following repository:"))
	fmt.Printf("Org : %s \n", org)
	fmt.Printf("Repository: %s \n", repoName)
	fmt.Printf("Description: %s \n", description)
	fmt.Printf("Private: %s \n", privateRepo)

	var approve string
	println(color.Colorize(color.Blue, "If you wish to Cancel, press C or Enter to continue: "))
	fmt.Scanln(&approve)

	if strings.ToLower(approve) == "c" {
		os.Exit(1)
	}

	url := "https://api.github.com/user/repos"
	method := "POST"

	payloadbody := fmt.Sprintf(`{"name":"%s","description":"%s","private":"%s","auto_init":"true","gitignore_template":"Go"}`, repoName, description, privateRepo)
	payload := strings.NewReader(payloadbody)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "token "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	//fmt.Printf("%+v\n", res)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	if res.StatusCode == 200 || res.StatusCode == 201 {
		println(color.Colorize(color.Blue, "Successfully created a new repository\n"))
		fmt.Println(string(body))

		cmd := exec.Command("git", "clone", "https://"+org+":"+token+"@github.com/"+org+"/"+repoName+".git")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		foldersArr := strings.Split(folders, ",")

		for i, s := range foldersArr {
			fmt.Println(i, s)
			exec.Command("mkdir", repoName+"/"+string(s)).Run()
		}

		createMakeFile(repoName)
		createMainFile(repoName)
		addToFile(repoName+"/.gitignore", ".env\n/bin\ngo.mod\ngo.sum")

	} else {
		println(color.Colorize(color.Red, "Error creating a new repository"))
		fmt.Println("\nReason: " + res.Status + "\n")
		fmt.Println(string(body))
	}
}
