package main 

import ( 
	
		"os"
		"fmt" 
	)


func main() {
	
	getDownloadsDir()
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

	downloadDirInfo, err := os.Stat(downloadDir)

	if (err != nil) {

		if os.IsNotExist(err) {
			fmt.Println("Does not exist")
		}
		
		panic(err)
	}

	fmt.Println(downloadDirInfo.IsDir())

	return ""
}

