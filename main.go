package main

import (
	"os/exec"
	"os"
	"bufio"
	"log"
	"io"
	"fmt"
)


func main() {

    pwd, _ := os.Getwd()
    // sourceFolder := capture_origin_folder(nil)
    sourceFolder := string(pwd) + "/tzweb-app-repo"
    targetFolder := string(pwd) + "/tzweb-app-prod"

	welcome_message(sourceFolder, targetFolder)
	pull_repo(sourceFolder)
	clear_folders(targetFolder)
	sync_folders(sourceFolder, targetFolder)

    deploy_app(targetFolder)

}

func pull_repo(sourceFolder string){
    fmt.Print("pulling repo\n")
    runCommand("git", sourceFolder, "pull", "origin", "master")

}
func sync_folders(sourceFolder string, targetFolder string){

    fmt.Print("target folders sinchronized\n")
    copy_folder(sourceFolder + "/src", targetFolder + "/src")
    copy_folder(sourceFolder + "/routers", targetFolder + "/routers")

}

func clear_folders(targetFolder string){
    fmt.Print("target folders cleared\n")
    os.RemoveAll(targetFolder + "/src")
    os.RemoveAll(targetFolder + "/routers")


}
func deploy_app(folder string){
    fmt.Print("eb publishing\n")
    runCommand("eb", folder, "deploy", "tzweb-app-testing")
    // runCommand("eb", folder, "list")

}

func runCommand(command string, pwd string, args ...string) *exec.Cmd {

    cmd := exec.Command(command, args...)
    cmd.Dir = pwd

  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Fatal(err)
  }
  cmd.Start()

  buf := bufio.NewReader(stdout) // Notice that this is not in a loop
  num := 1
  for {
    line, _, _ := buf.ReadLine()
    if num > 20 {
	    return cmd;
      // os.Exit(0)
    }
    num += 1
	if string(line) == "" {

	} else {
	    fmt.Println(string(line))
	}

  }

    return cmd
}

func capture_origin_folder(defaultValue string) string {
	text := defaultValue
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: " + defaultValue)
	text, _ = reader.ReadString('\n')
	if text == "" {
		text = defaultValue
	}
	// fmt.Println(text)
	return text

}

func copy_folder(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = copy_folder(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copy_file(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func copy_file(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}


func welcome_message(sourceFolder string, targetFolder string){

	message := 
`
=============================================================================
  _____                        ____                _               
 |_   _|_ __  __ _  ____ ___  |  _ \   ___  _ __  | |  ___   _   _ 
   | | | '__|/ _' ||_  // _ \ | | | | / _ \| '_ \ | | / _ \ | | | |
   | | | |  | (_| | / /|  __/ | |_| ||  __/| |_) || || (_) || |_| |
   |_| |_|   \__,_|/___|\___| |____/  \___|| .__/ |_| \___/  \__, |
                                           |_|               |___/ 

=============================================================================
`

    fmt.Print(message)
    fmt.Print("\n")
    fmt.Print("\n")

    fmt.Print("Source: ")
    fmt.Print(string(sourceFolder))
    fmt.Print("\n")
    fmt.Print("Target: ")
    fmt.Print(string(targetFolder))
    fmt.Print("\n")
}



/*
https://docs.aws.amazon.com/sdk-for-go/api/aws/awserr/
https://github.com/aws/aws-sdk-go
https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_UpdateEnvironment.html
https://docs.aws.amazon.com/sdk-for-go/api/service/elasticbeanstalk/#ElasticBeanstalk.UpdateApplication
https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/ElasticBeanstalk.html#createApplicationVersion-property
https://golang.hotexamples.com/examples/github.com.bluet-deps.aws-sdk-go.service.elasticbeanstalk/-/New/golang-new-function-examples.html
https://www.google.com/search?ei=ECV9Xu_HKvDD5OUP75aJ0A0&q=nodejs+%22createApplicationVersion%22&oq=nodejs+%22createApplicationVersion%22&gs_l=psy-ab.3...11084.12677..12861...0.0..0.209.714.0j2j2......0....1..gws-wiz.iT9I7cgH1J8&ved=0ahUKEwjvj8j2kLnoAhXwIbkGHW9LAtoQ4dUDCAo&uact=5





*/


