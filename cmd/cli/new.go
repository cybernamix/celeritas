package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

var appURL string

func doNew(appName string) {
	appName = strings.ToLower(appName)
	appURL = appName

	// sanitise the application name - (convert URL to single word)
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		//extracts the last entry from the slice
		appName = exploded[(len(exploded) - 1)]
	}

	log.Println("App name is:", appName)

	//git clone the skeleton application
	color.Green("\tCloning a Repository...please wait")

	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "https://github.com/cybernamix/celeritas-app.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		exitGracefully(err)
	}

	// remove the .git directory
	color.Yellow("\tRemoving .git ... still working!")
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		exitGracefully(err)
	}

	// create a ready to go .env file
	color.Yellow("\tCreating .env file ... going well")
	data, err := templateFS.ReadFile("templates/env.txt")
	if err != nil {
		exitGracefully(err)
	}

	env := string(data)
	env = strings.ReplaceAll(env, "${APP_NAME}", appName)
	env = strings.ReplaceAll(env, "${KEY}", cel.RandomString(32))

	err = copyDataToFile([]byte(env), fmt.Sprintf("./%s/.env", appName))
	if err != nil {
		exitGracefully(err)
	}

	// create a makefile
	if runtime.GOOS == "windows" {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.windows", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()

		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
		source.Close()
	} else {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.mac", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()

		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
		source.Close()
	}

	color.Yellow("\tRemoving Makefile template...")
	err = os.Remove("./" + appName + "/Makefile.mac")
	if err != nil {
		exitGracefully(err)
	}

	err = os.Remove("./" + appName + "/Makefile.windows")
	if err != nil {
		exitGracefully(err)
	}

	// update the go.mod file
	color.Yellow("\tCreating go.mod file ...")
	_ = os.Remove("./" + appName + "/go.mod")

	data, err = templateFS.ReadFile("templates/go.mod.txt")
	if err != nil {
		exitGracefully(err)
	}

	mod := string(data)
	mod = strings.ReplaceAll(mod, "${APP_NAME}", appURL)

	err = copyDataToFile([]byte(mod), "./"+appName+"/go.mod")
	if err != nil {
		exitGracefully(err)
	}

	// update the existing .go files with correct names/imports

	color.Yellow("\tUpdating source files.....")
	// change to the directory where new app has been created
	os.Chdir("./" + appName)
	//call updateSource from helpers.go - walks through all files and makes appropriate changes
	updateSource()

	// run go mod tidy in project directory
	color.Yellow("\tRunning go mod tidy....")

	cmd := exec.Command("go", "mod", "tidy")

	err = cmd.Start()
	if err != nil {
		exitGracefully(err)
	}
	color.Green("\t........................................")
	color.Green("\tAnd... we are all done building - " + appURL)
	color.Green("\tGo forth and create something beautiful ;)")

}
