





const form = document.getElementById("userform")
form.onsubmit = function (e) {
  e.preventDefault();
  let xhr = new XMLHttpRequest();
  let formData = new FormData(form);
  let csrf = document.getElementById("csrf_token")
  let path = "/signupAccount"

  if (Array.from(formData.keys()).length == 2) {
    path = "/authenticate"
  }
  // open the request
  xhr.open('POST', 'http://localhost:8080' + path)
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.setRequestHeader("x-csrf-token", csrf.value);

  // send the form data
  xhr.send(JSON.stringify(Object.fromEntries(formData)));

  // AJAX on form submission reset the form but not the page
  xhr.onreadystatechange = function () {
    if (xhr.readyState == XMLHttpRequest.DONE) {
      form.reset();
      window.location.href = "/";
    }

  }

  // avoid page refresh
  return false
}
