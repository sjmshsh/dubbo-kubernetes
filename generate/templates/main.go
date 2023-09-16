/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const templatesPath = "../../templates"

// This program generates zz_filesystem_generated.go file containing byte array variable named TemplatesZip.
// The variable contains zip of "./templates" directory.
func main() {
	f, err := os.OpenFile("../zz_filesystem_generated.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	srcOut := bufio.NewWriter(f)
	defer srcOut.Flush()

	_, err = fmt.Fprintln(srcOut, `/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(srcOut, "// Code generated by go generate; DO NOT EDIT.\npackage generate\n\nvar TemplatesZip = []byte{")
	if err != nil {
		log.Fatal(err)
	}

	zipWriter := zip.NewWriter(newGoByteArrayWriter(srcOut))
	buff := make([]byte, 4*1024)
	err = filepath.Walk(templatesPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		name, err := filepath.Rel(templatesPath, path)
		if err != nil {
			return err
		}
		if name == "." {
			return nil
		}
		name = filepath.ToSlash(name)
		if info.IsDir() {
			name = name + "/"
		}

		header := &zip.FileHeader{
			Name:   name,
			Method: zip.Deflate,
		}

		// Coercing permission to 755 for directories/executables and to 644 for non-executable files.
		// This is needed to ensure reproducible builds on machines with different values of `umask`.
		var mode fs.FileMode
		switch {
		case info.Mode()&fs.ModeSymlink != 0:
			mode = 0o777 | fs.ModeSymlink
		case info.IsDir() || (info.Mode().Perm()&0o111) != 0: // dir or executable
			mode = 0o755
		case info.Mode()&fs.ModeType == 0: // regular file
			mode = 0o644
		default:
			return fmt.Errorf("unsupported file type: %s", info.Mode().String())
		}
		header.SetMode(mode)

		w, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		switch {
		case info.Mode()&fs.ModeSymlink != 0:
			symlinkTarget, err := os.Readlink(path)
			if err != nil {
				return err
			}
			_, err = w.Write([]byte(filepath.ToSlash(symlinkTarget)))
			return err
		case info.Mode()&fs.ModeType == 0: // regular file
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.CopyBuffer(w, f, buff)
			return err
		default:
			return nil
		}
	})
	zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprint(srcOut, "\n}\n")
	if err != nil {
		log.Fatal(err)
	}
}

// goByteArrayWriter dumps bytes as a Go integer hex literals separated by commas into underlying Writer.
// Each line of the output will be indented by a tab and each line will contain at most 32 integer literals.
// This is useful when generating Go array literals.
type goByteArrayWriter struct {
	i                 uint32
	w                 io.Writer
	hexDigitWithComma []byte
}

func newGoByteArrayWriter(w io.Writer) *goByteArrayWriter {
	return &goByteArrayWriter{
		i:                 0,
		w:                 w,
		hexDigitWithComma: []byte("0x00,"),
	}
}

var (
	hexs          = []byte("0123456789abcdef")
	space         = []byte(" ")
	newLineAndTab = []byte("\n\t")
)

const bytesInLine = 32

func (g *goByteArrayWriter) Write(bs []byte) (written int, err error) {
	for _, b := range bs {
		if g.i == 0 {
			_, err = g.w.Write(newLineAndTab)
			if err != nil {
				return
			}
		} else {
			_, err = g.w.Write(space)
			if err != nil {
				return
			}
		}

		g.hexDigitWithComma[2] = hexs[b>>4]
		g.hexDigitWithComma[3] = hexs[b&0x0f]
		_, err = g.w.Write(g.hexDigitWithComma)
		if err != nil {
			return
		}

		if g.i == bytesInLine-1 {
			g.i = 0
		} else {
			g.i++
		}

		written += 1
	}

	return
}
