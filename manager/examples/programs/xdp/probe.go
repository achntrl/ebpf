// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ebpf/bin/probe.o

package main


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataProbeo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xb1\x6e\x13\x41\x10\x7d\x6b\x1b\xc7\x24\x29\x0c\x55\xb8\x10\xc9" +
	"\x0d\x88\x34\x0b\x49\x41\x45\x61\x45\x60\x1a\x17\x11\xa4\x40\xa2\x38\x1d\xe7\x0d\xb1\x62\x5f\xac\xbb\x0b\x18\x8c" +
	"\x44\x85\x44\x49\x83\x94\x0e\xf1\x15\x29\xf9\x15\x97\x48\x34\x50\x81\x40\x62\xd1\xac\x77\xed\xd5\xf8\x8c\x33\xd2" +
	"\x7a\x77\xde\xce\x9b\x37\x33\x5e\xdd\xdb\x07\xed\x56\x49\x08\x38\x13\xf8\x89\x99\x37\xb3\xf7\xa5\xd9\xb9\x69\x7f" +
	"\xd7\x21\x70\x6e\x83\xb3\xe0\xb7\xa6\x9d\xfc\xed\x55\xe0\x38\xf8\x65\xfc\x0d\x01\xa4\xaa\xd3\xa0\xf3\xad\x27\xf7" +
	"\xf7\x47\xc1\xf7\x29\xae\xf2\x46\x4c\xe7\x68\x90\x9f\x8e\x82\xaf\x53\x3c\x51\x2f\x4d\xfc\x20\x8a\x8f\x47\xc1\xd8" +
	"\xe0\x5f\x3e\x4f\x74\x56\x04\x30\xd6\x5a\x9f\x97\x80\x4d\x00\xef\x00\x54\x49\x17\x00\x95\xf8\x91\xd5\x4d\xb9\x28" +
	"\x0f\x69\x91\x0e\xd5\x42\x75\x6c\xaf\xe2\xe1\x7e\x1b\xf8\xab\xb5\xfe\xf4\x4d\x60\x83\xf1\x4c\x5b\xfe\x45\xc5\x5b" +
	"\x35\x00\xf5\x09\x5c\x77\xf1\xaf\x1f\xa1\xf6\x66\x4d\xac\x53\x0f\x76\x39\x3b\x2b\x98\x27\xb7\xb6\x49\xfd\x43\x17" +
	"\xdd\x95\x51\x9e\xc3\x6e\x00\xb8\x82\x4b\x53\xbf\x62\xf7\xeb\x06\xaf\xce\xe1\xf7\x00\x5c\xf5\xf2\x9c\x79\xf1\xae" +
	"\x5f\xf3\x17\xcb\x5c\x0d\x73\xc8\xbd\x83\x96\xa4\xc3\xb0\x33\xb8\xdd\x4d\x9e\xa7\x2a\xcb\x10\xbe\x50\x69\xd6\x3d" +
	"\x49\x10\xf6\xba\xb1\x4a\x32\x05\x99\xaa\x9e\x54\x47\xe1\x61\x1a\xf5\x15\x94\x09\x0b\xe3\x5e\x16\x1e\x9e\x26\x31" +
	"\xfa\x51\x37\x91\x31\x64\x96\xa7\x79\xf4\x0c\x32\x7b\xd5\x37\xfb\xde\x41\x0b\x32\x3d\xe9\x44\x79\x44\x77\x3b\x72" +
	"\xe7\xee\x05\x26\xf4\x7f\x7b\x6c\xe6\x34\x6f\x75\xfb\x6e\x87\x0c\xe7\x6f\x5c\xd8\x55\x65\x78\x73\x81\x5e\x85\xeb" +
	"\x2c\xe1\xf3\x37\x50\x63\xfe\x91\xe5\xef\x32\x7c\x6c\xf7\xcd\x02\x3d\xbf\x8f\x9b\xf6\xcc\x67\xf0\x67\x41\xbd\xbc" +
	"\xff\xad\x05\x7c\x17\xc8\xf9\xdc\x8f\x0b\x72\x9a\x38\x0b\x5e\x5b\xa2\xbf\xb2\x80\xbf\x65\xc1\xc6\x12\xfe\xae\xff" +
	"\x86\x3d\x6b\xda\xc0\x3b\x0c\xe7\xf3\x97\x00\x2e\x17\xe8\xbb\x84\x6e\xde\x6b\x36\xce\xf1\x1d\xfe\xb4\x40\x9b\x6c" +
	"\x60\xf5\x3f\x78\x75\x97\x3d\xbe\xfb\x4e\xfc\x0b\x00\x00\xff\xff\x99\xe3\xe2\xa9\x88\x05\x00\x00")

func bindataProbeoBytes() ([]byte, error) {
	return bindataRead(
		_bindataProbeo,
		"/probe.o",
	)
}



func bindataProbeo() (*asset, error) {
	bytes, err := bindataProbeoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "/probe.o",
		size: 1416,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1594295181, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"/probe.o": bindataProbeo,
}

//
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}


type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"": {Func: nil, Children: map[string]*bintree{
		"probe.o": {Func: bindataProbeo, Children: map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}