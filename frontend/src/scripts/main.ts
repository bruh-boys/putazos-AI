import { canvas, init, overwrite_entities, overwrite_map } from './game.js'

const socket = new WebSocket('ws://localhost:8080/game/socket/')
const keypress: { [key: string]: boolean } = {}

function socket_request(req: SocketRequest) {
    socket.send(JSON.stringify(req))

}

socket.onmessage = (event) => {
    const { type, data } = JSON.parse(event.data) as SocketResponse

    switch (type) {
        case 'data':
            overwrite_entities((data as SocketValue<'data'>))

            break
        case 'join':
            overwrite_map(JSON.parse(data as any) as SocketValue<'join'>)

            setTimeout(() => init(), 2000)
    }
}

canvas.addEventListener('keydown', (event) => {
    socket_request({ type: 'action', data: { state: true, key: event.key } })
    keypress[event.key] = true

})

canvas.addEventListener('keyup', (event) => {
    socket_request({ type: 'action', data: { state: false, key: event.key } })
    keypress[event.key] = false

})
