const onload = (request, resolve, reject) => () => {
  const data = JSON.parse(request.responseText);
  if (request.status >= 200 && request.status < 400) {
    resolve(data);
  } else {
    reject(data);
  }
};

const onerror = reject => () => {
  reject(arguments);
};

const post = (url, payload = {}) =>
  new Promise((resolve, reject) => {
    var request = new XMLHttpRequest();
    request.open("POST", url, true);
    request.setRequestHeader(
      "Content-Type",
      "application/x-www-form-urlencoded; charset=UTF-8",
    );
    request.onload = onload(request, resolve, reject);
    request.onerror = onerror(reject);
    request.send(payload);
  });
const get = url =>
  new Promise((resolve, reject) => {
    var request = new XMLHttpRequest();
    request.open("GET", url, true);
    request.onload = onload(request, resolve, reject);
    request.onerror = onerror(reject);
    request.send();
  });

const login = () => post("/api/auth/login");
const logout = () => post("/api/auth/logout");
const search = query => get(`/api/search?query=${query}`);

const loginBtn = document.getElementById("login");
const logoutBtn = document.getElementById("logout");
const searchBtn = document.getElementById("search");
const searchInput = document.getElementById("search-input");
const text = document.getElementById("text");

const showResponse = data => {
  text.innerHTML = JSON.stringify(data);
};

loginBtn &&
  loginBtn.addEventListener("click", ev => {
    ev.preventDefault();
    login().then(showResponse, showResponse);
  });

logoutBtn &&
  logoutBtn.addEventListener("click", ev => {
    ev.preventDefault();
    logout().then(showResponse, showResponse);
  });

searchBtn &&
  searchBtn.addEventListener("click", ev => {
    ev.preventDefault();
    search(searchInput.value).then(showResponse, showResponse);
  });
