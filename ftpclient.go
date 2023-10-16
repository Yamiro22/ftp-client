package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jlaffaye/ftp"
)

func main() {
	// FTP server details
	const server = "ftp.example.com"
	const port = "21"
	const username = "your_username"
	const password = "your_password"
	const filename = "file_to_download.txt" // Name of the file you want to download

	// Connect to the FTP server
	connection, err := ftp.Dial(server + ":" + port)
	if err != nil {
		log.Fatal(err)
	}

	// Login using provided credentials
	err = connection.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	// List the files in the root directory
	entries, err := connection.List("/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files on the FTP server:")
	for _, entry := range entries {
		fmt.Println(" -", entry.Name)
	}

	// Download a file
	r, err := connection.Retr(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	// Save the file to the local system
	err = ioutil.WriteFile(filename, buf, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File downloaded successfully:", filename)

	// Logout and close the connection
	connection.Logout()
	connection.Quit()
}
