FROM node:18.18.2

ENV CONTAINER_PATH /var/www/docker-vue
WORKDIR $CONTAINER_PATH
COPY package*.json ./
RUN npm install --production
COPY . .

EXPOSE 8080:8080
CMD ["npm", "run", "serve"]