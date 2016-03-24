package core

const (
   MAIL_SENDER_NAME = "Anna Berry"
   MAIL_SENDER_EMAIL = "anna@localization.expert"
   MAIL_DOMAIN = "localization.expert"

   SEARCH_APPS_LIMIT = 1
   PARSER_APPS_LIMIT = 1
   PARSER_MAX_URLS = 10

   EMAIL_REGEXP = `[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}`
)

var BLACKLIST_DOMAINS = []string{
    "google.com",
    "microsoft.com",
    "apple.com",
    "facebook.com",
    "adobe.com",
    "twitter.com",
    "youtube.com",
    "instagram.com",
    "mailto",
}

var BLACKLIST_EMAILS = []string{
    "@google.com",
    "@microsoft.com",
    "@apple.com",
    "@facebook.com",
    "@adobe.com",
    "@twitter.com",
}

var BLACKLIST_EXTENSTIONS = []string{
    ".exe",
    ".zip",
    ".rar",
    ".gif",
    ".png",
    ".jpg",
    ".jpeg",
    ".css",
    ".js",
}
