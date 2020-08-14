docker pull mongo:4.4.0
docker run -itd -p 27017:27017 --restart always mongo:4.4.0
docker run -itd -p 8081:8081 --link mongodb:mongo  --restart always  mongo-express

挂载配置文件
D:\dockerv\mongo\mongo/mongod.conf
docker run --name some-mongo -v /my/custom:/etc/mongo -d mongo --config /etc/mongo/mongod.conf

挂载存储文件
docker run --name some-mongo -v /my/own/datadir:/data/db -d mongo



docker run  -itd -p 27017:27017 -v D:/dockerv/mongo/mongo:/etc/mongo -v D:/dockerv/mongo/data:/data/db --restart always --name mongodb  mongo:4.4.0