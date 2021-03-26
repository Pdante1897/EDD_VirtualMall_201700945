import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { InicioComponent } from "./componentes/inicio/inicio.component";
import { TiendasComponent } from "./componentes/tiendas/tiendas.component";
import { CarritoComponent } from "./componentes/carrito/carrito.component";
import { PedidosComponent } from "./componentes/pedidos/pedidos.component";
import { ProductosComponent } from "./componentes/productos/productos.component";


const routes: Routes = [{
  path: '',
  component: InicioComponent,
},
{
  path: 'tiendas',
  component: TiendasComponent,
},
{
  path: 'productos/:nombre/:departamento/:calificacion',
  component: ProductosComponent,
},
{
  path: 'productos',
  component: ProductosComponent,
}
,

{
  path: 'carrito',
  component: CarritoComponent,
}
,
{
  path: 'pedidos',
  component: PedidosComponent,
}
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
