let socket = new WebSocket("ws://localhost:8080/ws")

// TODO: protect against script insertion
// send message from the chat form
document.forms.chat.onsubmit = function() {
    // get message
    let userMessage = this.message.value;
    // const outgoingMessage = {name: user, message: userMessage};

    // send message
    socket.send(userMessage);
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
    let message = name.bold() + msg;
    // if message is longer than the length of the message box add a line break
    if (message.length > 42){ 
        for (let i = 42; i < message.length; i++){
            if (i % 42 == 0) {
                // this is potentially expensive as we're potentially splitting the string multiple times
                // but the text limit is more accurate to the chat box size
                msgArray = message.split(""); 
                msgArray.splice(i,0,"<br>");
                message = msgArray.join("")
            }
        };
    }


    let msgcontainer = document.getElementById('messages');
    // append message to html message area
    msgcontainer.innerHTML += message;

    // auto scroll if overflow is reached
    msgcontainer.scrollTop = msgcontainer.scrollHeight;
}
