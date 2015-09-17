all:
	gb build

d:	all heroku_push heroku_logs
	@true

heroku_push:
	git push -f heroku master

heroku_logs:
	heroku logs -p web -t
