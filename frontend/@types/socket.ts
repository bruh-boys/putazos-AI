type SocketValue <T extends 'join' | 'update'> =
    T extends 'join' ? World : {
        projectiles: EntityModel<ProjectileModel>[],
        soldiers: EntityModel<SoldierModel>[]
    }

interface SocketResponse {
    type: 'join' | 'update',
    data: SocketValue<SocketResponse['type']>
}

interface SocketRequest {
    type: 'action',
    data: {
        state: boolean,
        key: string
    }
}
