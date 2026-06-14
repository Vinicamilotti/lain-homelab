FROM node:14-alpine AS download
WORKDIR /opt/node_app

RUN wget https://github.com/excalidraw/excalidraw/archive/refs/tags/v0.17.3.tar.gz -O excalidraw.tar.gz
RUN tar -zvxf excalidraw.tar.gz
RUN mv excalidraw-* excalidraw

FROM node:14-alpine AS build

ENV PUB_SRV_NAME draw.$DOMAIN_NAME
ENV PUB_SRV_NAME_WS draw-ws.$DOMAIN_NAME
ENV NODE_ENV development
ENV TZ $TZ
ENV REACT_APP_BACKEND_V2_GET_URL https://$PUB_SRV_NAME/api/v2/
ENV REACT_APP_BACKEND_V2_POST_URL https://$PUB_SRV_NAME/api/v2/post
ENV REACT_APP_WS_SERVER_URL https://$PUB_SRV_NAME_WS/

WORKDIR /opt/node_app

COPY --from=download /opt/node_app/excalidraw/package.json /opt/node_app/excalidraw/yarn.lock ./
RUN yarn --ignore-optional --network-timeout 600000

COPY --from=download /opt/node_app/excalidraw/ .
RUN yarn build:app:docker

FROM nginx:1.21-alpine

COPY --from=build /opt/node_app/build /usr/share/nginx/html

HEALTHCHECK CMD wget -q -O /dev/null http://localhost || exit 1