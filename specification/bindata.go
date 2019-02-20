// Code generated by "esc -o bindata.go -pkg specification -ignore .*\.go -private -modtime=1518458244 schemas"; DO NOT EDIT.

package specification

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return []os.FileInfo(fis[0:limit]), nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/schemas/metadata_schema_v0.1.json": {
		name:    "metadata_schema_v0.1.json",
		local:   "schemas/metadata_schema_v0.1.json",
		size:    1916,
		modtime: 1518458244,
		compressed: `
H4sIAAAAAAAC/7RUwY7iMAy99yuqsEegrLQnfgUh5G1dCKJJxjFIaMS/j9pAaaZpWoG4OvbLs/38vpM0
TVPxx+YHrECsU3FgNussO1qtFi661LTPCoKSF6t/mYvNxNxVyqIuqpChAIade91dVsu/yxrikcZXg3Wi
/n/EnB9RQ9ogsUQr1qmj0sQVVOhFPAzLJNX+jtG+lpoq4KYDbblBaBNuz1xxQbJSq1H4YHGBNidpOAaw
8aLNS5Cx6/R8OgkvvA1+XPdjDeT4Gu8KpGKQCskOAwARXH+PVTJW/RonGsKyrptlBZZSyXoqNnt+5fd1
CxIzQKj446TcN4OEkg4tQfh1loSFt0onyYCMmsj2Xtr50hd0Zyi9Tt0FDQ5xHp6Ld0jPcYYPKn5Yo0oK
LK6twQrkaRRyE3yN30bkRtpbCVd1vMDR63cyWZrT9nXP/eQ2Rlvt215sb8OG8pYchtz1LdCYe00yjAnG
8aKrhQUVlZgzm+SW/AQAAP//m8slr3wHAAA=
`,
	},

	"/schemas/metadata_schema_v0.2.json": {
		name:    "metadata_schema_v0.2.json",
		local:   "schemas/metadata_schema_v0.2.json",
		size:    1732,
		modtime: 1518458244,
		compressed: `
H4sIAAAAAAAC/7RUzY7yMAy89ymq8B0LRZ/2xKsghLzUBSOSdB2DhFa8+6oNf9mmP2LFdZyxx/bE30ma
pqn65zY71KAWqdqJVIs83ztrph6dWd7mBUMp0/lH7rGJyjyTipqkUaAAgbWPrk/z2f9ZneL2TM4V1g/t
5x43ckMrthWyEDq1SL2UBjegMUCCHE6YzFbdg5fswTwhO7LmNXKBbsNUSV+CZYA2kWvKrB0xx8NBBfAq
WlgDGQEyyK5bOTDD+VcVRYK6zfE7ZSxr3iQvsCRDdVsuf5QKhV2iwipgNPJ2Ub5Mp6DkSZZi/DoSYxHs
wjsm4oMGWV2pTyVDvz0NpdWpN3jnELP4XAKfP8YZ93u/7wctHFncnYMa6DCYchmN9pu7x+R3s8dZqrSs
QepWvLx2J6OtOW5f17fv3MZgqzvrpMk4am9dd+xPZug7M6N+9ogf/uL5iW++1wv+KiSX5CcAAP//f4Kg
RsQGAAA=
`,
	},

	"/schemas": {
		name:  "schemas",
		local: `schemas`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"schemas": {
		_escData["/schemas/metadata_schema_v0.1.json"],
		_escData["/schemas/metadata_schema_v0.2.json"],
	},
}
