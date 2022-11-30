package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Data struct {
	Data []struct {
		Name string `json:"CharacterName"`
		Gold string `json:"Coin"`
	}
}

func setActivity(s *discordgo.Session, r *discordgo.Ready) {
	err := s.UpdateListeningStatus("Yorushika Gekko Live")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == config.Prefix+"totalgold" {
		msg, _ := s.ChannelMessageSendReply(m.ChannelID, "Mohon ditunggu, mungkin dibutuhkan 10 detik untuk mendapatkan hasilnya...", m.Reference())
		totalGold, _ := getTotalGold()
		now := time.Now()
		time := fmt.Sprintf("%v - %v - %v", now.Day(), now.Month(), now.Year())
		_ = s.ChannelMessageDelete(m.ChannelID, msg.ID)
		_, _ = s.ChannelMessageSendReply(m.ChannelID, time+"\n"+strconv.Itoa(totalGold)+" :coin:\nDengan toleransi kesalahan berkisar 10-30k gold.", m.Reference())
	}

	if strings.Contains(m.Content, config.Prefix+"register") {
		regex, _ := regexp.Compile(`[^\s]*$`)
		regexContent := regex.FindAllString(m.Content, -1)
		splitContent := strings.Split(regexContent[0], ":")

		if len(splitContent) < 4 {
			_, _ = s.ChannelMessageSendReply(m.ChannelID, "Salah format, yang bener itu\n`suis!register username:password:pin:email`", m.Reference())
			return
		}

		username, password, pin, email := splitContent[0], splitContent[1], splitContent[2], splitContent[3]
		register := doRegister(username, password, pin, email)
		_, _ = s.ChannelMessageSendReply(m.ChannelID, register, m.Reference())
	}
	if strings.Contains(m.Content, config.Prefix+"randomregister") {
		split := strings.Split(m.Content, " ")
		amount := len(split)
		if amount > 1 {
			_amount, err := strconv.Atoi(split[1])
			if err != nil {
				s.ChannelMessageSendReply(m.ChannelID, "Parameter tidak sesuai, harusnya suis!randomregister angka", m.Reference())
			}
			amount = _amount
		}

		for i := 0; i < amount; i++ {
			result, payload := doRandomRegister()
			if result != "Success create account" {
				s.ChannelMessageSendReply(m.ChannelID, "Gagal membuat akun dikarenakan"+result, m.Reference())
			}

			s.ChannelMessageSendReply(m.ChannelID, "Berhasil membuat akun\n"+payload, m.Reference())
			time.Sleep(3 * time.Second)
		}
	}
}

func Run() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	goBot.AddHandler(setActivity)
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot online, with prefix: " + config.Prefix)
}

func getTotalGold() (int, error) {
	var (
		data  Data
		_coin int
		url   string
		Gold  int
	)
	user := map[string]bool{}

	client := http.Client{}
	for i := 0; i <= 5; i++ {
		url = baseUrl + "api/rank/char/gold/" + strconv.Itoa(i)
		if i == 0 {
			url = baseUrl + "api/rank/char/gold"
		}

		resp, err := client.Post(url, "application/json", nil)
		if err != nil {
			fmt.Println("error while trying to receive the response from api.")
			return 0, err
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}

		regex, _ := regexp.Compile(`\d{4}$`)
		for _, item := range data.Data {
			if user[item.Name] {
				continue
			}

			split := regex.Split(item.Gold, -1)

			gold, err := strconv.Atoi(split[0])
			if err != nil {
				fmt.Println(err.Error())
				return 0, err
			}

			findSmallCoin := regex.FindString(item.Gold)
			smallCoin, err := strconv.Atoi(findSmallCoin)
			if err != nil {
				fmt.Println(err.Error())
				return 0, err
			}

			Gold += gold
			_coin += smallCoin
			user[item.Name] = true
		}
		fmt.Println(Gold)
		splitCoin := regex.Split(strconv.Itoa(_coin), 2)[0]
		coin, _ := strconv.Atoi(splitCoin)
		Gold += coin
		time.Sleep(1 * time.Second)
	}

	Gold += (2000 * 18) * 5
	return Gold, nil
}
