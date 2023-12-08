package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"go.step.sm/crypto/randutil"
)

const random_id_len = 8

type User struct {
	Name     string   `json:"name"`
	Socket   net.Conn `json:"-"`
	Messages []string `json:"messages"`
	Banned   bool     `json:"banned"`
}

type CommandType int

const (
	Ban CommandType = iota
	Mute
	Elevate
)

type Executer interface {
	Execute() error
}

type Command struct {
	Type        CommandType `json:"type"`
	Description string      `json:"description"`
	Target      User        `json:"user"`
}

func (c *Command) Execute() error {
	switch c.Type {
	case Ban:
		fmt.Println("Call Ban command in user", c.Target)
	case Mute:
	case Elevate:
	}
	return nil
}

type CommandsInfo struct {
	Total  int `json:"total"`
	Banned int `json:"banned"`
}

func main() {
	commands := make(map[string]*Command, 10)

	for i := 0; i < 2; i++ {
		msg, _ := randutil.Alphanumeric(random_id_len)
		key, _ := randutil.Alphanumeric(random_id_len)
		commands[key] = &Command{
			Type:        Ban,
			Description: "Ban user",
			Target: User{
				Name:     "user@" + key,
				Socket:   nil,
				Messages: []string{msg, key},
			},
		}
	}
	buffer := bytes.NewBuffer([]byte{})

	enc := json.NewEncoder(buffer)
	enc.SetIndent("", "\t")

	/* enc.Encode(commands) */

	cmd := struct {
		CmdInfo  CommandsInfo        `json:"metadata"`
		Commands map[string]*Command `json:"commands"`
	}{
		CmdInfo:  CommandsInfo{len(commands), len(commands)},
		Commands: commands,
	}

	enc.Encode(cmd)

	/*
		for _, c := range commands {

			err := enc.Encode(c)
			if err != nil {
				log.Println(err)
				continue
			}

			decoded := bytes.NewReader(buffer.Bytes())
			dec := json.NewDecoder(decoded)
			err = dec.Decode(decoded)

			if err != nil {
				log.Println(err)
			}

			fmt.Println(decoded)
		}*/

	fmt.Println(buffer)

}
