package utils 



import ( 
	
	"os"
	"fmt" 
	"time"
	"path/filepath"
	
	
)

/**

	Function to get the path for the Downloads directory

*/

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


func evaluateFileDuration( fileInfo os.FileInfo ) bool {

	//	iterate through the .mkv, .mp4 and .srt files in the directory
	//	if they are beyond 2 weeks, delete them


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


func deleteDueFiles() {

	//	initialize the path to the Downloads path
	dir := getDownloadsDir()
	
	//	iterate through the files in the Downloads directory
	err := filepath.Walk(dir, func(path string, fileInfo os.FileInfo, err error) error {

		//	get the extension of the file path
		ext := filepath.Ext( path )

		//	if the extention is .mp4/ .mkv/ .srt, remove the files
		if (ext == ".mp4" || ext == ".mkv" || ext == ".srt") {

			
			isDueForDeletion := evaluateFileDuration( fileInfo )

			//	if it is due for deletion, then delete
			if (isDueForDeletion == true) {

				err := os.Remove(path)

				if (err != nil) {
					panic(err)
				}
			}		

		}
		

		if (err != nil) {
			return err
		}

		//	the iteration returns nil if there are no errors
		return nil
		
	})

	if (err != nil) {
		panic(err)
	}



}


func Execute()  {

	//	delete the due files
	deleteDueFiles()

	
}