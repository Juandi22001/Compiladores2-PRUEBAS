"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const Fila_1 = require("./Fila");
const Tipos_1 = require("./Tipos");
class Atributo {
    constructor(nombre, valor, linea, columna) {
        this.etiquetaContendora = null;
        this.nombre = nombre;
        this.valor = valor;
        this.linea = linea;
        this.columna = columna;
    }
    getAmbito() {
        let listaAmbito = [];
        for (let etiqueta = this.etiquetaContendora; etiqueta != null; etiqueta = etiqueta.padre) {
            listaAmbito.push(etiqueta.getName());
        }
        listaAmbito.push("GLOBAL");
        return listaAmbito;
    }
    getAsRowTable() {
        return (new Fila_1.Fila(this.nombre, Tipos_1.Tipos.ATRIBUTO, this.getAmbito(), this.linea, this.columna, this.imprimir()));
    }
    imprimir() {
        let texto = "";
        texto = this.nombre + "=" + this.valor;
        return texto;
    }
}
exports.Atributo = Atributo;
