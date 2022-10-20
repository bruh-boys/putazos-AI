/*import { map } from "./map.ts";

const obj: any = {
    radius: {
        x: 100,
        y: 100,
    },
    defs: []
}

const id_replacer: {
    [key: string]: string
} = {
    'Small_plataform': 'small-plataform',
    'Plataforms': 'plataform',
    'Collisions': 'collision',
    'Spawns': 'spawn-point'
}

for (const def of map.defs) {
    const id = id_replacer[def.__identifier]

    obj.defs.push({
        id: id,
        image_src: def.image_src,
        visible: def.visible,
        positions: [],
    })

    for (const tiles of def.gridTiles) {
        obj.defs[obj.defs.length - 1].positions.push({
            x: tiles.px[0],
            y: tiles.px[1],
        })
    }

}

Deno.writeTextFile("mymap.json", JSON.stringify(obj, null, 4), {
    create: true,
})

console.log(obj)*/