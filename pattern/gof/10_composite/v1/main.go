package main

import (
	"fmt"
	"os"
)

// FileSystemNode 接口定义了文件系统节点共有的行为
type FileSystemNode interface {
	CountNumOfFiles() int
	CountSizeOfFiles() int64
	GetPath() string
}

// File 结构体实现了 FileSystemNode 接口
type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{path: path}
}

func (f *File) CountNumOfFiles() int {
	return 1
}

func (f *File) CountSizeOfFiles() int64 {
	fileInfo, err := os.Stat(f.path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func (f *File) GetPath() string {
	return f.path
}

// Directory 结构体实现了 FileSystemNode 接口
type Directory struct {
	path     string
	subNodes []FileSystemNode
}

func NewDirectory(path string) *Directory {
	return &Directory{path: path, subNodes: []FileSystemNode{}}
}

func (d *Directory) CountNumOfFiles() int {
	numOfFiles := 0
	for _, subNode := range d.subNodes {
		numOfFiles += subNode.CountNumOfFiles()
	}
	return numOfFiles
}

func (d *Directory) CountSizeOfFiles() int64 {
	sizeOfFiles := int64(0)
	for _, subNode := range d.subNodes {
		sizeOfFiles += subNode.CountSizeOfFiles()
	}
	return sizeOfFiles
}

func (d *Directory) AddSubNode(subNode FileSystemNode) {
	d.subNodes = append(d.subNodes, subNode)
}

func (d *Directory) RemoveSubNode(removedNode FileSystemNode) {
	for i, subNode := range d.subNodes {
		if subNode.GetPath() == removedNode.GetPath() {
			d.subNodes = append(d.subNodes[:i], d.subNodes[i+1:]...)
			break
		}
	}
}

func (d *Directory) GetPath() string {
	return d.path
}

// 示例
func main() {
	/**
	 * /
	 * /wz/
	 * /wz/a.txt
	 * /wz/b.txt
	 * /wz/movies/
	 * /wz/movies/c.avi
	 * /xzg/
	 * /xzg/docs/
	 * /xzg/docs/d.txt
	 */
	root := NewDirectory("/")
	dirWz := NewDirectory("/wz/")
	dirXzg := NewDirectory("/xzg/")
	root.AddSubNode(dirWz)
	root.AddSubNode(dirXzg)

	fileWzA := NewFile("/wz/a.txt")
	fileWzB := NewFile("/wz/b.txt")
	dirWzMovies := NewDirectory("/wz/movies/")
	dirWz.AddSubNode(fileWzA)
	dirWz.AddSubNode(fileWzB)
	dirWz.AddSubNode(dirWzMovies)

	fileWzMoviesC := NewFile("/wz/movies/c.avi")
	dirWzMovies.AddSubNode(fileWzMoviesC)

	dirXzgDocs := NewDirectory("/xzg/docs/")
	dirXzg.AddSubNode(dirXzgDocs)

	fileXzgDocsD := NewFile("/xzg/docs/d.txt")
	dirXzgDocs.AddSubNode(fileXzgDocsD)

	fmt.Println("/ files num:", root.CountNumOfFiles())
	fmt.Println("/wz/ files num:", dirWz.CountNumOfFiles())
}
