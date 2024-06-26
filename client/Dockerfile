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


# Define build arguments
ARG REACT_APP_BACKEND_URL
ARG REACT_APP_BACKEND_PORT

# Set environment variables based on build arguments for the build process
ENV REACT_APP_BACKEND_URL=$REACT_APP_BACKEND_URL
ENV REACT_APP_BACKEND_PORT=$REACT_APP_BACKEND_PORT

RUN npm run build

# Stage 2: Create the production image
FROM nginxinc/nginx-unprivileged:alpine3.18-slim
COPY --from=builder /app/build /usr/share/nginx/html
#COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
