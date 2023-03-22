FROM node:14-alpine
WORKDIR /app/
ARG MOD

COPY package.json /app/
RUN npm install
COPY . .
RUN npm run-script build.${MOD}
WORKDIR /app/dist/
COPY node-server.js .
COPY express-package.json package.json
RUN npm install
ENTRYPOINT node node-server.js
