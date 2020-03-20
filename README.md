LINE Bot Icon/Nick Name Swtich demo: Sample code how to use Icon/Nick name switch in LINE Bot in Go
==============

[![Join the chat at https://gitter.im/kkdai/line-bot-icon-switch](https://badges.gitter.im/kkdai/line-bot-icon-switch.svg)](https://gitter.im/kkdai/line-bot-icon-switch?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![GoDoc](https://godoc.org/github.com/kkdai/line-bot-icon-switch.svg?status.svg)](https://godoc.org/github.com/kkdai/line-bot-icon-switch)  [![Build Status](https://travis-ci.org/kkdai/line-bot-icon-switch.svg?branch=master)](https://travis-ci.org/kkdai/line-bot-icon-switch.svg) [![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/line-bot-icon-switch)](https://goreportcard.com/report/github.com/kkdai/line-bot-icon-switch)

Icon / Nick name Switch in LINE Bot
=============

![](https://developers.line.biz/assets/img/icon-nickname-switch.7ad52e1a.jpg)

Icon/Nick name switch is new feature of LINE Bot open for public in 2020/03/17. (Please refer this [news](https://developers.line.biz/zh-hant/news/2020/03/17/icon-nickname-switch/) )

How to Use it
=============




Installation and Usage
=============

### 1. Got A LINE Bot API devloper account

[Make sure you already registered](https://business.line.me/zh-hant/services/bot), if you need use LINE Bot.

### 2. Just Deploy the same on Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Remember your heroku, ID.

<br><br>

### 3. Go to LINE Developer Console, setup basic API

3.1 Close auto-reply setting on "Messaging API" Tab.

3.2 Setup your basic account information. Here is some info you will need to know.

- `Callback URL`: https://{YOUR_HEROKU_SERVER_ID}.herokuapp.com:443/callback
- Verify your webhook.

3.3 You will get following info, need fill back to Heroku.

- Channel Secret
- Channel Access Token (You need to issue one here)


### 4. Back to Heroku again to setup environment variables

- Go to dashboard
- Go to "Setting"
- Go to "Config Variables", add following variables:
	- "ChannelSecret"
	- "ChannelAccessToken"

It all done.


### Video Tutorial:

- [How to deploy LINE BotTemplate](https://www.youtube.com/watch?v=xpP51Kwuy2U)
- [Hoe to modify your LINE BotTemplate code](https://www.youtube.com/watch?v=ckij73sIRik)


### Chinese Tutorial:

如果你看得懂繁體中文，這裡有[中文的介紹](http://www.evanlin.com/create-your-line-bot-golang/) 


License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

