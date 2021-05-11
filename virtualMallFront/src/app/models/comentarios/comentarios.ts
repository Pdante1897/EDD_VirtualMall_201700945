export class Comentarios {
    Comentarios: Comentario[]=[]
}
export class Comentario{
    Id : number
    Dpi: number
    Cadena : string
    SubComentarios: Comentario[]

    constructor(id:number, dpi:number, cadena:string, sub: Comentario[]){
        this.Id=id
        this.Dpi= dpi
        this.Cadena=cadena
        this.SubComentarios=sub
    }
}

export class SubComentarios {
    Comentarios: Comentario[]=[]
}
