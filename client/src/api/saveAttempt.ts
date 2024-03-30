import axios from 'axios';

export async function saveAttempt(saveAttemptData: any): Promise<void> {
  try {
    const response = await axios.post("http://localhost:8080/tutorials/code/attempt/create", saveAttemptData, {
      headers: {
        "Content-Type": "application/json",
      }
    });
  } catch (error) {
    console.error("An error occurred:", error);
    throw error;
  }
}
