package utils

import tele "gopkg.in/telebot.v4"

func GetMentions(entities tele.Entities, text string) []string {
	mentions := make([]string, 0)
	for _, entity := range entities {
		if entity.Type == tele.EntityMention {
			mentions = append(mentions, text[entity.Offset:entity.Offset+entity.Length])
		}
	}

	return mentions
}
