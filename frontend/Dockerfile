FROM node:18.12.1

RUN mkdir /frontend
WORKDIR /frontend

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY . .

RUN npm run build

CMD ["npm", "run", "dev"]