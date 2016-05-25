fileCommons contains following funcitons:


```
func MustDirExist(dirPath string) (err error)
func MustFileExist(filePath string) (err error)
func ReadToBytes(filePath string) ([]byte, error)
func ReadToString(filePath string) (content string, err error)
```