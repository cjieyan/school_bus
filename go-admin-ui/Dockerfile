FROM node:12.18.4
USER root
ENV HOST 0.0.0.0
RUN mkdir -p /app
COPY . /app
WORKDIR /app
RUN npm -g config set user root && \
    npm rebuild node-sass && \
    npm install --registry=https://registry.npm.taobao.org
EXPOSE 9527
CMD ["npm", "run", "dev"]
