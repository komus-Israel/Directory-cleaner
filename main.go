package main 

import ( 
	
		"os"
		"fmt" 
		"time"
	)


func main() {
	
	action()
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


	fileInfo, err := os.Stat(file)

	if (err != nil) {
		panic(err)
	}

	currentTime := time.Now()

	duration := currentTime.Sub(fileInfo.ModTime())

	twoWeeks, err := time.ParseDuration("336h") //	2 weeks

	if (err != nil) {
		panic(err)
	}

	return duration > twoWeeks

	
}


func action() {

	file := getDownloadsDir() + "/Rabbit Hole S01E03 - The Algorithms of Control (NetNaija.com).mkv"


	isMoreThanTwoWeeks := evaluateFileDuration(file)

	fmt.Println(isMoreThanTwoWeeks)

}

