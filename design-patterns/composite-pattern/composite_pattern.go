package composite_pattern

import "fmt"

/*
Composite to manage file and folders.
As we know Composite consists from following components:
- Component - main interface to which composite and its leaves must comply
- Leaf - Composite item. Implements Component.
- Composite - the one that manages the leaves. Also implements Component.
- Client - the user of the component.
Based on these definition we have interface FSItem which acts as a Component and implemented by leaves
FSFile and FSFolder. FSComposite also implements FSItem and contains []FSItem which can be either folder or file.
Using FSComposite we can use uniformly single FSItem the same way we use []FSItem as a group using single interface.
 */

type FSItemType int
const (
	FSItem_File FSItemType = iota
	FSItem_Folder
)

type FSItem interface {
	GetType() FSItemType
	GetPath() string
	Size() int
}

type FSFile struct {
	Parent FSItem
	FileName string
}

func (f *FSFile) GetPath() string {
	if f.Parent != nil {
		return fmt.Sprintf("%s/%s", f.Parent.GetPath(), f.FileName)
	} else {
		return f.FileName
	}
}

func (f *FSFile) GetType() FSItemType {
	return FSItem_File
}

func (f *FSFile) Size() int {
	return 1
}

type FSFolder struct {
	Parent FSItem
	FolderName string
	Items []FSItem
}

func (f *FSFolder) GetType() FSItemType {
	return FSItem_Folder
}

func (f *FSFolder) GetPath() string {
	if f.Parent != nil {
		return fmt.Sprintf("%s/%s", f.Parent.GetPath(), f.FolderName)
	} else {
		return f.FolderName
	}
}

func (f *FSFolder) Size() int {
	size := 0
	for _, item := range f.Items {
		size += item.Size()
	}

	return size
}

func (f *FSFolder) Add(item FSItem) {
	f.Items = append(f.Items, item)
}

type FSComposite struct {
	items []FSItem
}

func (f *FSComposite) Add(item FSItem) {
	f.items = append(f.items, item)
}

func (f *FSComposite) GetItems() []FSItem {
	return f.items
}

func (f *FSComposite) GetPath() string {
	return ""
}

func (f *FSComposite) Size() int {
	size := 0
	for _, item := range f.items {
		size += item.Size()
	}

	return size
}