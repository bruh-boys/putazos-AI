interface SpriteModel {
    animationType?: 'once' | 'loop';
    holdFrames?: number,
    maxFrames?: number,
    source: string,
    id: string,
}

interface Sprite extends SpriteModel {
    framesCurrent: number,
    framesElapsed: number,
    framesHold: number,
    framesMax: number,
    scale: number,
    image: HTMLImageElement

    update: (
        this: Sprite, ctx: CanvasRenderingContext2D,
        pos: Position
    ) => void
    draw: (
        this: Sprite, ctx: CanvasRenderingContext2D,
        pos: Position
    ) => void
}
