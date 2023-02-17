

const profileOptions = document.querySelector('[role="profileoptions"]');

profileOptions.addEventListener('click', showForm);

// function changeOptionFocus(e) {
//   const targetOption = e.target
//   targetOption.addEventListener('click', showForm);

// }

// show form when option is selected
function showForm(e) {
  form = e.target.querySelector('[role="form"]');
  form.removeAttribute("hidden");
}

// send new user name to web api
function sendUserName() {
  let username = profileOptions.querySelector('#username');
  sendJSON(username)
}


// send updated email address to web api
function sendEmail() {
  let email = profileOptions.querySelector('#email');
  sendJSON(email)
}

// send updated email address to web api
function sendPassword() {
  let password = profileOptions.querySelector('#password');
  sendJSON(password)
}

function sendDelete() {
  let deleteID = profileOptions.querySelector('#delete')
  sendJSON(deleteID)
}

// send data as JSON to web api
function sendJSON(element) {
  let path = ""
  let method = "PUT"
  let option = element.id
  let csrf = document.getElementById("csrf_token")

  // set path
  switch (option) {
    case "username":
      path = "/updateUserName"
      break;
    case "email":
      path = "/updateEmail"
      break;
    case "password":
      path = "/updatePassword"
      break;
    case "delete":
      path = "/deleteUser"
      method = "DELETE"
  }

  // create XHR object
  let xhr = new XMLHttpRequest();
  let url = "http://localhost:8080" + path;

  // open connection to web api
  xhr.open(method, url);

  // set header
  xhr.setRequestHeader("x-csrf-token", csrf.value)
  xhr.setRequestHeader("Content-Type", "application/json");

  // send data as JSON
  var data = JSON.stringify({ option: element.value })

  xhr.send(data);
}
