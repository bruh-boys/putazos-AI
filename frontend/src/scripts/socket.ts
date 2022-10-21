//import { canvas, ctx } from './@module.js'

const socket = new WebSocket('ws://localhost:8080/game/socket/', [])

socket.onopen = (ev) => {
    console.log('Connected to server')
}

socket.onmessage = (event) => {
    console.log(JSON.parse(JSON.stringify(event.data)))

}

socket.onclose = (ev) => {
    console.log('Disconnected from server')
}

socket.onerror = (ev) => {
    console.log(ev)
}

setTimeout(() => {
    socket.send(JSON.stringify({
        action: 'down',
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

