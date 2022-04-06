FROM golang:alpine AS agenclientbuilds
WORKDIR /appbuilds
COPY . .
RUN go mod tidy
RUN go build -o binary

# ---- Svelte Base ----
FROM node:lts-alpine AS totosveltebaseagen
WORKDIR /svelteapp
COPY [ "frontend/package.json" , "frontend/yarn.lock"  , "./"]

# ---- Svelte Dependencies ----
FROM totosveltebaseagen AS totosveltedepagen
RUN yarn
RUN cp -R node_modules prod_node_modules

#
# ---- Svelte Builder ----
FROM totosveltebaseagen AS totosveltebuilderagen
COPY --from=totosveltedepagen /svelteapp/prod_node_modules ./node_modules
COPY ./frontend .
RUN yarn build

FROM alpine:latest as agenclientrelease
WORKDIR /app
RUN apk add tzdata
COPY --from=totosveltebuilderagen /svelteapp/dist ./frontend/dist
COPY --from=agenclientbuilds /appbuilds/binary .
COPY --from=agenclientbuilds /appbuilds/.env /app/.env
ENV PORT=5051
ENV PATH_API_BACKEND="http://128.199.241.112:7072/"
ENV TZ=Asia/Jakarta

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT [ "./binary" ]