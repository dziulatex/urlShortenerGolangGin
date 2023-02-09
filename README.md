# urlShortenerGolangGin
This project is an url shortener writen in golang, gin framework.
<br>
In .env there are such variables as:
<br>
SHORTED_URL_LEN length of generated url shorting id.
<br>
REDIS_URL redis url.
<br>
REDIS_PWD redis password.
<br>
REDIS_DB_NUM redis db number, default 0
<br>
CACHE_DRIVER cache driver: redis,
<br>
DATABASE_URL=database url, needs ?parseTime=true&loc=Local at the end because i have not yet discovered how to parse time object without from database well &parseTime=true. And &loc=Local is because my createdAt, and updatedAt fields were with 1 hours diff, dont know why.
<br>
INTERNAL_TOKEN="Sy%beyP&Npj!u+h49=C6" internal token to access debugging endpoint /debug GET.
<br>
Swagger doc is there for your inquisitiveness.
![image](https://user-images.githubusercontent.com/38464243/217682085-c69d3894-74cb-4023-b5dd-8d153d73fd88.png)
![image](https://user-images.githubusercontent.com/38464243/217683075-c9f3f7ea-5a66-4afd-8ced-eed778b1caad.png)

