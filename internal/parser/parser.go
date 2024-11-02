package parser

import (
	"github.com/Resolution-hash/s1-report/config"

	"github.com/playwright-community/playwright-go"
)

var (
	OfflineStores []string
)

func ParseOfflineStores() ([]string, error) {

	userData, err := config.LoadUserConfig(true)
	if err != nil {
		return []string{}, err
	}

	// Инициализация Playwright
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	defer pw.Stop()

	// Запуск браузера
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	defer browser.Close()

	// Открытие новой страницы
	page, err := browser.NewPage()
	if err != nil {
		return nil, err
	}

	// Переход на целевую страницу
	url := "http://dmits01/Orion/SummaryView.aspx?ViewID=1" // Замените на нужный URL
	if _, err := page.Goto(url); err != nil {
		return nil, err
	}

	// Вводим логин и пароль
	if err := page.Locator("#ctl00_BodyContent_Username").Fill(userData.Login); err != nil {
		return nil, err
	}
	if err := page.Locator("#ctl00_BodyContent_Password").Fill("hFZ1qDLL"); err != nil {
		return nil, err
	}

	// Нажимаем на кнопку "Login"
	if err := page.Locator("#ctl00_BodyContent_LoginButton").Click(); err != nil {
		return nil, err
	}

	// Ожидание загрузки новой страницы (например, через наличие определенного элемента)
	if err := page.Locator("#Resource1782_ctl00_ctl01_ResourceWrapper_NodesRepeater_ctl01_link").WaitFor(); err != nil { // замените на нужный селектор
		return nil, err
	}

	// Полная ссылка для перехода
	fullLink := "http://dmits01/Orion/DetachResource.aspx?ViewID=1&ResourceID=1782&NetObject=&currentUrl=aHR0cDovL2RtaXRzMDEvT3Jpb24vU3VtbWFyeVZpZXcuYXNweD9WaWV3SUQ9MQ%3d%3d"

	// Переход по полной ссылке
	if _, err := page.Goto(fullLink); err != nil {
		return nil, err
	}

	// Используем Locator для получения всех элементов с классом "Property"
	propertyLocators := page.Locator(".Property")

	// Получаем количество элементов
	count, err := propertyLocators.Count()
	if err != nil {
		return nil, err
	}

	for i := 0; i < count; i++ {
		// Находим все теги <a> внутри текущего элемента "Property"
		anchors := propertyLocators.Nth(i).Locator("a")

		// Получаем количество тегов <a>
		anchorCount, err := anchors.Count()
		if err != nil {
			return nil, err
		}

		for j := 0; j < anchorCount; j++ {

			// Получаем текстовое содержимое тега <a>
			offlineStore, err := anchors.Nth(j).TextContent()
			if err != nil {
				return nil, err
			}

			OfflineStores = append(OfflineStores, offlineStore)
		}

	}
	return OfflineStores, nil
}
