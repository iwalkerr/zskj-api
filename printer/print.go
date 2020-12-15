package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"xframe/pkg/utils/exec"
	pd "xframe/printer/proto"

	"google.golang.org/grpc"
)

const printAddress = ":8111"

// 初始化打印服务
func main() {
	startPrint()

	select {}
}

// 打印服务
func Server(printName, pdfPath string) error {
	// 连接服务器
	conn, err := grpc.Dial(printAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close() // // 确保连接最终被关闭

	client := pd.NewDdfPrintProviderClient(conn)
	reply, err := client.GoPtPdf(context.Background(), &pd.RequestParam{PrinterName: printName, DpfPath: pdfPath})
	if err != nil {
		return err
	}
	log.Printf("发送打印指令成功,code=%d,msg=%s", reply.Code, reply.Msg)
	return nil
}

// 启动打印服务
func startPrint() {
	// 杀死服务
	_, _ = exec.ExecCommand("kill -9 `lsof -ti:8111`")
	// 启动服务
	curDir, err := os.Getwd()
	if err == nil {
		curDir += "/"
	}
	curDir += "printer/print-server.jar"
	cmd := fmt.Sprintf("nohup java -jar %s > /dev/null 2>&1 &", curDir)
	if out, err := exec.ExecCommand(cmd); err != nil {
		log.Println("打印服务启动失败", err)
	} else {
		log.Println("打印服务启动成功！", out)
	}
}
