FROM node:16.13.1-alpine3.13

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

EXPOSE 8090

CMD ["npm","run","start"]