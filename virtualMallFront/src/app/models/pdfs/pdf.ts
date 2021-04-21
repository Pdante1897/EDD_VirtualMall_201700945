export class Pdf {
    Nombre: string
    Ruta: string

    constructor(nombre: string, ruta: string){
        this.Nombre=nombre
        this.Ruta=ruta
    }
    

}
export class ArrPdf{
    Archivos: Pdf[]=[]
}
