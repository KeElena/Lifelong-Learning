mkdir confg html log
mkdir ./confg/conf.d
touch ./confg/nginx.conf
docker run -d --name mynginx -p 80:80\
	\-v $PWD/confg/nginx.conf:/etc/nginx/nginx.conf\
	\-v $PWD/confg/conf.d:/etc/nginx/conf.d\
	\-v $PWD/log:/var/log/nginx\
	\-v $PWD/html:/usr/share/nginx/html nginx

