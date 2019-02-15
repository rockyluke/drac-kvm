// -*- go -*-

package supermicro

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"text/template"
)

// KvmSupermicroDriver is Supermicro specific folder for KVM driver.
//
type KvmSupermicroDriver struct {
	Host     string
	Username string
	Password string
	Version  int
}

// KvmSupermicroTemplate is the Supermicro KVM template
type KvmSupermicroTemplate struct {
	Template         string
	JARVersion       string
	NativeLibVersion string
	Arguments        []string
}

// KvmSupermicroContext is the Supermicro KVM context
type KvmSupermicroContext struct {
	Host             string
	Username         string
	Password         string
	JARVersion       string
	NativeLibVersion string
	Arguments        []string
}

const (
	// DefaultUsername is the default username on Supermicro KVM
	DefaultUsername = "ADMIN"
	// DefaultPassword is the default password on Supermicro KVM
	DefaultPassword = "ADMIN"
)

// KvmSupermicroVersions is a map of each viewer.jnlp template for
// the various Supermicro iKVM versions, keyed by version number
var KvmSupermicroVersions = map[int]KvmSupermicroTemplate{
	16921: {
		ikvm169,
		"1.69.21",
		"1.0.5",
		[]string{"5900", "623", "2", "0"},
	},
	16927: {
		ikvm169,
		"1.69.27",
		"1.0.8",
		[]string{"5900", "623", "0", "0", "0", "3520"},
	},
	16937: {
		ikvm169,
		"1.69.37",
		"1.0.12",
		[]string{"63630", "623", "0", "0", "1", "5900"},
	},
}

// Viewer returns a viewer.jnlp template filled out with the
// necessary details to connect to a particular DRAC host
func (d *KvmSupermicroDriver) Viewer() (string, error) {

	t, ok := KvmSupermicroVersions[d.Version]

	if !ok {
		msg := fmt.Sprintf("no support for iKVM v%d", d.Version)
		return "", errors.New(msg)
	}

	log.Printf("Found iKVM version %d", d.Version)
	// Generate a JNLP viewer from the template
	// Injecting the host/user/pass information
	buff := bytes.NewBufferString("")
	c := KvmSupermicroContext{d.Host, d.Username, d.Password, t.JARVersion, t.NativeLibVersion, t.Arguments}
	err := template.Must(template.New("viewer").Parse(t.Template)).Execute(buff, c)
	return buff.String(), err
}

// GetHost return Configured driver Host
func (d *KvmSupermicroDriver) GetHost() string {
	return d.Host
}

// GetUsername return Configured driver Username
func (d *KvmSupermicroDriver) GetUsername() string {
	return d.Username
}

// GetPassword return Configured driver Password
func (d *KvmSupermicroDriver) GetPassword() string {
	return d.Password
}

// EOF
