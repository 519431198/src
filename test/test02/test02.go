package main

import (
	"fmt"
	"github.com/anacrolix/torrent"
	"log"
)

func main() {
	//if len(os.Args) < 2 {
	//	log.Fatal("Usage: go run db.go <magnet-link>")
	//}

	magnetURI := "magnet:?xt=urn:btih:947c36ae4f4fb3541e05e13e30022a8c46c24a36"

	// 创建一个新的 torrent 客户端
	client, err := torrent.NewClient(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 添加磁力链接
	t, err := client.AddMagnet(magnetURI)
	if err != nil {
		log.Fatal(err)
	}

	// 等待 torrent 元数据下载完成
	<-t.GotInfo()

	// 打印文件信息
	fmt.Println("Files:")
	for _, file := range t.Files() {
		fmt.Println("-", file.Path())
	}

	// 下载所有文件
	fmt.Println("Downloading...")
	t.DownloadAll()
	<-t.GotInfo()

	fmt.Println("Download complete!")
}
