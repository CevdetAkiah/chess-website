FROM node:21.1-alpine AS builder


ENV NODE_ENV=production
ENV PORT=3000
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json      ./
COPY package-lock.json ./
RUN npm install

COPY src ./src
COPY public ./public


RUN npm run build

# Stage 2: Create the production image
FROM nginxinc/nginx-unprivileged:alpine3.18-slim
COPY --from=builder /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
