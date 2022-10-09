const canvas = document.getElementById('canvas') as HTMLCanvasElement
const ctx = canvas.getContext('2d')

function game() {

}

setInterval(
    () => requestAnimationFrame(game), 1000 / 60
)
