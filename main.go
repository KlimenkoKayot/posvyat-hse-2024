package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Person struct {
	Full_name         string
	Telegram_username string
	Phone             string
	Education_program string
	Gender            string
}

func main() {
	Persons := []Person{
		{
			Full_name:         "Клименко Георгий Сергеевич",
			Telegram_username: "kayot123",
			Phone:             "+7 *** *** ** **",
			Education_program: "Программная инженерия",
			Gender:            "Мужской",
		},
	}

	for _, cur := range Persons {
		val := url.Values{}

		val.Add("full_name", cur.Full_name)
		val.Add("telegram_username", cur.Telegram_username)
		val.Add("phone", cur.Phone)
		val.Add("education_program", cur.Education_program)
		val.Add("gender", cur.Gender)
		val.Add("18yo_bool", "true")
		val.Add("needs_transfeer_bool", "true")

		// https://cstati.ru/v1/register
		resp, err := http.PostForm("https://cstati.ru/v1/register", val)
		// resp, err := http.PostForm("https://cstati.ru/v1/registration_form", val)
		bytes, _ := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		code := string(bytes)

		url := ""
		if strings.Index(code, "t.me/csposvyat2024_bot") != -1 {
			idxstart := strings.Index(code, "t.me/csposvyat2024_bot")
			for i := idxstart; true; i++ {
				if string(code[i]) == `'` || string(code[i]) == `"` {
					break
				}
				url += string(code[i])
			}
			fmt.Printf("%s\n%s\n\n", val.Get("full_name"), url)
		} else if strings.Index(code, "t.me") != -1 {
			idxstart := strings.Index(code, "t.me")
			for i := idxstart; true; i++ {
				if string(code[i]) == `'` || string(code[i]) == `"` {
					break
				}
				url += string(code[i])
			}
			fmt.Printf("%s\n%s\n\n", val.Get("full_name"), url)
		} else if strings.Index(code, "tg://") != -1 {
			idxstart := strings.Index(code, "tg://")
			for i := idxstart; true; i++ {
				if string(code[i]) == `'` || string(code[i]) == `"` {
					break
				}
				url += string(code[i])
			}
			fmt.Printf("%s\n%s\n\n", val.Get("full_name"), url)
		} else {
			url, _ := resp.Location()
			urls := ""
			if url != nil {
				urls = url.String()
			}
			fmt.Printf("%s\n%s\n%s\n%s\n", val.Get("full_name"), code, urls, resp.Request.URL)
			return
		}

		resp.Body.Close()
	}
}
