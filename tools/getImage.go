package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetImage(url string, name string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTPリクエストエラー:", err)
		return
	}
	defer response.Body.Close()

	// レスポンスのステータスコードを確認
	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTPリクエストエラー:", response.Status)
		return
	}
	fileName := "images/" + name + ".png"
	// 画像ファイルを作成
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("ファイル作成エラー:", err)
		return
	}
	defer file.Close()

	// レスポンスのボディを画像ファイルにコピー
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("ファイル書き込みエラー:", err)
		return
	}

	fmt.Println("画像を取得しました")
}
