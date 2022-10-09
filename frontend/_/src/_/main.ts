import { entities } from './models/@module.js'

const canvas = document.getElementById('canvas') as HTMLCanvasElement,
    ctx = canvas.getContext('2d');


function game() {

}

setInterval(
    () => requestAnimationFrame(game), 1000 / 60
);
