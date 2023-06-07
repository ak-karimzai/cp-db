import { store } from "./store";

let Security = {
  saveToken: (token) => {
    let date = new Date();
    let expDays = 30;
    date.setTime(date.getTime() + expDays * 24 * 60 * 60 * 1000);
    const expires = "expires=" + date.toUTCString();
    document.cookie =
      "token=" +
      JSON.stringify(token) +
      ";" +
      expires +
      ";Path=/" +
      ";" +
      "SameSite=Strict;" +
      "Secure";
  },
  setTokenNull: () => {
    document.cookie =
      "token=" +
      ";" +
      "Expires=Thu, 01 Jan 1970 00:00:01 GMT" +
      ";Path=/" +
      ";" +
      "SameSite=Strict;" +
      "Secure";
  },
  requestOptions: (payload, method = "POST") => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", "Bearer " + store.token);

    return {
      method: method,
      body: JSON.stringify(payload),
      headers: headers,
    };
  },
  requestOptionsWithoutBody: (method = "GET") => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", "Bearer " + store.token);

    return {
      method: method,
      headers: headers,
    };
  },
  isAdmin: () => {
    return store.user.user_role === 'admin';
  },
};

export default Security;
