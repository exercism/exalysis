// Code generated by go-bindata. DO NOT EDIT.
// sources:
// go_routines.md (130B)
// ifs_to_switch.md (75B)
// loop_rune_not_byte.md (152B)
// maprune.md (195B)
// move_map.md (148B)
// try_switch.md (92B)
// type_conversion.md (149B)
// unicode.md (123B)
// unicode_loop.md (99B)

package tpl

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
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
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _go_routinesMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcc\xbd\x0d\xc2\x30\x14\x45\xe1\x9e\x29\xee\x02\x64\x05\x1a\x1a\x46\xa0\x7c\x71\x2e\xb1\x15\xe2\x67\xbd\x1f\x45\xd9\x1e\x91\xfe\x9c\xef\x8e\x17\x9c\xc4\xa9\x09\x31\x22\xbd\xf5\x15\xab\xc2\x34\xa3\x75\xfa\x84\xb7\x26\x86\xe9\x98\x65\xfe\x9e\x38\xa4\x07\x17\x84\x62\x97\x8d\x68\x01\xcf\x41\xc3\x47\x3c\x26\x3c\xdb\x72\x51\xa5\xb2\x6c\x88\xca\xdb\xcc\x5e\xea\x2e\xb6\x39\xaa\x1e\xd8\xb3\xd4\xab\xa5\xfd\xdf\xe6\x8f\x5f\x00\x00\x00\xff\xff\x7b\x1b\x13\x88\x82\x00\x00\x00")

func go_routinesMdBytes() ([]byte, error) {
	return bindataRead(
		_go_routinesMd,
		"go_routines.md",
	)
}

func go_routinesMd() (*asset, error) {
	bytes, err := go_routinesMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go_routines.md", size: 130, mode: os.FileMode(420), modTime: time.Unix(1538237322, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1, 0xd1, 0xaa, 0xe9, 0x27, 0xe7, 0xbd, 0x1a, 0xf, 0x3, 0xc1, 0x10, 0x35, 0x32, 0xf0, 0x68, 0xfb, 0x85, 0x3a, 0x32, 0x11, 0x75, 0x33, 0xd5, 0x94, 0x72, 0x46, 0x33, 0x4b, 0x8b, 0xf4, 0x24}}
	return a, nil
}

var _ifs_to_switchMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\xc1\x09\x85\x30\x10\x04\xd0\x56\xa6\x81\xff\x9b\xb0\x91\x2c\x71\x64\x17\xc9\xe6\xb0\x13\xc4\xee\x05\x4f\xde\xdf\x0f\x9b\xb3\x9f\x98\x4b\x90\x13\xad\xae\x50\xf7\x86\x92\x89\x83\x29\x68\x62\x15\x11\x59\xa2\xed\xaf\x1a\x96\x37\x5a\x1c\x1f\x56\xff\x27\x00\x00\xff\xff\x20\x27\x3c\x25\x4b\x00\x00\x00")

func ifs_to_switchMdBytes() ([]byte, error) {
	return bindataRead(
		_ifs_to_switchMd,
		"ifs_to_switch.md",
	)
}

func ifs_to_switchMd() (*asset, error) {
	bytes, err := ifs_to_switchMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ifs_to_switch.md", size: 75, mode: os.FileMode(420), modTime: time.Unix(1538237055, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa4, 0x86, 0x72, 0xff, 0x80, 0xc8, 0xab, 0x94, 0x30, 0x68, 0x2b, 0xeb, 0x8a, 0x7c, 0x94, 0xc0, 0xd3, 0x8, 0xe6, 0xb8, 0x75, 0x6e, 0xb3, 0x39, 0x8c, 0xd7, 0x4d, 0x2f, 0x2, 0x9d, 0xc4, 0xc9}}
	return a, nil
}

var _loop_rune_not_byteMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcd\x41\xae\x83\x30\x0c\x84\xe1\xfd\x3b\xc5\xac\x23\xbd\xde\xa3\xfb\x1e\x20\xae\x71\xc1\x52\x70\x90\xed\x80\xb8\x7d\x15\x75\xff\xcd\xfc\xff\x78\xa6\x38\xa5\xda\x8a\x7e\x8a\x83\x50\x23\x5d\x6d\xad\xb8\xb4\x35\x1c\xde\x4f\x65\x41\xf5\x61\x52\x03\xd7\xa6\xbc\x41\x03\x04\xee\xfb\xd1\x24\x05\xbc\x91\x13\xe7\x5c\xdb\x02\x26\xc3\x1f\x77\x0b\x8d\x44\xff\xa0\x94\x7d\xb4\xd4\xa3\x49\x29\x78\xdf\x29\xf1\xc0\xcb\x6f\x8c\x98\xd1\x79\x1b\x50\x8b\x14\x5a\x26\xff\x89\x6f\x00\x00\x00\xff\xff\x9c\x11\x9d\x85\x98\x00\x00\x00")

func loop_rune_not_byteMdBytes() ([]byte, error) {
	return bindataRead(
		_loop_rune_not_byteMd,
		"loop_rune_not_byte.md",
	)
}

func loop_rune_not_byteMd() (*asset, error) {
	bytes, err := loop_rune_not_byteMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "loop_rune_not_byte.md", size: 152, mode: os.FileMode(420), modTime: time.Unix(1538236752, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x80, 0x86, 0x33, 0xbb, 0x93, 0xb2, 0x6a, 0x28, 0x25, 0x18, 0x75, 0x64, 0x3a, 0xe2, 0x56, 0x8b, 0x2, 0xc, 0xda, 0x34, 0x4d, 0x5d, 0xc8, 0x66, 0xe2, 0x7c, 0xbc, 0xe6, 0xeb, 0x34, 0xb6, 0x23}}
	return a, nil
}

var _mapruneMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\x41\x4e\x43\x31\x0c\x84\xe1\x3d\xa7\x98\x0d\xea\x8a\x1c\x80\x5d\xcf\x81\x90\x12\x92\x69\x6b\x1a\xec\x47\x6c\x17\xbd\xdb\xa3\x56\x5d\xcf\xa7\xd1\xff\x86\xdd\x12\xdd\x72\x0e\xa4\x13\x0d\xf5\xa7\x6d\x1f\x2b\x95\x9f\xa2\x51\x21\xea\xc1\x36\x60\x27\xd4\x57\xaf\x38\xd9\xc2\x90\xc5\x1e\x98\x66\xd7\xdc\xee\x4b\x5c\x08\xef\xb6\x88\x5b\x9b\xc9\x07\x6a\xa8\xf7\x97\xfa\x8e\x17\x35\xc4\xbe\x11\xdd\xf4\xc6\xe5\x62\x0a\x25\x07\x47\xc1\xf1\xa9\x20\x8e\xbe\xd8\x82\x03\x7f\x12\x17\xb8\xe8\x79\x12\xbf\x69\x41\x07\xcb\xb9\xe0\x70\x3c\xe0\x3b\x3d\x30\xe5\xfa\x28\xfd\xda\x83\xb5\xfc\x07\x00\x00\xff\xff\xf1\xe0\x83\x93\xc3\x00\x00\x00")

func mapruneMdBytes() ([]byte, error) {
	return bindataRead(
		_mapruneMd,
		"maprune.md",
	)
}

func mapruneMd() (*asset, error) {
	bytes, err := mapruneMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "maprune.md", size: 195, mode: os.FileMode(420), modTime: time.Unix(1538291623, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x28, 0xa0, 0xbc, 0xba, 0x9d, 0x42, 0xbc, 0x2, 0x70, 0xb2, 0xf4, 0xd2, 0x2d, 0xf9, 0xe3, 0x7a, 0x7b, 0x6d, 0xd2, 0xa3, 0x94, 0x26, 0xff, 0x3e, 0x9a, 0x29, 0xac, 0x26, 0x63, 0x5f, 0x9c, 0x15}}
	return a, nil
}

var _move_mapMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcd\x41\x0e\xc2\x30\x0c\x44\xd1\x3d\xa7\x98\x0b\xd0\xcb\xc0\x86\x5d\x2d\xd7\x34\x16\xd4\xae\x92\x09\xa8\xb7\x47\x41\x82\xfd\xfb\xfa\x67\xdc\xb2\x43\xaa\x41\xab\x09\x3d\x56\x08\x1a\x85\xae\xd8\x64\x87\x47\xf3\xc5\xc0\x62\x98\x2f\x9a\xd5\x66\xdc\x7b\x28\x3d\x63\xc2\xb5\x1e\xd8\xf2\x35\x22\x16\x6f\xd0\x5c\x0c\xd9\xf9\x4f\x7e\x14\xa7\x96\x70\x62\x98\xb1\xb1\x05\x19\xcf\x03\x19\x6a\x78\x3b\xcb\x57\xef\xa2\x0f\x59\x6d\xfa\x04\x00\x00\xff\xff\xdb\xe5\x4a\xc7\x94\x00\x00\x00")

func move_mapMdBytes() ([]byte, error) {
	return bindataRead(
		_move_mapMd,
		"move_map.md",
	)
}

func move_mapMd() (*asset, error) {
	bytes, err := move_mapMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "move_map.md", size: 148, mode: os.FileMode(420), modTime: time.Unix(1538248324, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb9, 0x1d, 0x8f, 0xd7, 0x76, 0xf7, 0x90, 0xb3, 0x24, 0x4f, 0x51, 0xd9, 0xdc, 0xf, 0x48, 0xb9, 0xaf, 0xb2, 0x32, 0xb5, 0xfd, 0x6e, 0x73, 0xda, 0xb0, 0xa6, 0x9d, 0xc1, 0x2, 0x18, 0xd6, 0x27}}
	return a, nil
}

var _try_switchMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xc8\xe1\x09\x03\x21\x0c\x05\xe0\x55\xde\x02\xdd\xc9\xe0\x25\xf6\x81\x8d\x62\x22\xe2\xf6\xe5\xfe\x7e\x1f\xd0\x70\xc7\x86\x2c\xc5\x9e\xb0\xb1\xc0\xc4\x0e\x7a\x83\xa0\xc4\x61\xd6\x6f\x01\x3d\x52\xe5\xc1\xb0\x57\x7f\x32\x0b\x0e\x7b\x07\xbd\x2e\x95\x50\xc4\x54\x7d\x10\x6c\x4e\x63\x15\xcf\x7e\xff\x01\x00\x00\xff\xff\xbf\xa6\x43\xdb\x5c\x00\x00\x00")

func try_switchMdBytes() ([]byte, error) {
	return bindataRead(
		_try_switchMd,
		"try_switch.md",
	)
}

func try_switchMd() (*asset, error) {
	bytes, err := try_switchMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "try_switch.md", size: 92, mode: os.FileMode(420), modTime: time.Unix(1538236086, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xeb, 0x78, 0xe3, 0x28, 0x74, 0x30, 0x3d, 0x2a, 0x43, 0xac, 0x2b, 0x3, 0x29, 0x3c, 0x6b, 0x2b, 0xdd, 0x5d, 0xf7, 0xf3, 0x14, 0xa0, 0xe9, 0xf6, 0xd1, 0x2d, 0x4b, 0x66, 0xf9, 0x64, 0x88, 0x3a}}
	return a, nil
}

var _type_conversionMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\x41\x0a\x84\x30\x0c\x46\xe1\xfd\x9c\xe2\xdf\xb9\x9a\xde\xc1\x63\xcc\xae\xa5\x8d\x1a\x46\x12\x49\x52\xc5\xdb\x4b\x71\xfb\xf8\x78\x5f\x84\xdd\x08\x45\x39\x95\x1b\xba\x08\x55\x72\x2f\x23\xde\x07\xa1\xaa\x9c\x64\xce\x2a\x9e\xf0\xd3\x8e\xaa\x7d\x6f\xb8\xd4\xfe\xb8\x38\xb6\x57\x65\xeb\x42\x19\x2c\x1e\x54\x1a\x74\x41\xf6\x30\x96\x75\x34\x7c\x62\x23\xe4\x45\x2d\x63\x57\x3d\x12\x66\x0c\x0f\x76\x54\xa3\x12\xd4\xde\x15\xa5\x35\x61\x9a\xa7\xf4\x04\x00\x00\xff\xff\x16\xf3\x99\x3b\x95\x00\x00\x00")

func type_conversionMdBytes() ([]byte, error) {
	return bindataRead(
		_type_conversionMd,
		"type_conversion.md",
	)
}

func type_conversionMd() (*asset, error) {
	bytes, err := type_conversionMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type_conversion.md", size: 149, mode: os.FileMode(420), modTime: time.Unix(1538236578, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x13, 0x10, 0xc, 0x79, 0x45, 0x8f, 0x4d, 0x57, 0x13, 0x9, 0xc2, 0x53, 0xbd, 0xb6, 0xb3, 0xbf, 0x38, 0x4, 0x32, 0x45, 0x43, 0xe9, 0xde, 0x4e, 0xea, 0x29, 0x6d, 0x60, 0xc7, 0x95, 0xdf, 0xfc}}
	return a, nil
}

var _unicodeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xca\x41\x0a\x02\x31\x0c\x05\xd0\xab\xfc\x8d\x4b\x05\xcf\x23\x42\x63\xf3\xeb\x04\x87\x66\x68\x52\xc1\xdb\xcb\x28\xb3\x7d\xbc\x33\x3e\x3e\x51\x7d\xae\x8a\x45\xde\x84\x60\x75\x7f\x41\x12\x65\x76\xab\xae\xbc\x9c\x6e\xd7\x7b\x14\x58\x0f\x53\x22\x17\xa2\xf9\xd8\xdb\xb6\x5b\x52\x14\xde\x50\x22\x87\xf5\x67\x1c\xfd\xc1\xe6\xe3\xdf\x7f\x35\x1d\xd6\xeb\xa0\x04\x11\x1b\xa9\xdf\x00\x00\x00\xff\xff\x61\x1a\x09\xa7\x7b\x00\x00\x00")

func unicodeMdBytes() ([]byte, error) {
	return bindataRead(
		_unicodeMd,
		"unicode.md",
	)
}

func unicodeMd() (*asset, error) {
	bytes, err := unicodeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "unicode.md", size: 123, mode: os.FileMode(420), modTime: time.Unix(1538240098, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd2, 0x96, 0x72, 0xf1, 0x76, 0x6b, 0x22, 0x6d, 0x8c, 0x97, 0xb0, 0xb1, 0x26, 0x3c, 0xbb, 0x27, 0x26, 0x83, 0x8d, 0x5a, 0xbe, 0xbc, 0x30, 0xb5, 0xcc, 0x21, 0x64, 0x9b, 0x50, 0x73, 0x26, 0x1d}}
	return a, nil
}

var _unicode_loopMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcb\x41\x0a\x42\x31\x0c\x45\xd1\xad\xbc\x89\x43\x05\xd7\x23\x42\x43\xfa\xf0\x07\x4b\x52\x9a\x54\x70\xf7\xf2\x07\x8e\xef\x3d\x57\x7c\x63\x43\x63\x8f\x8e\x43\x3e\x84\x60\x44\xbc\x21\x85\xb6\xdd\x34\x3a\x6f\x97\xc7\xfd\x99\x0d\x15\x58\x9c\x43\x94\x68\x59\xcb\xfc\x95\xff\x64\x8e\x3a\x78\xca\x79\x6e\xe6\xba\x28\x49\xe4\x24\xfb\x2f\x00\x00\xff\xff\xbc\xf2\x9d\xdb\x63\x00\x00\x00")

func unicode_loopMdBytes() ([]byte, error) {
	return bindataRead(
		_unicode_loopMd,
		"unicode_loop.md",
	)
}

func unicode_loopMd() (*asset, error) {
	bytes, err := unicode_loopMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "unicode_loop.md", size: 99, mode: os.FileMode(420), modTime: time.Unix(1538210688, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x52, 0x45, 0x9d, 0x18, 0x4e, 0x61, 0x5f, 0x1, 0x45, 0xcc, 0x1c, 0xa, 0xc9, 0xee, 0xad, 0xff, 0x36, 0x7, 0xd9, 0xa2, 0xdd, 0x42, 0xa3, 0x3b, 0x54, 0xf8, 0xdd, 0x95, 0x1, 0x7d, 0xec, 0x7c}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"go_routines.md": go_routinesMd,

	"ifs_to_switch.md": ifs_to_switchMd,

	"loop_rune_not_byte.md": loop_rune_not_byteMd,

	"maprune.md": mapruneMd,

	"move_map.md": move_mapMd,

	"try_switch.md": try_switchMd,

	"type_conversion.md": type_conversionMd,

	"unicode.md": unicodeMd,

	"unicode_loop.md": unicode_loopMd,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"go_routines.md":        &bintree{go_routinesMd, map[string]*bintree{}},
	"ifs_to_switch.md":      &bintree{ifs_to_switchMd, map[string]*bintree{}},
	"loop_rune_not_byte.md": &bintree{loop_rune_not_byteMd, map[string]*bintree{}},
	"maprune.md":            &bintree{mapruneMd, map[string]*bintree{}},
	"move_map.md":           &bintree{move_mapMd, map[string]*bintree{}},
	"try_switch.md":         &bintree{try_switchMd, map[string]*bintree{}},
	"type_conversion.md":    &bintree{type_conversionMd, map[string]*bintree{}},
	"unicode.md":            &bintree{unicodeMd, map[string]*bintree{}},
	"unicode_loop.md":       &bintree{unicode_loopMd, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
