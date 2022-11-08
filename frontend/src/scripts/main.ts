import { canvas, init, overwrite_map, overwrite_entities } from './game.js'

const socket = new WebSocket('ws://localhost:8080/game/socket/')
const keypress: { [key: string]: boolean } = {}

function socket_request(req: SocketRequest) {
    socket.send(JSON.stringify(req))

}

socket.onmessage = (event) => {
    const { type, data } = JSON.parse(event.data) as SocketResponse

    console.log(type)

    switch (type) {
        case 'update':
            overwrite_entities((data as SocketValue<'update'>))

            break
        case 'join':
            overwrite_map(JSON.parse(data as any) as SocketValue<'join'>)

            setTimeout(() => init(), 2000)
    }
}

window.addEventListener('keydown', (event) => {
    event.preventDefault();
    console.log(event.key)
    socket_request({ type: 'action', data: { state: true, key: event.key } })
    keypress[event.key] = true

}, false)

window.addEventListener('keyup', (event) => {
    event.preventDefault();
    socket_request({ type: 'action', data: { state: false, key: event.key } })
    keypress[event.key] = false

}, false)
