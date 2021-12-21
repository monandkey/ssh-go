package config

import "github.com/monandkey/ssh/pkg/fileutil"

// fileExist is a function to call fileutil's FileExist.
func fileExist(fileName string) bool {
	return fileutil.FileExist(fileName)
}

// setHomedir is a function to call fileutil's GetHomedir.
func setHomedir() string {
	return fileutil.GetHomedir()
}

// setSeparate is a function to call fileutil's GetSeparate.
func setSeparate() string {
	return fileutil.GetSeparate()
}
