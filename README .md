
# Discord Bot

It's a Simple Discord Bot designed to respond "pong" if "ping" input is served. GoLang is used to develop this  mini bot via using third praty library to use it.


## Deployment

You Need To add "config.json" in root with object "Token" and "BotPrefix"

Example config.json:


{
    "Token": "Your App Token",
    "BotPrefix" :"!"
}

How To get Token : 
To deploy this project run

```bash
  go mod init //Download Dependencies
  go build
  go run main.go
```

