// -*- go -*-

package hp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

// KvmHpDriver is HP specific folder for KVM driver.
//
type KvmHpDriver struct {
	Host     string
	Username string
	Password string
	Version  int
}

const (
	// DefaultUsername is the default username on HP iLO
	DefaultUsername = "Administrator"
	// DefaultPassword is the default password on HP iLO
	DefaultPassword = ""
)

// Viewer that logs in, fetch the sessionKey cookie to be able
// to generate a correct jnlp. With HP we can use `jnlp_template.html`
// url to fetch current jnlp template.
func (d *KvmHpDriver) Viewer() (string, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Dial: func(netw, addr string) (net.Conn, error) {
			deadline := time.Now().Add(5 * time.Second)
			c, err := net.DialTimeout(netw, addr, time.Second*5)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(deadline)
			return c, nil
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	// Post parameters to login to iLO
	values := map[string]string{"method": "login", "user_login": d.Username, "password": d.Password}
	jsonValue, _ := json.Marshal(values)

	if res, err := client.Post("https://"+d.Host+"/json/login_session", "", bytes.NewBuffer(jsonValue)); err == nil {
		defer res.Body.Close()
		if res.StatusCode == 200 {
			var f interface{}

			// Fetch sessionKey from json response using an
			// interface in order to build a cookie
			bodyBytes, _ := ioutil.ReadAll(res.Body)
			err = json.Unmarshal(bodyBytes, &f)

			if err != nil {
				fmt.Println("Couldn't decode json", err)
			}

			m := f.(map[string]interface{})
			sessionKey := m["session_key"].(string)

			cookie := http.Cookie{Name: "session_key", Value: sessionKey}
			req, _ := http.NewRequest("GET", "https://"+d.Host+"/html/jnlp_template.html", nil)
			req.AddCookie(&cookie)

			if res, err := client.Do(req); err == nil {
				defer res.Body.Close()
				if res.StatusCode == 200 {
					bodyBytes, _ := ioutil.ReadAll(res.Body)
					bodyString := string(bodyBytes)

					// We need to:
					// - replace placeholder with actual values
					// - skip the first and last line of the jnlp
					// template
					r := strings.NewReplacer("<%= this.baseUrl %>", "https://"+d.Host+"/",
						"<%= this.sessionKey %>", sessionKey,
						"<%= this.langId %>", "en")

					jnlp := r.Replace(bodyString)
					_jnlp := strings.Split(jnlp, "\n")
					jnlp = strings.Join(_jnlp[1:len(_jnlp)-1], "\n")

					return jnlp, err
				}

				return "", errors.New("Couldn't fetch jnlp template")
			}
		}
	}
	return "", errors.New("Couldn't login to iLO")
}

// GetHost return Configured driver Host
func (d *KvmHpDriver) GetHost() string {
	return d.Host
}

// GetUsername return Configured driver Username
func (d *KvmHpDriver) GetUsername() string {
	return d.Username
}

// GetPassword return Configured driver Password
func (d *KvmHpDriver) GetPassword() string {
	return d.Password
}

// EOF
