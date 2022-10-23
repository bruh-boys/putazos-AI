const socket = new WebSocket('ws://localhost:8080/game/socket/', [])
export let entities: any

socket.onopen = (ev) => {
    console.log('Connected to server')
}

socket.onmessage = (event) => {
    entities = JSON.parse(event.data)

}

socket.onclose = (ev) => {
    console.log('Disconnected from server')
}

socket.onerror = (ev) => {
    console.log(ev)
}

setTimeout(() => {
    socket.send(JSON.stringify({
        action: 'left',
        active: true
    }))
}, 1000)

/*
function game() {

}

setInterval(
    () => requestAnimationFrame(game), 1000 / 60
)
*/

