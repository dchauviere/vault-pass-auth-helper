package pass

import (
	"fmt"
	"log/slog"
	"os/exec"
)

type Pass struct {
	BasePath    string
	DefaultName string
}

func (p *Pass) Store(server string, namespace string, token string) {
	slog.Debug("store", "server", server)
	out, err := exec.Command("pass", "insert", "-f", fmt.Sprintf("%s/%s/%s", p.BasePath, p.getSecretName(server), p.getSecretName(namespace))).Output()
	if err != nil {
		slog.Error("store", "server", server, "error", err)
		fmt.Println("")
	} else {
		fmt.Println(string(out))
	}
}

func (p *Pass) Get(server string, namespace string) {
	slog.Debug("get", "server", server)
	out, err := exec.Command("pass", "show", fmt.Sprintf("%s/%s", p.BasePath, p.getSecretName(server), p.getSecretName(namespace))).Output()
	if err != nil {
		slog.Error("get", "server", server, "error", err)
		fmt.Println("")
	} else {
		fmt.Print(string(out))
	}
}

func (p *Pass) Erase(server string, namespace string) {
	slog.Debug("erase", "server", server)
	out, err := exec.Command("pass", "rm", "-f", fmt.Sprintf("%s/%s", p.BasePath, p.getSecretName(server), p.getSecretName(namespace))).Output()
	if err != nil {
		slog.Error("erase", "server", server, "error", err)
		fmt.Println("")
	} else {
		fmt.Println(string(out))
	}
}

func (p *Pass) getSecretName(server string) string {
	if server == "" {
		return p.DefaultName
	}
	return server
}
