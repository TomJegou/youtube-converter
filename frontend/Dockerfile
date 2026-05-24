FROM tomyj/my-env:node AS builder
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM tomyj/my-env:node
ENV NODE_ENV=production
ENV API_HOSTNAME=api:46460
ENV PORT=45450
WORKDIR /app
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
EXPOSE 45450
CMD ["node", "server.js"]
