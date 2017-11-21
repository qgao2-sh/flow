// Code generated by go-bindata.
// sources:
// model/model.swagger.json
// DO NOT EDIT!

package model

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

var _modelModelSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x5f\x6f\xdb\x38\x12\x7f\xcf\xa7\x18\xe8\xee\x61\x17\x68\x9b\xde\xe2\x70\x0f\x7d\xcb\x5e\xd3\x6d\x1f\x16\xdb\x6b\x72\xf7\x72\x57\x18\xb4\x38\xb6\xb8\xa1\x48\x85\xa4\xec\x0d\x8a\x7c\xf7\x03\x49\xfd\x37\x25\x4b\xb6\x63\x3b\xdb\x2c\xb0\xa8\x4d\x71\x46\xe4\xcc\x8f\xc3\xf9\xe7\x7c\xbb\x00\x88\xf4\x9a\x2c\x97\xa8\xa2\x77\x10\xfd\xf4\xe6\x6d\xf4\xca\x8e\x31\xb1\x90\xd1\x3b\xb0\xcf\x01\x22\xc3\x0c\x47\xfb\x3c\x95\x14\xf9\x9b\x4c\x49\x23\xdd\x3c\x80\x68\x85\x4a\x33\x29\xec\xd3\xe2\x23\x08\x69\x40\xa3\x89\x2e\x00\x1e\x1d\x37\x1d\x27\x98\xa2\x8e\xde\xc1\x7f\x3d\x51\x62\x4c\x56\x32\xb0\x9f\xb5\x9d\xfb\xd5\xcd\x8d\xa5\xd0\x79\x6b\x32\xc9\x32\xce\x62\x62\x98\x14\x97\xbf\x6b\x29\xea\xb9\x99\x92\x34\x8f\x47\xce\x25\x26\xd1\xf5\x96\x2e\x57\x7f\xbb\x5c\x70\xb9\xbe\x8c\x15\x12\x83\xd5\x03\x3b\x53\x6a\xd3\xf8\x0e\x10\xc9\x0c\x95\xe3\xf9\x89\xda\x7d\xfe\xd3\x91\xfc\xa2\x48\x96\x14\x9b\x70\xb3\x14\xea\x4c\x0a\x8d\xba\x45\x0c\x10\xfd\xf4\xf6\x6d\x67\x08\x20\xa2\xa8\x63\xc5\x32\x53\xc8\xae\xc1\xc8\x3d\x76\x22\x23\x1b\x64\x00\xd1\x5f\x15\x2e\x2c\xc5\x5f\x2e\x29\x2e\x98\x60\x96\x83\xbe\x74\x9a\x69\x2c\xec\x4b\xb1\x98\xa8\x45\xfe\x78\x11\xfa\xfc\xd8\xd8\x44\x46\x14\x49\xd1\xa0\xaa\x85\xea\xff\xeb\x2c\x5f\x90\xd4\x21\x62\x2e\xe9\x43\x77\xed\x4c\xf4\x3d\x51\x78\x9f\x33\x85\x56\x8a\x46\xe5\x78\xf0\x3d\xdf\xe7\xa8\xcd\x98\x2d\x7f\x6d\x6c\xd9\x90\x65\x77\xb3\xd1\x07\x2e\xd7\x37\xa8\x56\x2c\x6e\x48\xf0\xeb\x45\x93\x4d\x21\xb5\x1a\x49\xda\x28\x24\x69\x13\x49\x4b\x1c\x06\xd2\x8d\xa3\xb8\x5e\xa1\x30\xfa\x60\x48\xfa\xc1\xaf\x83\x89\x25\x54\x7c\x7e\x3c\x08\xbc\x9c\x90\xdd\x6a\x9f\x1a\x55\x4b\xfb\xa6\x37\x56\xaa\x33\x46\xc3\xf0\xba\xcf\x51\x0d\xe1\x6b\x41\xb8\xee\x02\xcc\x3c\x64\x8e\xbd\x36\x8a\x89\x65\x73\x13\x8f\xaf\x46\x2f\x4a\xc9\x74\xa6\xf1\xfe\xc9\x56\xd5\x79\xba\x90\x2a\x25\x16\x46\x51\xce\x84\xf9\xc7\xdf\xa3\xa3\xe1\xf9\x5b\x21\xff\xc7\x29\x90\xfe\x05\x8d\x83\xc9\x8d\xb1\x56\xf5\xac\xac\x63\x6b\x69\xc7\xb2\x8f\x83\x18\xb6\x77\xd2\x34\x13\x39\x80\xe0\xe3\x60\xe1\x32\x96\x69\xca\xcc\xb4\xfb\xd2\x93\x9c\x17\x18\x1a\x17\xc6\x67\x25\x63\xd4\x1a\xe9\x9f\x0f\x14\x23\xcc\xda\xe9\x6f\x70\x07\x8f\x73\xb8\xc1\x6b\x94\x53\xe4\xe4\x61\x12\xc8\xaf\x28\x7d\xef\x88\xce\x0a\xe6\x57\x94\xde\x18\xb2\xfc\x13\x9a\xbb\xe7\x80\xec\x12\x13\x85\x0a\xce\x04\xdb\x4c\xac\xe4\xdd\xb4\x88\xe7\x8a\xd2\x4f\x8e\xea\x43\x2e\x62\x07\xd0\x17\x94\xbf\xa0\xbc\x96\x7e\x1b\x1c\xe7\x05\x77\x6d\x57\x33\x15\xed\x6e\x0b\x2f\x20\x7f\x01\xf9\xa6\xf4\xcf\x09\xd6\xfa\xf2\x9b\xfb\xd7\x8d\x90\x35\x69\x3b\xe6\xdb\x62\xb5\x2b\x4b\x50\x62\x2a\xe7\x67\xe6\xa1\x77\x57\xf7\x5d\xe2\xbe\xd4\xee\x79\xac\xc6\xb0\x14\x65\x6e\x66\xa9\xfe\x8e\xb2\x1f\x81\x93\x16\xcb\x34\xe3\x38\x35\x6b\x5c\x10\x39\x48\x5f\xff\x61\x50\x09\xc2\xf9\x99\xc5\x0b\x3d\x8b\x7c\x39\x7b\x27\x5f\xcd\xc9\x6f\xc0\x5e\x68\x9c\xc9\x85\xb8\x22\x3c\x9f\xec\xe7\xfd\xc7\x12\xbd\x38\x7b\xe7\x75\xf0\x4e\x0e\xf5\x2b\x4a\x4b\xb4\x37\x10\x72\x64\xa4\x57\xd5\xd3\xc6\xf2\xea\xfa\x65\x30\x99\x5d\x2c\x33\x53\xa8\x51\x18\x52\x20\xb2\x3a\x10\xa5\x56\xe4\xfc\x77\x8c\x6b\x6f\x2f\xca\x94\x3d\x18\x86\x75\xf0\x5e\xce\x6f\x9d\x80\x3e\xcd\x36\xc1\xa7\x0d\x31\xf9\xc6\xd9\x19\x43\x49\x31\x43\x41\x51\xc4\xdd\xb5\x34\xe8\x89\x52\xa4\x0d\x82\x88\x19\x4c\xbb\xf3\x47\x26\xcb\x1f\x83\xe6\x65\x04\x0a\xf6\x90\x6b\x79\xd8\x76\x10\x50\xd7\xc8\xc1\x38\xa3\xcd\xa4\x28\x5c\xfc\x20\xd7\x58\x52\x9c\x71\x19\x77\x11\x33\x7a\x59\x31\xe1\x1c\xd5\x94\x3d\x5d\x74\x98\xd4\x3d\x06\x83\x62\x07\x5f\xae\xd7\x40\x40\xe0\x1a\xdc\x15\x0d\x6b\x66\x12\x20\xa0\x33\x8c\xd9\x82\xc5\xe0\x85\x14\x54\xe8\x66\x3a\xee\x34\x7a\x74\xf9\xe5\xd9\x06\x68\x87\xbc\xe1\x86\x2f\xdc\x71\x85\xcf\x59\x95\x1b\x02\x6f\x68\xd0\x09\xc1\xeb\x30\xac\xad\x81\xb4\xd2\x69\xd4\xb6\x28\x96\xb2\x23\x39\x51\xcb\x09\x67\xf7\xe3\xed\xed\xe7\x2f\x78\xff\x9e\x98\x3c\x7d\x1e\xca\xee\xd7\x17\x10\x4a\xad\xca\x4b\x01\x02\x13\xab\x72\xdd\x41\xd5\x9f\x81\xb2\x2b\x6f\x71\x27\x7b\xfb\x5b\x45\x1d\x96\x3c\x97\x3a\x57\x53\x4c\xf9\xcf\x5c\xce\x07\xb0\x40\x31\x3b\xd6\x95\x79\xc6\x08\x0c\x61\xae\x73\x51\x94\x3e\x86\xf7\x32\xc0\x48\x30\x09\x82\x6b\xee\x18\x84\x62\xe1\x88\x9f\x06\x8b\x55\x2c\x7a\x48\x29\xf9\x1d\x81\x42\x93\x2b\xa1\x9d\x18\xbc\x9c\x3e\xbd\x07\xb9\x70\xdf\xbd\xb1\xa6\xfd\x46\xba\x2f\x49\xf7\x4c\xc4\xd4\x0e\xfa\xac\x83\x74\x10\xdf\xaa\xdf\xa9\xac\x4f\xf1\x1e\x12\x9a\x73\x39\xdf\x71\x97\xb1\x14\x06\x85\x99\xed\xea\xdb\x73\x14\x4b\x93\xec\xe6\xb6\x74\x53\x78\x03\x20\xad\xc4\x04\x89\xe4\xee\x14\x2b\x5c\xa0\x42\x11\xa3\x3d\xb1\x04\xac\x08\x1c\x42\x89\xd6\x32\x66\x0e\xa4\x6e\x4c\x1b\xa9\x42\x48\x0d\x74\x19\x1c\x11\xa3\xfd\x80\xd8\x92\x56\xf9\xbe\xcf\xd1\xb9\xde\x34\xc3\x4a\x83\x94\xa8\x3b\x0d\x44\x00\xfe\xc1\xb4\x61\x62\x59\xd8\x55\xa2\xa1\x4c\x1d\x8f\x0e\x5a\xb6\x65\x64\x9f\x1f\x40\x74\x1e\xc7\xa8\xf5\x22\xe7\x7d\xd4\x73\x29\x39\x12\xd1\x67\x47\xca\xc7\x13\x8e\x57\xcb\x21\x0b\x88\xac\x63\xb8\x22\x14\xce\x42\xd7\x49\x9b\x28\x17\x77\x42\xae\xc5\xac\xf6\x0a\x9b\x8e\x7d\x1c\x63\x66\xae\x99\x49\x50\xb5\xc6\xb3\x8c\x3f\xdc\xca\xcd\x07\x26\x41\x71\xe5\x88\x7e\x96\xad\xe4\x9a\x7f\x62\xc9\xba\x83\x5f\x72\x11\xe6\xd0\x1d\xb5\x5b\x96\x1a\x03\xc3\x73\x26\x5a\xc3\xeb\x62\xb6\xab\x65\x34\xc6\x13\x22\x28\x6f\x8d\xe8\xbc\xbb\x24\xd6\xdb\x0f\x12\xb7\xa2\xf7\xa8\x1b\xf5\xb6\xe4\xc3\xf9\x6f\x8b\xd6\x80\x78\x68\x0f\x60\x81\xf8\x5a\x8f\xed\xa7\x56\x00\x4c\x6e\x54\x52\x0c\xaa\x94\x09\xa7\xa7\x8f\x52\xde\xf5\xd2\x94\xb2\x2a\x1e\x57\x79\xbb\x88\xe2\x82\x78\x33\x36\xa0\xf9\x0d\x83\xd0\x84\x19\xf8\xbc\xf0\x1c\xbd\x67\x65\x91\x06\x44\x50\x98\x63\x42\x56\x4c\xe6\xca\xfa\x58\xa4\xb0\x0c\x85\xbb\xd5\xe7\x85\x6e\x98\xc8\x3d\x4e\xfd\x13\x1c\xbf\xa6\x8a\x3b\xbe\x0d\x6c\x31\xfe\xdd\x88\x66\xab\xcd\x65\xa2\x14\x42\xe5\x1a\x38\xd3\x09\xde\x5d\xb0\x12\xb6\xce\x82\x9b\xe0\x13\x91\xf0\x43\xbd\xe5\xcb\x05\x61\x1c\xe9\x8f\x21\x19\x6f\xfe\x90\x60\x1f\xdb\xba\x5f\xa6\xe0\x80\xfe\x45\xe0\x37\x21\x67\xe1\xf8\xec\xed\x05\x63\x9a\x99\x87\x09\x50\xbb\xb6\xf3\x07\x22\x68\xeb\x3e\x1e\x2e\x1e\x47\xa5\xa4\x9a\xb2\x3a\x3b\x7f\x80\x9f\xbf\x71\x3d\xfd\x58\x9e\x45\x60\xb6\x18\x60\x9b\x18\x93\xcd\x14\xde\x1f\x34\x2f\x55\x30\xd5\xd9\x64\xae\x3a\x1b\x16\x81\x99\x92\x31\x71\xb5\x91\xf1\x06\xc6\x1b\x10\x14\x31\xc9\x72\x9d\x73\x9f\xa4\xe4\x1c\x32\xa9\x35\x9b\x73\xf4\x76\xc6\x5a\x73\x62\x5d\xbc\x07\x98\xb7\xa2\x8f\xd2\xa1\x73\xf6\x9c\x09\x67\xcf\xed\x69\xe9\x35\xea\x3e\x35\x1a\x27\x48\x73\x8e\xd4\xff\xae\x66\x1f\xbb\xbe\x87\x4b\x66\x58\x8a\x4f\x91\x8d\x36\x3b\xb2\xa4\xc4\xe0\x6b\xbb\xa8\x71\x19\x8d\x22\x93\x9c\x10\x6d\xc5\xaf\xac\x36\x5e\x83\x49\x98\x2e\x3c\x71\xab\x0a\x85\x9c\x18\xb6\x42\x3f\xc3\xdd\xbd\x02\xd0\x0a\x1d\xac\x1b\x54\xa5\xa3\x85\xa4\x08\x4c\x83\xc2\x58\xae\x50\x21\x0d\x68\xae\x61\x49\x46\x28\xac\x5a\x67\x4d\x06\xaa\x2c\xd2\xb9\x04\x55\xce\x39\x48\x05\xce\xa0\xf5\x06\x02\x0d\x03\x71\xe0\x9a\xde\x28\xab\x74\x6b\x09\x83\x4a\x4e\x51\xeb\x76\x8b\xe5\x30\xee\xfa\x15\x59\xbf\xad\x25\x20\x01\xce\x98\xfa\x83\x67\x95\x2c\xe3\x38\x57\xc5\x91\x63\x56\x75\xf5\x39\x7b\x55\x3b\x5b\x4c\x50\x16\xbb\x53\xec\x32\x59\x24\xd7\x95\x9f\xe5\xf9\x59\x5f\xc1\x7e\x2b\x36\x60\xb5\xce\x84\x07\xa1\x73\x0f\x07\x95\x70\xdb\x16\xe4\x0e\x91\x84\xbf\x21\x36\x4c\x7c\xd1\x64\xb5\xf9\xc0\xbb\x2e\x51\xa8\x1c\x11\xa0\xa9\x1d\x10\xe7\xa8\x07\xa8\x3d\x57\x2e\xb5\xe9\xf8\xf5\x84\x33\x3a\x2b\xef\x9b\x76\xff\xc1\xa0\x83\xdc\xde\x50\x48\xad\x56\x68\xe0\x40\x86\x4e\xaf\x4c\x78\xef\xbe\x50\x88\x8d\xf1\x61\x89\xc2\x7a\xd0\x5d\xf5\x06\x94\xf1\x81\x10\xfd\xa9\xaa\x23\x54\x45\xc3\x93\x5a\xd2\x27\xca\x7e\x1c\xd2\x90\x76\x72\x22\x07\xca\x2a\x83\xe5\x55\x66\xd3\x3f\x10\x72\x53\xa5\x3a\x42\x16\xb4\xad\xba\x1b\x6f\xb4\x4f\x7b\x05\x3e\x8d\x84\x77\x89\x03\x26\x49\xb9\x71\xe9\x05\xe4\x1c\xfe\xcd\xe3\xbe\x22\xee\x15\x55\x87\x8b\x7b\x46\x28\x65\xde\x9e\x7e\x0e\xf3\x84\xfe\x13\x32\xba\xcb\xa5\xc1\x2c\x5c\xa8\x3a\x9b\x78\xcc\xed\xe7\x70\xb6\xea\x74\xc5\xec\xe3\xf8\x76\x4e\x5c\xf6\x66\xae\x12\xa7\xee\xda\x5e\x33\xce\x41\x48\xe0\x52\x2c\x51\x41\x9c\x10\x11\xac\x4f\x79\x69\xfb\x22\xd6\x8b\xac\x47\xc9\xba\x28\xf9\xf5\x09\x73\x7f\x23\xbd\x19\x6c\xee\x58\x3e\xda\xe3\x7c\x36\x48\x9d\xdb\x38\x2b\x77\x3d\xfe\xe2\xde\x84\xd6\x00\x7b\x97\x8f\x24\xc2\xd8\x55\x4c\x7c\xc5\x6d\x99\xca\x14\xcb\xad\xaf\xa9\xaf\xdc\xa9\xfb\x68\x1b\xa4\xe0\x2b\x7c\xcf\x92\x2e\x03\xd6\x29\x49\xbe\x40\xa4\x1b\x7c\x85\xbf\xc0\x09\xa5\x93\xd8\xbb\x0b\xe1\xca\x12\x6d\x65\xbd\x8b\x80\x1c\xfb\x31\x02\xaa\x5f\x21\xf5\x6e\x6f\xb0\x74\x03\x2f\x58\x10\xa2\x67\x75\xe7\xcc\xac\xbc\xf7\xc7\xbf\x69\xc0\xef\x1a\xf5\xc6\x5d\xe4\x37\xe8\xa6\x8f\xbe\x32\x7b\x7f\xab\x7f\x16\xc9\xcc\xf0\x49\x7d\x96\x77\xcd\xee\x59\xb6\xfe\x14\xc1\x31\xef\xaf\x46\xe9\xa7\x93\x11\x12\x12\x52\xa9\xb0\xf0\x16\x34\xc4\x44\xc0\x1c\x21\x25\x14\xbd\x4b\xcd\xb4\x4f\x22\xfc\x4f\xb8\xcf\xce\xc9\x98\x23\x2c\x18\xe7\x72\x8d\x14\xe6\x0f\x40\x4a\x47\xc4\xb2\x6f\xa6\x8f\xc4\x43\xfb\xcd\x52\xde\x69\x48\xc8\x0a\x41\xe5\xa1\x1e\xb7\x8f\xb7\xb7\x9f\x3f\x22\xa1\xa8\xf6\x81\xc9\x1d\x6e\x64\xc0\x77\x6e\x23\xde\x21\x1e\xa9\x37\x01\x6b\x45\x32\x0d\x04\x34\x13\x4b\x8e\x90\xf8\xd1\x3b\x7c\xb8\xec\x4b\x68\x59\xe2\x5f\xd1\x24\x92\xee\x97\x47\x49\x3d\x8f\xe6\x65\x88\xad\x94\x86\x5d\x4b\xf3\x7b\xd6\x49\x79\x64\xed\xcc\x09\xc5\x6e\x3d\x54\x66\xbe\xfb\xbe\x49\x43\x4c\x9c\x8c\xca\x8a\x74\x96\xd7\x16\x9e\xdf\x7f\x9d\x12\xa9\xfb\x01\xec\x53\x28\x68\xc3\xb2\xab\x72\xf0\xfb\x74\x12\x49\x3a\xa5\x84\xb2\xa5\xe6\xe1\x95\x7e\x80\x36\xc4\x2d\x75\x82\xe2\xdc\x34\x68\xc2\x41\x5f\xda\x45\xd7\x18\xe6\xbf\x36\x64\x0e\xdb\xe1\x5f\x6a\xa1\x51\xee\x25\x02\x64\x6e\x96\x92\x89\x25\x48\x05\x4c\xc4\xd2\xfd\x39\x2e\xa7\xd2\xe2\x1e\x7b\x05\xcc\x14\x67\xc6\x86\xf2\x85\xe4\x5e\x15\x1a\x77\x31\x8e\x2c\xaa\xd4\xe0\x94\xd4\x87\x81\xb2\x62\xf2\x02\x82\x1e\x10\xf8\x22\xf0\x2c\x96\xb4\xd7\xe2\x31\x61\x70\xd9\x6a\xd0\x18\x2c\x70\x6c\x03\x44\xa1\x92\xd2\x24\x0e\xa2\xe1\xfa\x5f\xff\xbe\xbe\xb9\xed\x43\x43\x51\xc0\x76\x79\xd1\x12\x12\x42\xf6\x43\xa2\xeb\x09\x9f\x28\x97\x26\xa7\x14\xfe\x4e\xda\x32\x7d\xec\x5f\x1b\x1d\x21\xa5\x7b\x4e\x7d\x72\x57\x65\xe3\x35\xd1\xe0\xa2\xba\x6d\xed\xd6\xa1\x58\xeb\x25\x93\xbf\x4b\x49\xd4\x0b\xbe\x8a\xd7\x00\x5e\x03\x95\x6b\xe1\xff\x38\xa4\x7f\xaa\xcb\x2a\xb6\x51\x6c\xb9\xec\xa9\x75\x06\x62\xd3\x13\x29\xa4\x0c\xac\x67\xe7\x92\xe6\x1f\x87\xfb\x72\xd9\xc0\x84\x91\x00\xe5\xe2\xad\x42\xaa\xcf\x2e\xd2\xb0\x13\xb1\xee\x0d\x75\xed\xf8\x04\xd3\xb2\x99\xa9\xc1\xaa\x96\x40\x8f\xbe\xaa\x9e\x8f\xbd\x55\x15\x68\x38\xd9\xc1\x0e\xb4\x16\xd5\x2e\x7e\xb7\xfb\xba\x85\x34\x09\xaa\x76\x13\xc5\x80\xad\x28\xdb\x3b\x8e\x57\x10\xef\x8d\x76\x07\x37\x5f\x90\x54\x41\x52\xe3\x27\xe3\x65\x49\xda\x85\xdd\xad\x3e\x40\x70\xb5\x4f\x5a\xdf\x5f\x3d\xbb\xdf\xbf\x12\xad\xbb\x7f\xac\xd3\x35\xcc\x21\xed\x54\x9b\x37\x2a\xc8\x31\x11\x31\xf2\xce\xe0\x1d\x73\x23\xc5\xc0\x60\x68\xd4\x7e\x6f\x48\x62\x45\xc1\xb8\xd9\x48\x59\x09\xaa\xea\x9d\x54\x2d\x88\x5c\xd8\xff\x1f\x2f\xfe\x1f\x00\x00\xff\xff\xb1\x00\xae\xb6\xdb\x5a\x00\x00")

func modelModelSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_modelModelSwaggerJson,
		"model/model.swagger.json",
	)
}

func modelModelSwaggerJson() (*asset, error) {
	bytes, err := modelModelSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/model.swagger.json", size: 23259, mode: os.FileMode(420), modTime: time.Unix(1511277605, 0)}
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
	"model/model.swagger.json": modelModelSwaggerJson,
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
	"model": &bintree{nil, map[string]*bintree{
		"model.swagger.json": &bintree{modelModelSwaggerJson, map[string]*bintree{}},
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

