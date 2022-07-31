package main

import (
	"fmt"
	"sort"
	"strings"
)

type inode struct {
	name     string
	isDir    bool
	children map[string]*inode
	content  []byte
}

type FileSystem struct {
	root         *inode
	inodeCounter int
}

func NewFileSystem() FileSystem {
	return FileSystem{
		root: &inode{
			name:     "/",
			isDir:    true,
			children: make(map[string]*inode),
		},
	}
}

func (fs *FileSystem) Ls(path string) []string {
	var nodes []string
	node := fs.root
	for _, dir := range fs.splitPath(path) {
		var ok bool
		node, ok = node.children[dir]
		if !ok {
			panic("not possible")
		}
	}
	if !node.isDir {
		return []string{node.name}
	}
	for _, child := range node.children {
		nodes = append(nodes, child.name)
	}
	sort.Strings(nodes)
	return nodes
}

func (fs *FileSystem) splitPath(path string) []string {
	if path == "/" {
		return nil
	}
	pathArr := strings.Split(path, "/")
	return pathArr[1:]
}

func (fs *FileSystem) Mkdir(path string) {
	fs.getOrCreateNode(path, true)
}

func (fs *FileSystem) createNode(name string, isDir bool) *inode {
	return &inode{
		name:     name,
		isDir:    isDir,
		children: make(map[string]*inode),
	}
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	node := fs.getOrCreateNode(filePath, false)
	node.content = append(node.content, []byte(content)...)
}

func (fs *FileSystem) getOrCreateNode(path string, isDir bool) *inode {
	node := fs.root
	for _, name := range fs.splitPath(path) {
		if _, ok := node.children[name]; !ok {
			node.children[name] = fs.createNode(name, isDir)
		}
		node = node.children[name]
	}
	return node
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	node := fs.getOrCreateNode(filePath, false)
	return string(node.content)
}

func main() {
	fs := NewFileSystem()
	fmt.Println(fs.Ls("/"))
	fs.Mkdir("/a/b/c")
	fs.AddContentToFile("/a/b/c/d", "hello")
	fmt.Println(fs.Ls("/"))
	fmt.Println(fs.Ls("/a/b/c/d"))
	fmt.Println(fs.ReadContentFromFile("/a/b/c/d"))
	fs.AddContentToFile("/a/b/c/d", "_world")
	fmt.Println(fs.ReadContentFromFile("/a/b/c/d"))
}
