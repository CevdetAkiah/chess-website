FROM node:21.1-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json      ./
COPY package-lock.json ./

RUN npm install

COPY . ./

EXPOSE 8080

CMD ["npm", "start"]
