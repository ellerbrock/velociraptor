package main

import (
	"errors"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	repack        = app.Command("repack", "Repack a binary.")
	repack_config = repack.Arg(
		"config_file", "The config file to repack into the binary.").
		Required().String()

	repack_binary = repack.Arg("binary", "The binary to repack.").Required().String()
)

func RepackClient(repack_binary string, repack_config string) error {
	binary, err := ioutil.ReadFile(repack_binary)
	if err != nil {
		return err
	}

	config_file, err := ioutil.ReadFile(repack_config)
	if err != nil {
		return err
	}

	start := regexp.MustCompile("# START C[O]NFIGURATION")
	end := []byte("# END CONFIGURATION")
	match := start.FindIndex(binary)
	if match == nil {
		return errors.New("No magic start")
	}

	offset := match[0]
	start_end := match[1]

	fmt.Printf("Found configuration signature at offset %v bytes\n", offset)

	new_binary := append(binary[:start_end], config_file...)
	new_binary = append(new_binary, end...)
	new_binary = append(new_binary,
		binary[start_end+len(config_file)+len(end):]...)

	err = ioutil.WriteFile(repack_binary, new_binary, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	command_handlers = append(command_handlers, func(command string) bool {
		switch command {
		case repack.FullCommand():
			err := RepackClient(*repack_binary, *repack_config)
			if err != nil {
				kingpin.FatalIfError(err, "Can not repack client")
			}
		default:
			return false
		}
		return true
	})
}
