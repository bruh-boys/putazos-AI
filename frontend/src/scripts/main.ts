//import { canvas, ctx, entites } from './@module.js'

const socket = new WebSocket('ws://127.0.0.1:8080/game/gateway', [])

socket.onopen = (ev) => {
    console.log('Connected to server')

    /*socket.send(JSON.stringify({

    }))*/
}

socket.onmessage = (event) => {
    console.log(event.data)

}

socket.onclose = (ev) => {
    console.log('Disconnected from server')
}

socket.onerror = (ev) => {
    console.log(ev)
}

/*
function game() {

}

setInterval(
    () => requestAnimationFrame(game), 1000 / 60
)
*/