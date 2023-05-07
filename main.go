package main 

import ( 
	
		"os"
		"fmt" 
		"time"
		"path/filepath"
	)


func main() {
	
	//action()
	getMatchingFileNames()
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


func evaluateFileDuration( file string ) bool {

	//	iterate through the .mkv, .mp4 and .srt files in the directory
	//	if they are beyond 2 weeks, delete them

	//	get file info
	fileInfo, err := os.Stat(file)

	if (err != nil) {
		panic(err)
	}


	//	get current time
	currentTime := time.Now()

	//	get the duration of the file based on last modified date
	duration := currentTime.Sub(fileInfo.ModTime())

	//	initialize two weeks
	twoWeeks, err := time.ParseDuration("336h") //	2 weeks

	if (err != nil) {
		panic(err)
	}

	//	return true if the last modified status of the app is two weeks
	return duration > twoWeeks

	
}


func action() {

	file := getDownloadsDir() + "/Rabbit Hole S01E03 - The Algorithms of Control (NetNaija.com).mkv"


	isMoreThanTwoWeeks := evaluateFileDuration(file)

	fmt.Println(isMoreThanTwoWeeks)

}

func getMatchingFileNames() {

	folder := getDownloadsDir()
	
	extensions := []string{".mkv", ".mp4", ".srt"}




}

