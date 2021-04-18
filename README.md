# telebot-chat-sbs
Q&amp;A's step by step bot for telegram

Made with Telebot framework
```bash
go get -u gopkg.in/tucnak/telebot.v2
```

# Overview

This simpel demo provides step-by-step communication with telegram bot. After every question, bot waits for an answer and then fires the next question.
Supports parallel processing w/o using database. Unique temporary file is creating for each new telegram bot user to store state machine's conditions and to gather replys.

