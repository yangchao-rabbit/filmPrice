package spider

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	LoginWithQRCodeScreenshot()
}

func TestFetchItemPrice(t *testing.T) {
	url := "https://item.taobao.com/item.htm?id=10643344222"
	title, price, shop, err := FetchItemPrice(url)
	if err != nil {
		fmt.Println("❌ 错误:", err)
		return
	}
	fmt.Println("📦 商品名称:", title)
	fmt.Println("📦 店铺名称:", shop)
	fmt.Println("💰 当前价格:", price)
}

func TestFetchItemPrice2(t *testing.T) {
	url := "https://item.taobao.com/item.htm?id=910237609779"
	title, price, shop, err := FetchItemPrice(url)
	if err != nil {
		fmt.Println("❌ 错误:", err)
		return
	}
	fmt.Println("📦 商品名称:", title)
	fmt.Println("📦 店铺名称:", shop)
	fmt.Println("💰 当前价格:", price)
}
