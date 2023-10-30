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


// // Function to set the grid layout when the chss game is loaded
// function setGridLayout(){
//     const gridContainer = document.querySelector('.grid-container');
//     const chessContainer = document.querySelector('.chess-container');
    

//     // Check if the chess container is loaded
//     // For demonstration purposes, let's assume it's loaded after 3 seconds (3000 ms)
//     setTimeout(() => {
//         gridContainer.style.gridTemplateColumns = '1fr 1fr' // Set two columns for chat and chess
//         chessContainer.style.display = 'block'; // Show the chess container

//         const chessCanvas = chessContainer.querySelector('chessCanvas');


//     }, 3000); // Adjust this time based on your actual loading time for the chess game
// }

// // Call the function when the page is loaded
// window.addEventListener('load', setGridLayout)