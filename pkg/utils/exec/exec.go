package exec

import (
	"bytes"
	"os/exec"
)

// 运行cmd
func ExecCommand(arg string) (string, error) {
	cmd := exec.Command("bash", "-c", arg)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out //剧透，坑在这里
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		// log.Printf("运行命令行 %s ===> 错误提示: %s", err.Error(), stderr.String())
		return "", err
	}
	return out.String(), nil
}
