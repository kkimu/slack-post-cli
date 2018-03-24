# slack-post-cli
Send slack message from command line.

## Install
```
go get github.com/kkimu/slack-post-cli
```

## Setup
- Get a slack webhook url, https://slack.com/services/new/incoming-webhook
- Create a config file `slack-post-cli.toml` to your home directory.
Example:

```toml
url = "your webhook url"

[Payload]
    UserName  = "slack-post-cli"
    IconEmoji = ":ghost:"
    LinkNames = 1
```

## Usage
```
$ slack-post-cli -m "Hello slack! :wave:" -to random -config [path-to-config-file]
```

Also, can send to user.
```
$ slack-post-cli -m "Hi user! :+1:" -to @username
```
