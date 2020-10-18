FROM node:12.18.4-alpine3.9
ENV HOST 0.0.0.0
RUN mkdir -p /app
COPY . /app
WORKDIR /app
EXPOSE 9527
RUN npm rebuild node-sass && npm install -g npm --registry=https://registry.npm.taobao.org && \
    npm install --registry=https://registry.npm.taobao.org
CMD ["npm", "run", "dev"]
