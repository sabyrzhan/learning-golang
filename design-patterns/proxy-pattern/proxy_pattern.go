package proxy_pattern

import (
	"fmt"
	petname "github.com/dustinkirkland/golang-petname"
	"math/rand"
	"strings"
	"time"
)

/*
Here we have RemoteFileManager and CachedRemoteFileManager where both implement FileManager interface.
CachedRemoteFileManager is the proxy to RemoteFileManager.
*/

type FileType int
var txtFiles = map[string]int {"txt": 1, "yaml": 1, "yml": 1, "json": 1, "conf": 1, "properties": 1, "log": 1, "java": 1, "cpp": 1, "c": 1, "go": 1, "php": 1, "js": 1, "html": 1, "css": 1}
var audioFiles = map[string]int {"mp3": 1, "wav": 1, "ogg": 1, "m4a": 1, "wma": 1, "aac": 1, "flac": 1}
var videoFiles = map[string]int {"mp4": 1, "mkv": 1, "mpg": 1, "avi": 1, "mov": 1, "webm": 1, "wmv": 1, "mpeg": 1}
var allTypes = func () []string {
	var allExtensions []string
	var copyMap = func(data map[string]int) {
		for k,_ := range data {
			allExtensions = append(allExtensions, k)
		}
	}
	copyMap(txtFiles)
	copyMap(audioFiles)
	copyMap(videoFiles)

	return allExtensions
}()

const (
	FileTypeTxt FileType = iota
	FileTypeAudio
	FileTypeVideo
	FileTypeBinary

)

type File struct {
	CreateDate time.Time
	ModifiedDate time.Time
	FileType FileType
	FileTypeName string
	FileSize int
	FileName string
	Content []byte
}

func (f File) GetFileTypeDescription() string {
	if f.FileType == FileTypeTxt {
		return "Text file"
	} else if f.FileType == FileTypeAudio {
		return "Audio file"
	} else if f.FileType == FileTypeVideo {
		return "Video file"
	} else {
		return "Binary file"
	}
}

func (f File) String() string {
	var source string
	if strings.HasPrefix(f.FileName, "http") {
		source = "Remote file data"
	} else {
		source = "Cached file data"
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("========= %s ======\n", source))
	builder.WriteString(fmt.Sprintf("CreateDate: %s\n", f.CreateDate.UTC().Format("2006-01-02 15:04:05")))
	builder.WriteString(fmt.Sprintf("ModifiedDate: %s\n", f.ModifiedDate.UTC().Format("2006-01-02 15:04:05")))
	builder.WriteString(fmt.Sprintf("FileType: %s\n", f.FileTypeName))
	builder.WriteString(fmt.Sprintf("FileTypeDesc: %s\n", f.GetFileTypeDescription()))
	builder.WriteString(fmt.Sprintf("FileSize: %d bytes\n", f.FileSize))
	builder.WriteString(fmt.Sprintf("FileName: %s\n", f.FileName))
	builder.WriteString("==========================\n")

	return builder.String()
}


type FileManager interface {
	CreateFile(fileName string) string
	GetFile(fileName string) File
	DeleteFile(fileName string)
	ListFiles(dir string) []File
	UpdateFileContent(fileName string, content []byte)
}

func getFileNameExtension(fileName string) string {
	index := strings.LastIndex(fileName, ".")
	if index != -1 {
		ext := fileName[index+1:]
		return ext
	} else {
		return ""
	}
}

func getGuessFileTypeFromFileName(fileName string) FileType {
	ext := getFileNameExtension(fileName)
	if ext != "" {
		if _, ok := txtFiles[ext]; ok {
			return FileTypeTxt
		} else if _, ok := audioFiles[ext]; ok {
			return FileTypeAudio
		} else if _, ok := videoFiles[ext]; ok {
			return FileTypeVideo
		} else {
			return FileTypeBinary
		}
	} else {
		return FileTypeBinary
	}
}

func getRandomExtension() string {
	return allTypes[rand.Intn(len(allTypes))]
}

type RemoteFileManager struct {
	ApiUrl string
}

func (r RemoteFileManager) CreateFile(fileName string) string {
	fmt.Println("Create remote file: " + fileName)
	return fmt.Sprintf("%s/%s", r.ApiUrl, fileName)
}

func (r RemoteFileManager) GetFile(fileName string) File {
	file := File{
		CreateDate:   time.Now(),
		ModifiedDate: time.Now(),
		FileType: getGuessFileTypeFromFileName(fileName),
		FileTypeName: getFileNameExtension(fileName),
		FileSize: rand.Intn(1000000) + 1,
		FileName: r.ApiUrl + "/" + fileName,
		Content: []byte(fileName),
	}

	return file
}

func (r RemoteFileManager) DeleteFile(fileName string) {
	fmt.Println("Delete remote file at " + r.ApiUrl + "/" + fileName)
}

func (r RemoteFileManager) ListFiles(dir string) []File {
	var result []File
	for i := 0; i < 5; i++ {
		randomName := petname.Generate(2, "")
		result = append(result, r.GetFile(randomName + "." + getRandomExtension()))
	}

	return result
}

func (r RemoteFileManager) UpdateFileContent(fileName string, content []byte) {
	fmt.Println("Updated the content of the remote file: " + r.ApiUrl + "/" + fileName)
}

type CachedRemoteFileManager struct {
	fileManager FileManager
}

func (c CachedRemoteFileManager) CreateFile(fileName string) string {
	fmt.Println("Creating file on remote server...")
	result := c.fileManager.CreateFile(fileName)
	fmt.Println("Creation Done")
	fmt.Println("Added file to local cache")

	return result
}

func (c CachedRemoteFileManager) GetFile(fileName string) File {
	fmt.Println("Getting file from cache: " + fileName)
	result := c.fileManager.GetFile(fileName)
	result.FileName = "cached://" + fileName

	return result
}

func (c CachedRemoteFileManager) DeleteFile(fileName string) {
	fmt.Println("Deleting from cache done. Deleting on remote server...")
	c.DeleteFile(fileName)
}

func (c CachedRemoteFileManager) ListFiles(dir string) []File {
	fmt.Println("Listing files from local cache...")
	files := c.fileManager.ListFiles(dir)

	return files
}

func (c CachedRemoteFileManager) UpdateFileContent(fileName string, content []byte) {
	fmt.Println("Invalidating local cache file: " + fileName)
	c.fileManager.UpdateFileContent(fileName, content)
}

