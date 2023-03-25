let socket = new WebSocket("ws://localhost:8080/ws")

// TODO: protect against script insertion
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

        // build message
    let message = name.bold() + ":  " + msg;
    // if message is longer than the length of the message box add a line break
    if (message.length > 38){
        msgArray = message.split("");
        for (let i = 38; i < message.length; i++){
            if (i % 38 == 0) {
                msgArray.splice(i,0,"<br>");
            }
        };
        message = msgArray.join("")
    }


    let msgcontainer = document.getElementById('messages');
    // append message to html message area
    msgcontainer.innerHTML += message;

    // auto scroll if overflow is reached
    msgcontainer.scrollTop = msgcontainer.scrollHeight;
}
