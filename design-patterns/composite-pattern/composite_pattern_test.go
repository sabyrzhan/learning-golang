package composite_pattern

import (
	"fmt"
	"testing"
)

func TestFileSize(t *testing.T) {
	folder := &FSFolder{ nil, "my-photos", []FSItem{}}
	file := &FSFile{folder, "photo1.png"}
	folder.Add(file)
	folder.Add(&FSFile{folder, "photo2.png"})
	folder.Add(&FSFile{folder, "photo3.png"})
	folder2 := &FSFolder{ folder, "my-docs", []FSItem{}}
	folder2.Add(&FSFile{folder2, "document.png"})
	folder2.Add(&FSFile{ folder2,"documen2.png"})
	folder2.Add(&FSFile{folder2, "documen3.png"})
	folder.Add(folder2)

	fmt.Printf("Total folder size: %d\n", folder.Size())
	fmt.Printf("Folder1 size: %d\n", folder.Size())
	fmt.Printf("Folder2 size: %d\n", folder2.Size())
	fmt.Printf("File size %d\n", file.Size())
	fmt.Printf("Folder path %s\n", folder.GetPath())
	fmt.Printf("Folder first item path %s\n", folder.Items[0].GetPath())
	fmt.Printf("Folder2 path %s\n", folder2.GetPath())
	fmt.Printf("Folder2 first item path %s\n", folder2.Items[0].GetPath())
}