# Tweetbot - Imgur Proxy

This is a small Go proxy to use with Tweetbot to upload images to imgur.

## Usage

The proxy requires one env var to run, `IMGUR_CLIENT_ID`.
You can view your existing apps [here](http://imgur.com/account/settings/apps).

When you have it running (I deploy mine on a Dokku instance), point your Tweetbot instance at it:

![Adding to Tweetbot](http://i.imgur.com/dB1ICNa.gif)

## License

MIT. For more info see `LICENSE.TXT`.
