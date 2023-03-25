let socket = new WebSocket("ws://localhost:8080/ws")

// send message from the chat form
document.forms.chat.onsubmit = function() {
    // TODO: get name from session

    // get name
    let user = "anonymous" 

    // get message
    let userMessage = this.message.value;
    const outgoingMessage = {name: user, message: userMessage};

    // send message
    socket.send(JSON.stringify(outgoingMessage));
    // reset input message box
    this.message.value = "";
    return false;
}


// message receive - show the message in div#messages
socket.onmessage = function(event){
    // parse message received from server
    let data = JSON.parse(event.data);
    let name = data.name;
    let msg = data.message + "<br>";
    // if message is longer than the length of the message box add a line break
    if (msg.length > 30){
        msgArray = msg.split("");
        for (let i = 30; i < msg.length; i++){
            if (i % 30 == 0) {
                msgArray.splice(i,0,"<br>");
            }
        };
        msg = msgArray.join("")
    }

    // build message
    message = name.bold() + ":  " + msg;
    let msgcontainer = document.getElementById('messages');
    // append message to html message area
    msgcontainer.innerHTML += message;

    // auto scroll if overflow is reached
    msgcontainer.scrollTop = msgcontainer.scrollHeight;
}
