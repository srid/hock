# hock

Tail buf, a Sep 2015 offsite hack project.

## Deploying

```
heroku create
heroku buildpacks:set https://github.com/paxan/heroku-buildpack-gb.git
git push heroku master
```

## Configuring drains

Go to your app and add a drain pointing to:

```
https://<your-hock>.herokuapp.com/logs/<appname>
```
