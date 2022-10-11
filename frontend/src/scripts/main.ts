import { canvas, ctx, entites } from './@module.js'

const socket = new WebSocket('ws://localhost:8080')

function game() {

}

setInterval(
    () => requestAnimationFrame(game), 1000 / 60
)
