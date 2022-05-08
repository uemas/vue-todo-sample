FROM node:16

WORKDIR /usr/src/client

RUN npm install -g http-server
RUN npm install -g @vue/cli

COPY package.json ./
RUN npm install

COPY . .
RUN npm run build

CMD [ "http-server", "dist" ]
