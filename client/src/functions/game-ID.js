import { v4 as uuidv4 } from 'uuid';

export const randGameID = () => {
  const timestamp = Date.now(); // Get current timestamp
  const randomString = uuidv4(); // Generate a random string using uuid

  // Combine timestamp and random string
  const gameID = `${timestamp}${randomString}`;

  return gameID;
}