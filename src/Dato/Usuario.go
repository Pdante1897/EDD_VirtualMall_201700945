package Dato

type JsonUsuarios struct {
	Usuarios []Usuario `json:"Usuarios"`
}

type Usuario struct {
	Dpi      int64  `json:"Dpi"`
	Nombre   string `json:"Nombre"`
	Correo   string `json:"Correo"`
	Password string `json:"Password"`
	Cuenta   string `json:"Cuenta"`
}

type BusquedaUsuario struct {
	Dpi int64 `json:"Dpi"`
}
