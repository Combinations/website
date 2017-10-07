run:
				docker build -t static-web .
				docker run -p=8080:8080 static-web
deploy:
				heroku container:push
