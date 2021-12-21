package commandhandlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type ImportChannelCommand struct {
	GuildID   string
	ChannelID string
}

func ImportChannelCommandHandler(c *ImportChannelCommand, s *discordgo.Session, log *logrus.Entry) {
	pinned, err := s.ChannelMessagesPinned(c.ChannelID)
	if err != nil {
		log.WithError(err).Error("Could not get channel pins")
		return
	}

	for _, m := range pinned {
		PinMessageCommandHandler(&PinMessageCommand{
			GuildID: c.GuildID,
			Message: m,
		}, s, log)
	}
}
