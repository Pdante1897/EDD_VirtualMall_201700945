export class Tienda {
    Departamento:string
    Nombre: string
    Descripcion: string
    Contacto: string
    Calificacion: number
    Logo: string

    

    constructor(_departamento:string, _nombre: string, _descripcion: string, _contacto: string, _calificacion: number, _logo: string){
        this.Departamento = _departamento
        this.Nombre=_nombre
        this.Contacto=_contacto
        this.Descripcion=_descripcion
        this.Calificacion=_calificacion
        this.Logo=_logo
    }
}

export class ArrTienda{
    Tiendas: Tienda[]=[]
    
    constructor(){

    }
    setTiendas(ListaTiendas: Tienda[]){
        this.Tiendas=ListaTiendas
    }
}

