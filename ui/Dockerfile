FROM node:16-alpine

WORKDIR /usr/frontend

COPY ./ui/package.json ./ui/yarn.lock ./

RUN yarn install --frozen-lockfile

COPY . .

ENV NODE_ENV=production

RUN yarn build

EXPOSE 3000

CMD [ "yarn", "start" ]