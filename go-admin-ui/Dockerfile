FROM node:12.18.4-alpine3.9
ENV HOST 0.0.0.0
RUN mkdir -p /app
COPY . /app
WORKDIR /app
EXPOSE 9527
RUN npm install -g cnpm --registry=https://registry.npm.taobao.org && \
    cnpm install --registry=https://registry.npm.taobao.org
CMD ["cnpm", "run", "dev"]
