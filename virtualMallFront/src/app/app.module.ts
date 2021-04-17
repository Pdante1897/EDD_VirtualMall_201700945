import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { InicioComponent } from './componentes/inicio/inicio.component';
import { TiendasComponent } from './componentes/tiendas/tiendas.component';
import { CarritoComponent } from './componentes/carrito/carrito.component';
import { PedidosComponent } from './componentes/pedidos/pedidos.component';
import {HttpClientModule} from '@angular/common/http';
import { ProductosComponent } from './componentes/productos/productos.component';
import { AdminComponent } from './componentes/admin/admin.component';


@NgModule({
  declarations: [
    AppComponent,
    InicioComponent,
    TiendasComponent,
    CarritoComponent,
    PedidosComponent,
    ProductosComponent,
    AdminComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
