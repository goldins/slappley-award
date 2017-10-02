## Slack of all Science | WIP

Simple server to return images from [masterofallscience.com](https://masterofallscience.com)

Build and run with:

```sh
docker build -t goldins/science .
docker run -it --rm --name science -p 8080:80 -e SCIENCE_TOKEN=<YOUR_SLACK_TOKEN> goldins/science:latest
```

### Inspiration/Resources:

http://guzalexander.com/2017/09/15/cowsay-slack-command.html

https://github.com/snare/humorbot

