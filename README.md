# go-discord-bot
This is my attempt at building a discord bot using Golang. The [discordGo](https://github.com/bwmarrin/discordgo) package is the primary driver of the bot.

### How to run
In development I simply run `make` to build (for macOS) and run the bot in one command. 

- Run `make build` to build.
- Then run `make run` to start the bot.

If you need to build for linux or windows, simply uncomment the needed build command in the Makefile.

**Please Note:**
This bot is only intended to run on my personal discord servers at this time. Therefore, you may see some hardcoded values that are specific to these servers. You will need to update these specific values if you plan to fork/clone this bot and add it to your server.
