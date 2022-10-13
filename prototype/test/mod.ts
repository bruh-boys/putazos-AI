const socket = new WebSocket("ws://localhost:8080/game/socket/");

socket.onopen = () => {
    console.log("connected");

    socket.send(JSON.stringify({
        action: "move-left",
        active: true
    }))
}

socket.onmessage = (event) => {
    console.log(event.data);
}

socket.onclose = () => {
    console.log("disconnected");
}

socket.onerror = (error) => {
    console.error(error);
}
