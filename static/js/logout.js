// attach a form to the logout link to change method from GET to POST

window.addEventListener("load", function () {
  this.document.getElementById("logoutLink").addEventListener("click", function (e) {
    e.preventDefault();
    var myForm = document.getElementById("logoutForm")
    myForm.submit();
  })
})