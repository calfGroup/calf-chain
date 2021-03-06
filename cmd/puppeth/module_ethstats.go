
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/log"
)


var ethstatsDockerfile = `
FROM puppeth/ethstats:latest

RUN echo 'module.exports = {trusted: [{{.Trusted}}], banned: [{{.Banned}}], reserved: ["yournode"]};' > lib/utils/config.js
`


var ethstatsComposefile = `
version: '2'
services:
  ethstats:
    build: .
    image: {{.Network}}/ethstats
    container_name: {{.Network}}_ethstats_1{{if not .VHost}}
    ports:
      - "{{.Port}}:3000"{{end}}
    environment:
      - WS_SECRET={{.Secret}}{{if .VHost}}
      - VIRTUAL_HOST={{.VHost}}{{end}}{{if .Banned}}
      - BANNED={{.Banned}}{{end}}
    logging:
      driver: "json-file"
      options:
        max-size: "1m"
        max-file: "10"
    restart: always
`


func deployEthstats(client *sshClient, network string, port int, secret string, vhost string, trusted []string, banned []string, nocache bool) ([]byte, error) {
	workdir := fmt.Sprintf("%d", rand.Int63())
	files := make(map[string][]byte)

	trustedLabels := make([]string, len(trusted))
	for i, address := range trusted {
		trustedLabels[i] = fmt.Sprintf("\"%s\"", address)
	}
	bannedLabels := make([]string, len(banned))
	for i, address := range banned {
		bannedLabels[i] = fmt.Sprintf("\"%s\"", address)
	}

	dockerfile := new(bytes.Buffer)
	template.Must(template.New("").Parse(ethstatsDockerfile)).Execute(dockerfile, map[string]interface{}{
		"Trusted": strings.Join(trustedLabels, ", "),
		"Banned":  strings.Join(bannedLabels, ", "),
	})
	files[filepath.Join(workdir, "Dockerfile")] = dockerfile.Bytes()

	composefile := new(bytes.Buffer)
	template.Must(template.New("").Parse(ethstatsComposefile)).Execute(composefile, map[string]interface{}{
		"Network": network,
		"Port":    port,
		"Secret":  secret,
		"VHost":   vhost,
		"Banned":  strings.Join(banned, ","),
	})
	files[filepath.Join(workdir, "docker-compose.yaml")] = composefile.Bytes()

	if out, err := client.Upload(files); err != nil {
		return out, err
	}
	defer client.Run("rm -rf " + workdir)

	if nocache {
		return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s build --pull --no-cache && docker-compose -p %s up -d --force-recreate --timeout 60", workdir, network, network))
	}
	return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s up -d --build --force-recreate --timeout 60", workdir, network))
}

type ethstatsInfos struct {
	host   string
	port   int
	secret string
	config string
	banned []string
}

func (info *ethstatsInfos) Report() map[string]string {
	return map[string]string{
		"Website address":       info.host,
		"Website listener port": strconv.Itoa(info.port),
		"Login secret":          info.secret,
		"Banned addresses":      strings.Join(info.banned, "\n"),
	}
}

func checkEthstats(client *sshClient, network string) (*ethstatsInfos, error) {
	infos, err := inspectContainer(client, fmt.Sprintf("%s_ethstats_1", network))
	if err != nil {
		return nil, err
	}
	if !infos.running {
		return nil, ErrServiceOffline
	}
	port := infos.portmap["3000/tcp"]
	if port == 0 {
		if proxy, _ := checkNginx(client, network); proxy != nil {
			port = proxy.port
		}
	}
	if port == 0 {
		return nil, ErrNotExposed
	}
	host := infos.envvars["VIRTUAL_HOST"]
	if host == "" {
		host = client.server
	}
	secret := infos.envvars["WS_SECRET"]
	config := fmt.Sprintf("%s@%s", secret, host)
	if port != 80 && port != 443 {
		config += fmt.Sprintf(":%d", port)
	}
	banned := strings.Split(infos.envvars["BANNED"], ",")

	if err = checkPort(host, port); err != nil {
		log.Warn("Ethstats service seems unreachable", "server", host, "port", port, "err", err)
	}
	return &ethstatsInfos{
		host:   host,
		port:   port,
		secret: secret,
		config: config,
		banned: banned,
	}, nil
}
