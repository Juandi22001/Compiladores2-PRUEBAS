"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class EtiquetaInicio {
    constructor(listaAtributos, linea, columna) {
        this.linea = linea;
        this.columna = columna;
        listaAtributos.forEach(atributo => {
            if (atributo.nombre == "version") {
                this.version = atributo.valor;
            }
            else if (atributo.nombre == "encoding") {
                this.encoding = atributo.valor;
            }
        });
    }
}
exports.EtiquetaInicio = EtiquetaInicio;
