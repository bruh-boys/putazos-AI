import { Entity } from './models/@module.js'

const canvas = document.getElementById('canvas') as HTMLCanvasElement
const ctx = canvas.getContext('2d')

const entites: Entity[] = []

export { canvas, ctx, entites }
