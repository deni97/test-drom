package main

import (
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// getAutohistoryJSON тут поголовье безголовья, сама работа с браузером
func getAutohistoryJSON(number string) string {
	// // для дебага явь головастого браузера
	//u := launcher.New().Headless(false).MustLaunch()
	//browser := rod.New().ControlURL(u).MustConnect()
	browser := rod.New().MustConnect()

	defer browser.MustClose()

	resChan := make(chan string)
	router := browser.HijackRequests()

	// слушаем запросы
	// POST https://auto.drom.ru/ajax/?mode=check_autohistory_gibdd_info
	// их бываец не один - повторяются покуда state не fetched
	router.MustAdd("*mode=check_autohistory_gibdd_info*", func(ctx *rod.Hijack) {
		ctx.MustLoadResponse()

		body := ctx.Response.Body()

		if strings.Contains(body, "fetched") {
			resChan <- body
		}
	})

	go router.Run()
	defer router.MustStop()

	// вперед на vin.drom.ru!
	page := browser.
		MustPage("https://vin.drom.ru").
		MustWaitLoad()

	// вбили номер в поле ввода
	page.
		MustElement(`input`).
		MustFocus().
		MustInput(number).
		MustType(input.Enter)

	// надеемся на товарища роутера - будь добр, перехвати!
	return <-resChan
}
