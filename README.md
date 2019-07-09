# botsend

A small docker image for sending telegram notification via bot.

## Usage

```bash
docker run \
    -e BOT_TOKEN=123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11 \
    -e CHAT=12345678 \
    botsend \
    "your message"
```

If you want send messages to yourself, you can obtain your chat_id using info bot:

https://t.me/userinfobot