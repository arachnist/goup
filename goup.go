// Copyright 2016 Robert S. Gerus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

func main() {
	config := setup()

	filename := path.Base(os.Args[1])
	ext := path.Ext(filename)

	hash := sha1.New()
	temp, err := ioutil.TempFile("./", "goup_")
	if err != nil {
		log.Fatalln("Error creating temporary file:", err)
	}

	defer os.Remove(temp.Name())

	multi := io.MultiWriter(hash, temp)
	_, err = io.Copy(multi, os.Stdin)

	destFile := fmt.Sprintf("%x%s", hash.Sum(nil), ext)
	dest := path.Join(config.FilesDir, destFile)
	old := temp.Name()
	temp.Close()
	err = os.Rename(old, dest)
	if err != nil {
		log.Fatalln("Error renaming temporary file:", err)
	}

	filename = strconv.FormatInt(time.Now().UnixNano(), 16) + "." + filename
	nameLink := path.Join(config.NamesDir, filename)

	err = os.Symlink(dest, nameLink)
	if err != nil {
		log.Fatalln("Error creating name link:", err)
	}

	_ = os.Chmod(dest, 0644)
	_ = os.Chmod(nameLink, 0644)

	fmt.Println(config.FilesUrlBase + "/" + destFile)
	fmt.Println(config.NamesUrlBase + "/" + filename)
}
