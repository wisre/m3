// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by "esc -prefix openapi/ -pkg openapi -ignore .go -o openapi/assets.go ."; DO NOT EDIT.

package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
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
	return nil, nil
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

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
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

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/asset-gen.sh": {
		local:   "asset-gen.sh",
		size:    221,
		modtime: 1527738292,
		compressed: `
H4sIAAAAAAAC/0zKQUrGMBDF8X1O8TrNqpAG18WFiBew7kSktpN0EGdKEkEQ7y5W0G82A7//67v4Ihrr
7tzNPN89PM/3t9f+yiUrEIiC/J9THCZs5gBAEjo8ImzwgqcJbWc9w8+tpk30nU9I4s7P626gzMplaaIZ
qdgbvNBvrSvCUTjJB8h/St8P8SsSwvGa/4EQJKsVxpgNwS6mS63c6piNQCO5zZTddwAAAP//XfBIdt0A
AAA=
`,
	},

	"/coordinator.yml": {
		local:   "openapi/coordinator.yml",
		size:    9182,
		modtime: 1527739323,
		compressed: `
H4sIAAAAAAAC/+xZXXObRhe+16/Yl7wXzYWFY7vpjO5wrKjMKI7H9mSmyfRiYQ9oU9hDdg9JnEz/e2dB
QkhQPqLESd1yYQv2Od8P5yxgPvA4Bj1jzsn02JlIFeFswhhJSmDGnBenF+fOhDEBJtQyI4lqxhyPCWlI
yyAnEIxkCsyAlmCY4MQDboDlRqqYvTi9vXnNogQ5PT1jIaaZBmMkqin7DXMWcsUiqQTDnFiKGhgP7E9r
lXFib1ZE2cx101MRTGNJqzyYSixO3d9/6lh8zFAzVOzNQtKveVBizcx117gQ0xJo/zye2gjfgzZldE+m
xzYVjIWoiIdk88GY4mmZkPMLtkCME2ALjXnmFKu5TmbMqazYBTONC1hhLEKdp+6j/5X/rVUrl8gQlIEd
A17GwxWwZbnETkpXGhZa4nCDBAM35YZAu0v/2fzyZu5MVmjICqKhwsIvJ8dPnImt0RWn1Yw5Ls+k+/6J
MyEem9nkaOOI/WcyHkKz/s9QRTLOdVnii3NWYY2zVZAlPIQUFA1QUMOacAUpFJ4U4TqTjNPK2CS5lZ0y
ZTGsi8NY6TxbH0d77tvD5GnK9d2MOQugHY/LdcxAc+ufL+rRL4A2iBCVyQvXKis8yxIZFmLuW4NqA800
ijwcBNVgMlQGau6fHB9vT/Zz59RWilzxOpax/2uIZsx55AqIpJJWyriXtXCu1wa3in7+6vYWoEDLcK41
6lJBZnnYUqzuUnlCMM4agL+plSfEt61VxjVPgUDXwGu2ByjutqmSqnGpmbvuSnlCXMO7HAz9QEw5uwem
3A8jt73E/Vz99C/+LFUJSIBgPF8vCrkRlC0Fvhtra5Hvkdc23e0lDe9yqUHMGOkcqst0l1ktdi+g4nul
aZm3Yo7otAj5/kl6dnz2UG6Gav52DtajvaHeHKu0AtaA7FK/Wn4Yk/WqFs4PN1k7qlVMVsWkMsRVCIxw
VPEewqi9qgXzPUZtN3UezqjtGacdJF2P0zHELEW8JHkAvaVryN3nUHAtZnZAs/GtDZ7IT+NqacUeTpex
0fzXZu6Fr583Y23Yhr6/A9UnZaQx/YKW9N2IvM3FP2uL/+/gbW3FyrY8/pcq13XA4C2Ea8Jl2hKO5LYM
RcW7S4dZZWvIW4CXJdypu/ayrmKQXwEiGdI8myseJCAaPgaICfCKzVGSm9VA7ActCcwtPsM0lbTEuE8g
tCf5UFc0ZFzqwWACZZPzclCWr/fgVftRPDMrpIFWpRLwcZhFvwa14tetDg+qaRXrFWiJ4pIrNA1PpSKI
oXY7RWh3MuXK07PN9SDB8I8b+QkO05JHEejnOeX6ayi64oYOj+qCE59/zKS+6yvjHtyLCPQlkheGYMyB
SfYbFBlUYxhGwMPL1/bWcRQXY2monuLurna9xu+Yvt5RMrjfll8RGkHXBe3BhSi84MlVQ82oNtx8KBjh
cLkN6ixo29gdYWHvTVL/dN+kaPMBbhR7Tk92XB7h52ZX9I0q56/V18aI3bc95yGh7otR5enNimvReytJ
U+D679AwJ3wP+la27A8GdjNpXki7Zek3lvKPhVs3QL7oMrdJ0piyiZ79jTSYFHdF8XW2B/wJVd9+6QPI
eEV9SQMlMpSKepSZ9qpyrXn9WZMg7WdYkeIdxb35tkf1Ibjb0wx1Z9CF9cPqtueYIU7Q0zRKWllgJYW5
DsHvY8Wa/weNc6sjir5Yxdb3nbTV/QSVp+XiEXP8S//W95b+a/9y4Wwueq88f+mdL+fVleXce7VGtLzP
/CoN8YvoWW+Aba9AfgzPRnXbttFT6+1W+7D+3qao/lQ4Zoe2xbeQ6q8AAAD//5v31qveIwAA
`,
	},

	"/doc.html": {
		local:   "openapi/doc.html",
		size:    560,
		modtime: 1527738292,
		compressed: `
H4sIAAAAAAAC/0ySQY8TMQyF7/wKk8te2smgRQKVpCAoxxVoxYVjmrgdL5l4FLtdVcB/R520S0+xv5c8
P1lxrzffvvz4+f0rDDrm9SvXDgA3YEjnAsApacb1w/3mMzzihqOzjTR1RA0Qh1AF1ZuD7pbvjb3VShjR
myPh88RVDUQuikW9eaakg094pIjLuVkAFVIKeSkxZPRvzMUoU/kFQ8WdN4PqJCtrd1xUuj3zPmOYSLrI
o40iH3dhpHzyD2cdaw26uu/7xdu+X7zr+z+PvGXlW2SgYvZG9JRRBkS9Dp1JqwG2nE7w+9IAjKHuqayg
//CCppASlf0N+9t87IuRs9e1urPfZU7FxBFkwrg81OzvEkexokEp2shcE5WgXLvTmO/Wzs7XrxFjpUlB
avy/mJhK9yQJMx1rV1Btmcb26FMOiqJ2eygpozTYiYaSQuaC3ZOY9Tnv7NoCt5zOto/xLwAA//9xV16p
MAIAAA==
`,
	},

	"/": {
		isDir: true,
		local: "openapi",
	},
}