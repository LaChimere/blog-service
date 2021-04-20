// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xd1\x6e\xdb\x30\x0c\x7c\xf7\x57\x10\xd8\x73\x1d\xa5\x59\xd3\x42\x6f\x6d\x9a\xa2\x19\x9a\x2d\x98\x5d\xf4\x71\x60\x62\xc6\xd1\x20\x5b\x0a\x45\xa7\xee\xbe\x7e\x90\xdd\xc5\x4d\x33\x3d\x09\x77\x87\x3b\x1d\xc5\x8c\xf8\x40\xac\x13\x80\x9f\x4d\xbd\x74\x05\x69\x28\x68\xdd\x94\x09\xc0\xa3\x88\x5f\x39\x16\x0d\x37\x4a\xa9\xa8\x20\x2c\x72\x53\x91\x6b\x44\xc3\x34\x22\x2f\x6c\x84\x3e\x42\xb7\xde\x47\xaf\x7b\xda\x62\x63\x65\x85\x25\x65\xe6\x0f\x69\x18\x47\xf5\x12\xdb\x8f\x88\x1a\x84\x33\x57\x0b\xb5\x72\x6a\xfe\xe4\xca\x0c\x0f\xb4\x42\xd9\x69\x08\xe2\x18\x4b\x1a\x59\x57\x86\x9e\x7b\x30\x96\xbe\x63\x45\x1a\xd0\xfb\x01\x9a\xb7\xa2\x21\xb5\x2e\x36\x78\xf6\xd6\x61\x71\x6e\xd2\x74\x78\x18\x14\xdd\x10\x9e\xd9\x6a\xd8\x89\x78\x3d\x1a\x8d\x2f\xaf\x53\x95\xaa\x74\xac\x63\xf7\x51\x10\x14\xb3\x39\xea\x17\x15\x96\xb4\xc4\xb6\x6f\x72\x05\x5f\x60\x79\x77\x4a\xde\x5a\xeb\x5e\xe7\xad\x84\x38\x0d\x80\x0b\x48\x7f\xfb\x72\xb8\xd2\xf1\xee\xeb\x32\x99\x57\x68\x6c\x14\x3e\xba\x20\x1a\x42\x25\x3e\xdd\xef\xd3\x8d\xab\x12\x80\xfe\x07\xbe\x4e\xaf\x62\x40\x20\xae\xbb\xca\xed\xa7\x13\x85\x18\xc2\xab\xe3\xe2\x8c\x4c\x00\x16\x21\xcb\x9e\x34\x08\x37\x94\x00\x3c\xb0\xab\xfe\x6b\x91\xbb\x7f\xcf\xfd\x4c\x76\x82\x6f\x2f\x79\xe4\x33\xda\x30\x89\x06\x8b\x9b\x9d\xa9\x88\xa9\x0b\x08\x0d\xb1\x86\xb5\x75\xe5\x45\x20\x3e\x98\x4d\x84\xe7\xad\x37\x4c\x1a\xae\x2f\x95\x4a\xee\x51\x70\x8d\x81\xba\x05\xb9\xcb\xdf\x3c\x69\xa8\xde\xc2\xde\x9e\x34\x63\xe7\xe4\xac\xcd\x71\x36\xc3\xbf\x4c\x26\x6a\xda\x19\xf5\x3b\x10\x83\x7f\x0d\xc1\x39\xae\x2d\xad\x98\xb6\xa6\x7d\xe7\x12\x80\xd9\x0e\x39\xc4\x97\x37\xb2\xbd\xe9\x32\x38\x74\xdb\xab\x21\xef\x27\xb3\xc4\x76\x51\x58\x9a\xb9\xba\x0e\xc3\xd2\xfe\xf0\x54\xbf\x43\x13\x95\xfc\x0d\x00\x00\xff\xff\x99\xd4\x7c\x26\x32\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 818, mode: os.FileMode(436), modTime: time.Unix(1618886195, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"configs/config.yaml": configsConfigYaml,
}

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
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
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

var _bintree = &bintree{nil, map[string]*bintree{
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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
