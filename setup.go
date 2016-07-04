// Copyright 2016 Robert S. Gerus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"

	"gopkg.in/yaml.v2"
)

type goupConfig struct {
	FilesDir     string `yaml:"files_dir"`
	NamesDir     string `yaml:"names_dir"`
	LogFile      string `yaml:"logfile"`
	FilesUrlBase string `yaml:"files_url_base"`
	NamesUrlBase string `yaml:"names_url_base"`
}

func setup() goupConfig {
	var err error
	var data []byte
	var config goupConfig
	var usr *user.User

	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "<file to upload>")
	}

	usr, err = user.Current()
	if err != nil {
		log.Fatalln("Couldn't get current user:", err)
	}

	data, err = ioutil.ReadFile(path.Join(usr.HomeDir, ".goup.conf"))
	if err != nil {
		log.Fatalln("Error reading configuration file:", err)
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatalln("Error parsing configuration file:", err)
	}

	if config.LogFile != "" {
		logfile, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalln("Error opening logfile:", err)
		}

		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	}

	return config
}
