const socket = new WebSocket("ws://localhost:8080/game/socket/", ["tpc"]);

socket.onopen = () => {
    console.log("connected");

    socket.send("");
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
