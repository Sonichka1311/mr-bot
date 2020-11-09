# Merge requests telegram bot.
Bot which checks gitlab repo for new merge requests and sends message to telegram group, choosing responsible for this MR.

## Run
```
git clone https://github.com/Sonichka1311/mr-bot.git
export token=... # gitlab access token
export bot_token=... # bot token from BotFather
export project_id=... # gitlab project ID
make run
```
