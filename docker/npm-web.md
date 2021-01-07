
330MB
```
FROM node:10-alpine
 
WORKDIR /app
COPY app /app
RUN npm install -g webserver.local
RUN npm install && npm run build
 
EXPOSE 3000
CMD webserver.local -d ./build
```

91.5MB
```
FROM node:10-alpine AS build
WORKDIR /app
COPY app /app
RUN npm install && npm run build
 
 
FROM node:10-alpine
WORKDIR /app
RUN npm install -g webserver.local
COPY --from=build /app/build ./build
EXPOSE 3000
CMD webserver.local -d ./build
```
22.4MB
```
FROM node:10-alpine AS build
WORKDIR /app
COPY app /app
RUN npm install && npm run build
 
 
FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```