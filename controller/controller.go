package controller

import (
	"IroBot/config"
	"IroBot/model"
	"fmt"

	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"
)

var Bot model.BotEnvironment

/*
 * @brief Initialization of the bot.
 */
func PowerBot() error {
	var (
		err error
	)

	// Load configuration.
	Bot.Conf = config.Config()
	if err != nil {
		return err
	}
	// Generate the manager which master the function of the bot.
	Bot.Manager = OPQBot.NewBotManager(Bot.Conf.Basic.Qid, Bot.Conf.Basic.Url)
	Bot.Manager.SetMaxRetryCount(Bot.Conf.Basic.Retry)
	err = Bot.Manager.Start()
	if err != nil {
		return err
	}

	// Register the events.
	regEvent()

	Bot.Manager.Wait()
	defer Bot.Manager.Stop()
	return nil
}

/*
 * @brief Register the events.
 */
func regEvent() {
	_, err := Bot.Manager.AddEvent(OPQBot.EventNameOnFriendMessage, PrivateChatEvents)
	if err != nil {
		logrus.Error(fmt.Sprintf("Add private chat events failed with error messages: [ %s ]", err.Error()))
	}

	_, err = Bot.Manager.AddEvent(OPQBot.EventNameOnGroupMessage, GroupChatEvents)
	if err != nil {
		logrus.Error(fmt.Sprintf("Add group chat events failed with error messages: [ %s ]", err.Error()))
	}

}
