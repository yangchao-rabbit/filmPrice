package spider

import (
	"encoding/json"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func Login() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	defer pw.Stop()
	defer browser.Close()

	ctx, _ := browser.NewContext()
	page, _ := ctx.NewPage()

	fmt.Println("👉 请扫码登录淘宝...")
	page.Goto("https://login.taobao.com/", playwright.PageGotoOptions{})
	time.Sleep(30 * time.Second)

	cookies, _ := ctx.Cookies()
	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		log.Fatalf("could not marshal cookies: %v", err)
	}

	err = ioutil.WriteFile("taobao_cookies.json", data, 0644)
	if err != nil {
		log.Fatalf("could not write cookies to file: %v", err)
	}
}

func LoginWithQRCodeScreenshot() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("❌ 启动 Playwright 失败: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true), // 无头模式，不弹出窗口
	})
	if err != nil {
		log.Fatalf("❌ 启动浏览器失败: %v", err)
	}
	defer pw.Stop()
	defer browser.Close()

	ctx, err := browser.NewContext(playwright.BrowserNewContextOptions{
		Locale: playwright.String("zh-CN"),
	})
	if err != nil {
		log.Fatalf("❌ 创建浏览器上下文失败: %v", err)
	}
	page, err := ctx.NewPage()
	if err != nil {
		log.Fatalf("❌ 创建页面失败: %v", err)
	}

	fmt.Println("🔗 打开淘宝登录页面...")
	_, err = page.Goto("https://login.taobao.com/")
	if err != nil {
		log.Fatalf("❌ 页面跳转失败: %v", err)
	}

	fmt.Println("📸 等待二维码加载...")
	page.WaitForSelector("img.qrcode-login") // 等待二维码元素出现
	time.Sleep(2 * time.Second)

	fmt.Println("📸 保存二维码截图: qrcode.png")
	_, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String("qrcode.png"),
		FullPage: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("❌ 截图失败: %v", err)
	}

	fmt.Println("📲 请扫描 qrcode.png 中的二维码完成登录（60秒内）")
	time.Sleep(60 * time.Second) // 等待登录

	// 获取 Cookie 并保存
	cookies, err := ctx.Cookies()
	if err != nil {
		log.Fatalf("❌ 获取 Cookie 失败: %v", err)
	}
	data, _ := json.MarshalIndent(cookies, "", "  ")
	_ = ioutil.WriteFile("taobao_cookies.json", data, 0644)
	fmt.Println("✅ Cookie 已保存为 taobao_cookies.json")
}

// LoadCookies 从文件读取 cookies.json
func LoadCookies(path string) ([]playwright.OptionalCookie, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cookies []playwright.OptionalCookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return nil, err
	}
	return cookies, nil
}

// FetchItemPrice 访问淘宝商品页面并获取标题和价格
func FetchItemPrice(itemURL string) (title string, price string, shop string, err error) {
	pw, err := playwright.Run()
	if err != nil {
		return "", "", "", fmt.Errorf("Playwright 启动失败: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		return "", "", "", fmt.Errorf("浏览器启动失败: %v", err)
	}
	defer browser.Close()

	context, err := browser.NewContext()
	if err != nil {
		return "", "", "", fmt.Errorf("浏览器上下文创建失败: %v", err)
	}

	cookies, err := LoadCookies("taobao_cookies.json")
	if err != nil {
		return "", "", "", fmt.Errorf("Cookie 读取失败: %v", err)
	}
	err = context.AddCookies(cookies)
	if err != nil {
		return "", "", "", fmt.Errorf("Cookie 注入失败: %v", err)
	}

	page, err := context.NewPage()
	if err != nil {
		return "", "", "", fmt.Errorf("页面创建失败: %v", err)
	}

	fmt.Println("🔗 打开商品页面：", itemURL)
	_, err = page.Goto(itemURL, playwright.PageGotoOptions{
		Timeout: playwright.Float(30000),
	})
	if err != nil {
		return "", "", "", fmt.Errorf("页面加载失败: %v", err)
	}
	time.Sleep(3 * time.Second)

	// 获取标题
	title, err = page.Title()
	if err != nil {
		title = "未知标题"
	}

	// 获取价格
	price = "未知价格"
	priceSelectors := []string{
		".tb-rmb-num",
		".price [class*=price]",
		"[class*=highlightPrice]",
	}
	for _, sel := range priceSelectors {
		el, _ := page.QuerySelector(sel)
		if el != nil {
			txt, _ := el.InnerText()
			if txt != "" {
				price = strings.Trim(txt, "¥￥ \n\t")
				break
			}
		}
	}

	// 获取店铺名
	shop = "未知店铺"
	shopSelectors := []string{
		".shop-name a",
		".tb-shop-name a",
		".shop-info .shop-name",
		"#J_ShopInfo .shop-name a",
		"[class*=shopName] span",
	}
	for _, sel := range shopSelectors {
		el, _ := page.QuerySelector(sel)
		if el != nil {
			txt, _ := el.InnerText()
			if txt != "" {
				shop = strings.TrimSpace(txt)
				break
			}
		}
	}

	return title, price, shop, nil
}
