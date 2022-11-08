import { Sprite } from "./sprite.js";
import { Entity } from "./entity.js";

class Soldier extends Entity {
    public readonly id: string;

    public health: number = 100;
    public direction: boolean = false;
    public actions: string[] = [];
    public faction: 'red' | 'blue';
    
    public sprite: Sprite;

    constructor(id: string, faction: Soldier['faction'], direction: boolean, pos: Position, actions?: string[]) {
        super(pos);

        this.actions = actions || [];

        this.direction = direction;
        this.faction = faction;

        this.id = id;

        this.sprites.set('idle', new Sprite({
            source: `./assets/images/soldiers/${this.faction}/idle.png`,
            animationType: 'loop',
            holdFrames: 3,
            maxFrames: 5,
            id: 'idle',
        }));

        const sprite = this.sprites.get('idle')!;

        this.sprite = Object.assign(
            Object.create(Object.getPrototypeOf(sprite)), sprite
        )
    }

    update(soldier: SoldierModel): void {
        for (const [key, value] of Object.entries(soldier)) {
            this[key as 'direction'] = value;
        }
    }

    draw(ctx: CanvasRenderingContext2D, scale: number) {
        this.sprite.draw(ctx, scale, this.position, !this.direction);
    }

}

export { Soldier }
