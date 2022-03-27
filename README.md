# fintech-test

How to start :

- Run db with docker-compose up --build
- Build go app with go build -o app
- Send request to http://localhost:3000/slowest-queries/:page/size/:pagesize/type/:query-type(SELECT,INSERT...)