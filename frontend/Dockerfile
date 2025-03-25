# using node base image
FROM node:20

WORKDIR /app

# copying dependencies from source folder
COPY package*.json ./

# isntalling dependencies
RUN npm install

# Copying rest of the files
COPY . .

# exposes Angular on the port 4200
EXPOSE 4200

# Démarrer Angular en mode dev
CMD ["npm", "run", "start"]
