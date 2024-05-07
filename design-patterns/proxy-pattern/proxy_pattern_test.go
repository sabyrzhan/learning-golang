package proxy_pattern

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	fileManager := RemoteFileManager{ApiUrl: "https://my-remote-cdn.com"}
	cachedFileManager := CachedRemoteFileManager{fileManager}
	cachedFileManager.CreateFile("test-file.txt")

	javaFile := cachedFileManager.GetFile("JavaSourceCode.java")
	javaFileRemote := fileManager.GetFile("JavaSourceCode.java")
	fmt.Println(javaFile)
	fmt.Println(javaFileRemote)

	audioFile := cachedFileManager.GetFile("AudioSourceCode.mp3")
	fmt.Println(audioFile)
	videoFile := cachedFileManager.GetFile("AudioSourceCode.avi")
	fmt.Println(videoFile)

	filesInDir := fileManager.ListFiles("animals-documentaries")
	fmt.Println(filesInDir)
}
