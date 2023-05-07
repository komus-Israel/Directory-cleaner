package main 

import ( 
	
		"os"
		"fmt" 
		"time"
	)


func main() {
	
	getFileCreationTime()
}



func getDownloadsDir() string  {


	//	get the user home directory
	homeDir, err := os.UserHomeDir()

	
	//	initialize the downlaod path
	var downloadDir = homeDir + "/Downloads"

	//	create panic if there is an error
	if(err != nil) {
		panic(err)
	}

	_, err = os.Stat(downloadDir)

	if (err != nil) {

		if os.IsNotExist(err) {
			fmt.Println("Does not exist")
		}
		
		panic(err)
	}

	return downloadDir
}


func getFileCreationTime() {

	//	iterate through the .mkv, .mp4 and .srt files in the directory
	//	if they are beyond 2 weeks, delete them

	file := getDownloadsDir() + "/Rabbit Hole S01E03 - The Algorithms of Control (NetNaija.com).mkv"

	fileInfo, err := os.Stat(file)

	if (err != nil) {
		panic(err)
	}

	fmt.Println(fileInfo.ModTime())




}

