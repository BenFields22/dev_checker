package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
    fmt.Println("************************************") 

    fmt.Print("IP INFO") 
	out, err := exec.Command("curl", "https://ipecho.net/plain").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Print("\nPublic IP: ")
	output := string(out[:])
	fmt.Print(output+"\n")

	// let's try the pwd command herer
	out, err = exec.Command("ipconfig", "getifaddr", "en0").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Print("Private IP: ")
	output = string(out[:])
	fmt.Print(output)
    fmt.Print("------------------------------------\n") 
    fmt.Print("AWS CLI ACCOUNT\n") 
	out2, err2 := exec.Command("aws","sts","get-caller-identity").Output()

	// if there is an error with our execution
	// handle it here
	if err2 != nil {
		fmt.Printf("%s", err2)
	}

    output2 := string(out2[:])
	fmt.Print(output2)

    fmt.Print("------------------------------------\n") 
    fmt.Print("VERSION CHECKS\n")

    out3, err3 := exec.Command("go","version").Output()
    // if there is an error with our execution
	// handle it here
	if err3 != nil {
		fmt.Printf("%s", err3)
    }

    fmt.Print("GO:") 
    output2 = string(out3[:])
	fmt.Print(output2)

    out3, err3 = exec.Command("node","--version").Output()
    // if there is an error with our execution
	// handle it here
	if err3 != nil {
		fmt.Printf("%s", err3)
    }

    fmt.Print("NODE:") 
    output2 = string(out3[:])
	fmt.Print(output2)
    
    out3, err3 = exec.Command("npm","--version").Output()
    // if there is an error with our execution
	// handle it here
	if err3 != nil {
		fmt.Printf("%s", err3)
    }

    fmt.Print("NPM:") 
    output2 = string(out3[:])
	fmt.Print(output2)

    out3, err3 = exec.Command("python3","--version").Output()
    // if there is an error with our execution
	// handle it here
	if err3 != nil {
		fmt.Printf("%s", err3)
    }

    fmt.Print("PYTHON3:") 
    output2 = string(out3[:])
	fmt.Print(output2)


    out3, err3 = exec.Command("g++","--version").Output()
    // if there is an error with our execution
	// handle it here
	if err3 != nil {
		fmt.Printf("%s", err3)
    }

    fmt.Print("\nC++:\n") 
    output2 = string(out3[:])
	fmt.Print(output2)

    out3, err3 = exec.Command("swift","--version").Output()
    // if there is an error with our execution
	// handle it here
	if err3 != nil {
		fmt.Printf("%s", err3)
    }

    fmt.Print("\nSWIFT:\n") 
    output2 = string(out3[:])
	fmt.Print(output2)


    out3, err3 = exec.Command("java","--version").Output()
    // if there is an error with our execution
	// handle it here
	if err3 != nil {
		fmt.Printf("%s", err3)
    }

    fmt.Print("\nJava:\n") 
    output2 = string(out3[:])
	fmt.Print(output2)
    fmt.Print("************************************\n") 
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
