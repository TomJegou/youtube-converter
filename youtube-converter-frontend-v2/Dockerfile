FROM tomyj/my-env:node
ENV API_HOSTNAME=api:46460
WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm ci
COPY . .
RUN npx next build
CMD [ "npx", "next", "start", "-p", "45450" ]
