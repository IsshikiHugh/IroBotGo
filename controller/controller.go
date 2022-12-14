package controller

import (
	"IroBot/config"
	"IroBot/model"
	"fmt"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

type BotEnvironment struct {
	Manager *OPQBot.BotManager
	Conf    *model.Configuration
}

func PowerBot() error {
	var (
		err error
		bot BotEnvironment
	)

	bot.Conf = config.Config()
	if err != nil {
		return err
	}
	bot.Manager = OPQBot.NewBotManager(bot.Conf.Basic.Qid, bot.Conf.Basic.Url)
	bot.Manager.SetMaxRetryCount(bot.Conf.Basic.Retry)
	err = bot.Manager.Start()
	if err != nil {
		return err
	}

	bot.regEvent()

	bot.Manager.Wait()
	defer bot.Manager.Stop()
	return nil
}

func (bot *BotEnvironment) regEvent() {
	_, err := bot.Manager.AddEvent(OPQBot.EventNameOnFriendMessage, bot.PrivateChatEvents)
	if err != nil {
		logrus.Error(fmt.Sprintf("Add private chat events failed with error messages: [ %s ]", err.Error()))
	}

	_, err = bot.Manager.AddEvent(OPQBot.EventNameOnGroupMessage, bot.GroupChatEvents)
	if err != nil {
		logrus.Error(fmt.Sprintf("Add group chat events failed with error messages: [ %s ]", err.Error()))
	}

}
