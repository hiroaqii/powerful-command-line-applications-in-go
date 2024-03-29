package scan

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

var (
	ErrExists    = errors.New("Host already in the list")
	ErrNotExists = errors.New("Host not in the list")
)

type HostsList struct {
	Hosts []string
}

func (hl *HostsList) search(host string) (bool, int) {
	sort.Strings(hl.Hosts)

	i := sort.SearchStrings(hl.Hosts, host)
	if i < len(hl.Hosts) && hl.Hosts[i] == host {
		return true, i
	}

	return false, -1
}

func (hl *HostsList) Add(host string) error {
	if ok, _ := hl.search(host); ok {
		return ErrExists
	}

	hl.Hosts = append(hl.Hosts, host)
	return nil
}

func (hl *HostsList) Remove(host string) error {
	if found, i := hl.search(host); found {
		hl.Hosts = append(hl.Hosts[:i], hl.Hosts[i+1:]...)
		return nil
	}

	return fmt.Errorf("%w: %s", ErrNotExists, host)
}

func (hl *HostsList) Load(hostFile string) error {
	f, err := os.Open(hostFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hl.Hosts = append(hl.Hosts, scanner.Text())
	}

	return nil
}

func (hl *HostsList) Save(hostFile string) error {
	output := ""

	for _, h := range hl.Hosts {
		output += fmt.Sprintln(h)
	}

	return ioutil.WriteFile(hostFile, []byte(output), 0644)
}
