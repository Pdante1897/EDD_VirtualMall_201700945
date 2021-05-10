export class Usuario {
    Dpi: number 
    Nombre   :string 
	Correo   :string 
	Password :string 
	Cuenta   :string 

    constructor(_dpi: number, nombre :string, correo :string, password: string, cuenta: string){
        this.Dpi=_dpi
        this.Nombre=nombre
        this.Correo=correo
        this.Password = password
        this.Cuenta = cuenta
    }
}

export class BusquedaU {
    Dpi: number 
    constructor(_dpi: number){
        this.Dpi=_dpi

    }
}