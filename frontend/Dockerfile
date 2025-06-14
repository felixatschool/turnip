# --- DEPENDENCIES STAGE
FROM node:20-alpine AS deps

WORKDIR /app

COPY svelte.config.js ./

COPY package*.json ./

RUN npm install

# --- DEVELOPMENT STAGE
FROM node:20-alpine AS dev

WORKDIR /app

ENV NODE_ENV=development

COPY --from=deps /app/node_modules ./node_modules

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host"]

# --- BUILD STAGE
FROM node:20-alpine AS builder

WORKDIR /app

ENV NODE_ENV=development

COPY --from=deps /app/node_modules ./node_modules

COPY . .

RUN npm run build

# --- PRODUCTION STAGE
FROM nginx:alpine AS prod

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/build .

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=3s CMD wget --spider -q http://localhost/ || exit 1

CMD ["nginx", "-g", "daemon off;"]
