FROM node:lts-alpine AS build_stage


WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build


FROM node:lts-alpine AS run_stage


RUN npm install -g http-server

WORKDIR /app/

COPY --from=build_stage /app/dist /app/dist

COPY --from=build_stage /app/node_modules /app/node_modules

EXPOSE 8080

CMD [ "http-server", "dist" ]