package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	const (
		token          string = "TOKEN"
		channelName1   string = "CHANNEL NAME"
		channelID1     string = "CHANNEL ID"
		updateDelayMin int64  = 10
	)
	koreaUTC := int((9 * time.Hour).Seconds())

	client, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("discordgo.New Error!", err)
		return
	}

	func() {
		var leftDays int
		var newChannelName string
		var day time.Time
		day = time.Date(YEAR, MONTH, DAY, HOUR, MIN, SEC, NANOSEC, time.FixedZone("Korea", koreaUTC))//set it to D-Day's tomorrow 0h 0m 0s 0ms
		leftDays = getLeftDays(day)
		newChannelName = string(channelName1 + " D-" + strconv.Itoa(leftDays))
		fmt.Println("Now up and running!")

		channel, err := client.Channel(channelID1)
		if err != nil {
			fmt.Println("client.Channel Error!", err)
			time.Sleep(1 * time.Minute)
			return
		}
		channelEdit := discordgo.ChannelEdit{
			Name:                 newChannelName,
			Topic:                channel.Topic,
			NSFW:                 channel.NSFW,
			Position:             channel.Position,
			Bitrate:              channel.Bitrate,
			UserLimit:            channel.UserLimit,
			PermissionOverwrites: channel.PermissionOverwrites,
			ParentID:             channel.ParentID,
			RateLimitPerUser:     channel.RateLimitPerUser}
		channel, err = client.ChannelEditComplex(channelID1, &channelEdit)
		if err != nil {
			fmt.Println("ChannelEdit Error!", err)
			return
		}
		fmt.Println("Name changed!", newChannelName)

		for {
			time.Sleep(10 * time.Minute)
			nowLeftDays := getLeftDays(day)

			if nowLeftDays == leftDays {
				continue
			}

			leftDays = nowLeftDays
			if leftDays == 0 {
				newChannelName = string(channelName1 + " D-DAY")
			} else {
				newChannelName = string(channelName1 + " D-" + strconv.Itoa(leftDays))
			}

			channelEdit.Name = newChannelName
			client.ChannelEditComplex(channelID1, &channelEdit)
			fmt.Println("D-Day Updated!", leftDays)
		}
	}()
}

func getLeftDays(day time.Time) int {
	daySec := int64((24 * time.Hour).Seconds())

	dDay := day.Unix()          //seconds int64
	nowDay := time.Now().Unix() //seconds int64
	return int((dDay - nowDay) / daySec)
}
