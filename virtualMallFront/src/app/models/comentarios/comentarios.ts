export class Comentarios {
    Comentarios: Comentario[]=[] 
}
export class Comentario{
    Id : number
    Cadena : string

    constructor(id:number, cadena:string){
        this.Id=id
        this.Cadena=cadena
    }
}
