const sprites = new Map<string, Sprite>();

function draw(
    this: Sprite, ctx: CanvasRenderingContext2D,
    pos: Position
) {
    ctx.drawImage(
        this.image,
        this.framesCurrent * (this.image.width / this.framesMax),
        0,
        this.image.width / this.framesMax,
        this.image.height,
        pos.x,
        pos.y,
        (this.image.width / this.framesMax) * this.scale,
        this.image.height * this.scale,
    )
}

function update(
    this: Sprite, ctx: CanvasRenderingContext2D,
    pos: Position
) {
    this.draw(ctx, pos)

    if (this.framesCurrent < this.framesMax - 1) this.framesCurrent++
    else this.framesCurrent = 0;
}

function sprite(this: Sprite, id: string, sprite: SpriteModel) {
    for (const key in sprite) { // @ts-ignore
        this[key as keyof Sprite] = sprite[key as keyof SpriteModel];
    }

    this.framesHold = this.framesHold
        ? this.framesHold
        : 30;

    this.framesMax === undefined &&
        (this.framesMax = 1)

    this.image.src = this.source;
    this.image = new Image();

    this.framesCurrent = 0;
    this.framesElapsed = 0;

    this.draw = draw.bind(this);
    this.update = update.bind(this);

    sprites.set(id, this);
}

const Sprite = sprite as unknown as {
    new (id: string, sprite: SpriteModel): Sprite;
}

export { Sprite, sprites };
