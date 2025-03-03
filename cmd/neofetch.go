package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func mostrarNeofetch(comando string) {
	user := os.Getenv("USER")
	hostname, _ := os.Hostname()
	osName := runtime.GOOS
	arch := runtime.GOARCH
	cpuInfo, _ := exec.Command("sh", "-c", "grep -m1 'model name' /proc/cpuinfo | cut -d':' -f2").Output()
	memInfo, _ := exec.Command("sh", "-c", "grep MemTotal /proc/meminfo | awk '{print $2}'").Output()

	color.Green(`
        \u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u28c0\u2840\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800
        \u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2880\u28e0\u28f4\u287f\u281b\u28bb\u28ff\u28f7\u2840\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800
        \u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2880\u28f4\u287f\u280b\u2800\u2800\u2800\u2808\u2819\u28bf\u2847\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800
        \u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u28fc\u281f\u2800\u2800\u2800\u2800\u2880\u28c0\u2800\u28bb\u28c7\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800
        \u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u28a0\u285f\u2800\u2800\u2800\u28a0\u287e\u281b\u281b\u28b7\u2840\u28bb\u2844\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u2800
        \u2800\u2800\u2800\u28e0\u28e4\u28e4\u28e4\u28e4\u287e\u2801\u2800\u2800\u2800\u28fe\u2801\u28e0\u2840\u2808\u28f7\u2800\u28b7\u28e4\u28e4\u28e4\u28c4\u2800\u2800\u2800\u2800\u2800
        \u2800\u2800\u28b8\u284f\u2800\u28e0\u28f4\u285f\u2800\u2800\u2800\u2800\u2818\u28b7\u2840\u283b\u2803\u28e0\u287f\u2800\u2808\u28b7\u2844\u2818\u28ff\u2840\u2800\u2800\u2800\u2800
        \u2800\u2800\u2818\u28f7\u2840\u283b\u283f\u2803\u2800\u2840\u2800\u2800\u2800\u2808\u283b\u2836\u281e\u2809\u2800\u2800\u2800\u2818\u28bf\u2840\u28ff\u2803\u2800\u2800\u2800\u2800
        \u2800\u2800\u2800\u2808\u283b\u28e6\u28c4\u28c0\u28c0\u283b\u28e6\u2840\u2800\u2800\u2800\u2800\u2800\u2800\u2800\u28e0\u28c0\u28c0\u28fe\u287f\u280b\u2800\u2800\u2800\u2800\u2800
        \u2800\u2800\u2800\u2800\u2800\u2800\u2809\u2819\u281b\u281b\u283b\u2803\u2800\u2800\u2800\u2800\u2800\u2800\u2818\u283f\u283f\u281b\u2801\u2800\u2800\u2800\u2800\u2800\u2800\u2800

        \U0001f98e Salamandra Shell \U0001f98e
	`)

	color.Yellow("Usuario: ")
	fmt.Println(user)
	color.Yellow("Host: ")
	fmt.Println(hostname)
	color.Yellow("OS: ")
	fmt.Println(osName)
	color.Yellow("Arquitectura: ")
	fmt.Println(arch)
	color.Yellow("CPU: ")
	fmt.Println(strings.TrimSpace(string(cpuInfo)))
	color.Yellow("Memoria: ")
	fmt.Printf("%s KB\n", strings.TrimSpace(string(memInfo)))
}
