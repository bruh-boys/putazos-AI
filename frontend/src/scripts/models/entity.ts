import { Sprite } from "./sprite.js"

class Entity {
    protected readonly sprites: Map<string, Sprite> = new Map()
    public position: Position

    constructor(pos?: Position, sprites?: SpriteModel[]) {
        for (const sprite of (sprites || [])) {
            this.sprites.set(sprite.id, new Sprite(sprite))
        }

        this.position = pos || { x: 0, y: 0 }
    }

    protected draw(ctx: CanvasRenderingContext2D, scale: number, pos: Position) {

    }
}

export { Entity }
